#!/usr/bin/env python3
"""Check the generated guide's responsive layout and navigation in Chromium.

Requires playwright==1.57.0 and its Chromium browser. The tested HTML, CSS, font,
and script are the actual build artifacts, loaded inline so CI needs no server
or external network. Run after `go run ./cmd/workbench generate --out docs`.
"""
import argparse
import base64
import json
from pathlib import Path
import re

from playwright.sync_api import sync_playwright


def document(site, filename):
    source = (site / filename).read_text(encoding="utf-8")
    styles = (site / "style.css").read_text(encoding="utf-8")
    styles += (site / "assets/mobile.css").read_text(encoding="utf-8")
    font = base64.b64encode((site / "assets/JuliaMono-Regular.woff2").read_bytes()).decode()
    styles = styles.replace('url("assets/JuliaMono-Regular.woff2")',
                            f'url("data:font/woff2;base64,{font}")')
    source = re.sub(r"<link\b[^>]*>", "", source)
    source = re.sub(r"<script\b[^>]*src=[^>]*>\s*</script>", "", source)
    return source.replace("</head>", "<style>" + styles + "</style></head>"), (site / "tabs.js").read_text(encoding="utf-8")


def check_menu(page, output):
    summary = page.locator(".mobile-nav > summary")
    menu = page.locator("[data-mobile-nav]")
    summary.tap()
    assert menu.evaluate("el => el.open"), "touch must open the menu"
    links = page.locator(".mobile-nav nav a")
    assert summary.bounding_box()["height"] >= 44
    for index in range(links.count()):
        assert links.nth(index).bounding_box()["height"] >= 44, "menu links need a 44px touch target"
    page.screenshot(path=str(output / "menu-mobile.png"))
    page.locator(".mobile-nav .nav-component").filter(has_text="Deployment comparison").tap()
    page.wait_for_timeout(100)
    assert not menu.evaluate("el => el.open"), "choosing a section must close the menu"
    assert page.evaluate("document.activeElement.id") == "screen/deployment", "focus must leave the hidden menu"
    top = page.locator('[id="screen/deployment"]').bounding_box()["y"]
    assert 48 <= top < 150, f"fragment target is obscured or misplaced: {top}"
    page.screenshot(path=str(output / "deployment-mobile.png"))
    summary.tap()
    page.locator('.mobile-nav .nav-level').filter(has_text="Screens").tap()
    page.wait_for_timeout(100)
    assert not menu.evaluate("el => el.open"), "atomic navigation must close the menu"
    assert page.evaluate("document.activeElement.id") == "level-screens"
    assert 48 <= page.locator('[id="level-screens"]').bounding_box()["y"] < 150
    summary.tap()
    page.keyboard.press("Escape")
    assert not menu.evaluate("el => el.open")
    assert summary.evaluate("el => el === document.activeElement"), "Escape must restore focus"
    summary.press("Space")
    page.wait_for_function(
        "() => document.querySelector('[data-mobile-nav]').open",
        timeout=2000,
    )
    page.mouse.click(10, 800)
    assert not menu.evaluate("el => el.open"), "outside click must close the menu"
    summary.tap()
    page.set_viewport_size({"width": 1440, "height": 844})
    page.set_viewport_size({"width": 390, "height": 844})
    page.wait_for_timeout(100)
    assert not menu.evaluate("el => el.open"), "breakpoint changes must not leave the menu open"


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--site", type=Path, default=Path("docs/dist"))
    parser.add_argument("--output", type=Path, default=Path("artifacts/mobile-guide"))
    parser.add_argument("--chromium-path")
    args = parser.parse_args()
    args.output.mkdir(parents=True, exist_ok=True)
    pages = sorted(args.site.glob("*.html"))
    assert len(pages) >= 7, "build the complete guide before running browser checks"
    sources = {file.name: document(args.site, file.name) for file in pages}
    results = []
    with sync_playwright() as playwright:
        options = {"executable_path": args.chromium_path} if args.chromium_path else {}
        browser = playwright.chromium.launch(**options)
        for width in (320, 360, 390, 430, 760, 768, 1440):
            for filename, (html, javascript) in sources.items():
                context = browser.new_context(viewport={"width": width, "height": 844},
                                              device_scale_factor=1, has_touch=width <= 760,
                                              is_mobile=width <= 760)
                page = context.new_page()
                page.set_default_timeout(5000)
                page.set_content(html)
                page.add_script_tag(content=javascript)
                page.evaluate("document.fonts.ready")
                values = page.evaluate("""() => ({
                    width: innerWidth, scrollWidth: document.documentElement.scrollWidth,
                    headingTop: document.querySelector('h1').getBoundingClientRect().top,
                    menuVisible: !!document.querySelector('.mobile-nav-bar').offsetHeight,
                    menuOpen: document.querySelector('[data-mobile-nav]').open,
                    desktopVisible: !!document.querySelector('.desktop-nav').offsetWidth
                })""")
                assert values["scrollWidth"] <= width + 1, (filename, values)
                assert values["menuVisible"] == (width <= 760), (filename, values)
                assert values["desktopVisible"] == (width > 760), (filename, values)
                if width <= 760:
                    assert values["headingTop"] < 180, (filename, values)
                    assert not values["menuOpen"], (filename, values)
                if filename in ("examples.html", "development.html") and width in (390, 1440):
                    page.screenshot(path=str(args.output / f"{filename.removesuffix('.html')}-{width}.png"))
                expected_levels = ["Foundations", "Atoms", "Molecules", "Organisms", "Templates", "Screens"]
                actual_levels = page.locator(".desktop-nav .nav-level").all_text_contents()
                assert actual_levels == expected_levels, (filename, actual_levels)
                if filename == "components.html":
                    assert page.locator(".atomic-level").evaluate_all("els => els.map(e => e.id)") == ["level-atoms", "level-molecules", "level-organisms"]
                    assert page.locator('[id="level-molecules"] .component-source h5').all_text_contents() == ["Go usage#"] * 10
                if filename == "development.html":
                    tasks = page.locator(".dev-paths a")
                    assert tasks.count() == 3, "developer guide needs three task paths"
                    for task in tasks.all():
                        target = task.get_attribute("href")
                        assert target.startswith("#") and page.locator(f'[id="{target[1:]}"]').count() == 1, target
                if filename == "examples.html" and width == 390:
                    check_menu(page, args.output)
                results.append({"page": filename, **values})
                context.close()
        context = browser.new_context(viewport={"width": 390, "height": 844}, java_script_enabled=False)
        page = context.new_page()
        page.set_content(sources["examples.html"][0])
        menu = page.locator("[data-mobile-nav]")
        assert menu.get_attribute("open") is None
        page.locator(".mobile-nav summary").click()
        assert menu.get_attribute("open") is not None, "native menu must work without JavaScript"
        context.close()
        browser.close()
    (args.output / "results.json").write_text(json.dumps(results, indent=2), encoding="utf-8")
    print(f"Verified {len(results)} page/width layouts, touch, keyboard, fragments, resize, and no-JS menu.")


if __name__ == "__main__":
    main()
