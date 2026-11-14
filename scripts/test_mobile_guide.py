#!/usr/bin/env python3
"""Check the generated guide's responsive layout and navigation in Chromium.

Requires playwright==1.57.0 and its Chromium browser. The tested HTML, CSS, font,
and script are the actual build artifacts, loaded inline so CI needs no server
or external network. Run after `go run ./cmd/workbench generate --out docs`.
"""
import argparse
import base64
import json
import mimetypes
from pathlib import Path
import re
from urllib.parse import unquote, urlparse

from playwright.sync_api import sync_playwright


def document(site, filename):
    source = (site / filename).read_text(encoding="utf-8")
    styles = (site / "style.css").read_text(encoding="utf-8")
    styles += (site / "assets/mobile.css").read_text(encoding="utf-8")
    font = base64.b64encode((site / "assets/JuliaMono-Regular.woff2").read_bytes()).decode()
    styles = styles.replace('url("assets/JuliaMono-Regular.woff2")',
                            f'url("data:font/woff2;base64,{font}")')
    def inline_image(match):
        file = site / match[1]
        if file.is_file():
            mime = mimetypes.guess_type(file)[0] or 'application/octet-stream'
            data = base64.b64encode(file.read_bytes()).decode()
            return 'src="data:' + mime + ';base64,' + data + '"'
        return match[0]

    source = re.sub(r'src="(assets/[^"<>]+)"', inline_image, source)
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
    page.locator(".mobile-nav .nav-component").filter(has_text="Service workbench").tap()
    page.wait_for_timeout(100)
    assert not menu.evaluate("el => el.open"), "component navigation must close the menu"
    assert page.evaluate("document.activeElement.id") == "screen/services"
    assert 48 <= page.locator('[id="screen/services"]').bounding_box()["y"] < 150
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


def check_catalog_navigation(browser, site, output):
    # Route real generated files under the same prefix shape as a PR preview.
    base = "https://guide.test/charmed/pr-preview/pr-123/"

    def serve(route):
        path = unquote(urlparse(route.request.url).path).removeprefix("/charmed/pr-preview/pr-123/")
        file = site / path
        if file.is_file():
            route.fulfill(path=str(file))
        else:
            route.abort()

    context = browser.new_context(viewport={"width": 1440, "height": 1000})
    context.route("https://guide.test/**", serve)
    page = context.new_page()
    errors = []
    page.on("pageerror", lambda error: errors.append(str(error)))
    page.goto(base + "atoms.html")
    page.locator('.group-card[href="atoms-chart.html"]').click()
    page.wait_for_url(base + "atoms-chart.html")
    assert page.locator('.desktop-nav [aria-current="page"]').inner_text() == "Charts"
    assert page.locator(".component").count() >= 10
    page.locator('.desktop-nav .nav-level[href="molecules.html"]').click()
    page.locator('.group-card[href="molecules-commands.html"]').click()
    page.wait_for_url(base + "molecules-commands.html")
    page.locator('.dependency').filter(has_text="input/text").click()
    page.wait_for_url(base + "atoms-input.html#input/text")
    assert page.locator('[id="input/text"]').is_visible()

    redirects = {
        "components.html#input/text": "atoms-input.html#input/text",
        "components.html#input%2Ftext-code": "atoms-input.html#input/text-code",
        "components.html#level-molecules": "molecules.html#level-molecules",
        "components.html#level-molecules-chart": "molecules-chart.html#level-molecules-chart",
        "components.html#family-chart": "atoms-chart.html#level-atoms-chart",
        "examples.html#screen/deployment-code": "screens-pattern.html#screen/deployment-code",
        "examples.html#template/application": "templates-pattern.html#template/application",
        "foundations.html#foundations/theme": "foundations-theme.html#foundations/theme",
        "components.html#panel-chart/sparkline-1": "atoms-chart.html#panel-chart/sparkline-1",
        "components.html#tab-chart/sparkline-2": "atoms-chart.html#tab-chart/sparkline-2",
    }
    for old, new in redirects.items():
        page.goto(base + old)
        page.wait_for_url(base + new)
        target = new.split("#", 1)[1]
        assert page.locator(f'[id="{target}"]').is_visible(), (old, new)
        if target.startswith("tab-"):
            assert page.locator(f'[id="{target}"]').get_attribute("aria-selected") == "true"
    page.goto(base + "components.html")
    page.evaluate("location.hash = 'feedback/status'")
    page.wait_for_url(base + "atoms-feedback.html#feedback/status")
    assert not errors, errors
    page.goto(base + "atoms.html")
    page.screenshot(path=str(output / "category-desktop.png"))
    context.close()

    context = browser.new_context(viewport={"width": 390, "height": 844},
                                  has_touch=True, is_mobile=True)
    context.route("https://guide.test/**", serve)
    page = context.new_page()
    page.goto(base + "atoms.html")
    page.locator(".mobile-nav > summary").tap()
    page.locator('.mobile-nav .nav-group[href="atoms-text.html"]').tap()
    page.wait_for_url(base + "atoms-text.html")
    assert not page.locator("[data-mobile-nav]").evaluate("el => el.open")
    page.locator(".mobile-nav > summary").tap()
    page.locator('.mobile-nav .nav-level[href="molecules.html"]').tap()
    page.wait_for_url(base + "molecules.html")
    assert page.locator("h1 .heading-permalink > span").first.inner_text() == "Molecules"
    context.close()

    context = browser.new_context(java_script_enabled=False)
    context.route("https://guide.test/**", serve)
    page = context.new_page()
    page.goto(base + "components.html#input/text")
    page.locator(".legacy-links summary").click()
    page.locator('[data-legacy-anchor="input/text"]').click()
    page.wait_for_url(base + "atoms-input.html#input/text")
    assert page.locator('[id="input/text"]').is_visible()
    context.close()


def check_rhythm(page, filename, config):
    measurements = page.evaluate("""() => {
        const box = node => node.getBoundingClientRect();
        const components = Array.from(document.querySelectorAll('.component')).map(component => {
            const heading = component.querySelector('h2');
            const preview = component.querySelector('[id$="-example"][role="tabpanel"] .terminal');
            const source = component.querySelector('.source-code');
            const integration = component.querySelector('.component-integration');
            const run = component.querySelector('.component-run');
            const canvas = document.createElement('canvas').getContext('2d');
            canvas.font = getComputedStyle(preview).font;
            const cell = canvas.measureText('0').width;
            const columns = Math.max(...preview.textContent.split('\\n').map(line => canvas.measureText(line).width / cell));
            return {id: component.id,
                firstArtifact: box(preview).top - box(heading).bottom,
                previewHeight: box(preview).height,
                columns: Math.round(columns),
                artifactGap: box(source).top - box(preview).bottom,
                integrationGap: box(integration).top - box(source).bottom,
                runGap: box(run).top - box(integration).bottom,
                clippedCode: source.scrollHeight > source.clientHeight + 1};
        });
        const prose = Array.from(document.querySelectorAll('main > p, .component-use, .component-integration p, .catalog-description, .dev-lede')).filter(node =>
            !node.matches('.eyebrow, .catalog-count, .catalog-breadcrumb, .example-imports, .composition-line')
        ).map(node => {
            const canvas = document.createElement('canvas').getContext('2d');
            canvas.font = getComputedStyle(node).font;
            const cell = canvas.measureText('0').width;
            return {text: node.textContent.slice(0, 45), width: box(node).width / cell,
                available: box(node.parentElement).width / cell};
        });
        const sectionGaps = Array.from(document.querySelectorAll('main > h2')).map(node => ({id: node.id, gap: parseFloat(getComputedStyle(node).marginTop)}));
        return {components, prose, sectionGaps};
    }""")
    limits = config['rendered']
    for item in measurements['components']:
        checks = {'firstArtifact': 'first_artifact_px', 'previewHeight': 'preview_height_px',
                  'artifactGap': 'artifact_gap_px', 'integrationGap': 'section_gap_px', 'runGap': 'section_gap_px'}
        for metric, rule in checks.items():
            low, high = limits[rule]
            assert low - 1 <= item[metric] <= high + 1, (filename, item['id'], rule, item[metric], limits[rule])
        assert not item['clippedCode'], (filename, item['id'], 'usage code must be fully visible vertically')
        low, high = config['component']['preview_columns']
        assert low <= item['columns'] <= high, (filename, item['id'], 'preview columns', item['columns'])
    for item in measurements['sectionGaps']:
        low, high = limits['section_gap_px']
        assert low <= item['gap'] <= high, (filename, item['id'], 'section spacing', item)
    for item in measurements['prose']:
        low, high = limits['prose_width_ch']
        assert item['width'] <= high + 1, (filename, 'prose too wide', item)
        if item['available'] >= low:
            assert item['width'] >= low - 1, (filename, 'prose too narrow', item)
    return measurements


def check_component_graph(browser, site, output):
    base = "https://guide.test/charmed/pr-preview/pr-123/"
    document = json.loads((site / "component-graph.json").read_text())
    assert document["version"] == 1
    assert not any(edge["from"] == "diagram/trace-replay" and edge["to"] == "input/radio-group" for edge in document["relations"])
    assert any(example["kind"] == "interactive" and any(use["component"] == "input/radio-group" for use in example["uses"]) for example in document["examples"])

    def serve(route):
        path = unquote(urlparse(route.request.url).path).removeprefix("/charmed/pr-preview/pr-123/")
        file = site / path
        if file.is_file():
            route.fulfill(path=str(file))
        else:
            route.abort()

    for width in (1440, 390):
        context = browser.new_context(viewport={"width": width, "height": 1000})
        try:
            context.route("https://guide.test/**", serve)
            page = context.new_page()
            errors = []
            page.on("pageerror", lambda error: errors.append(str(error)))
            page.goto(base + "dependencies.html#component-collection%2Fstatus-matrix")
            cards = page.locator("[data-component]")
            selector = page.locator("#graph-component")
            assert selector.locator("option").count() == len(document["components"])
            assert page.locator("[data-component]:visible").count() == 1
            assert page.locator('[data-component="collection/status-matrix"]').is_visible()
            for option in selector.locator("option").all():
                selector.select_option(option.get_attribute("value"))
                preview = page.locator("[data-component]:visible").get_by_role("region", name="Native example", exact=True)
                assert preview.locator("pre.terminal").inner_text().strip(), "Missing compiled consumer output"
                assert preview.get_by_role("link", name="Go usage").get_attribute("href")
                if option.get_attribute("value") == "feedback/canary-membership":
                    page.evaluate("document.fonts.ready")
                    positions = preview.locator("pre.terminal").evaluate("""el => {
                        const positions = [];
                        const nodes = document.createTreeWalker(el, NodeFilter.SHOW_TEXT);
                        let node;
                        while (node = nodes.nextNode()) {
                            for (const glyph of ['◉', '◌']) {
                                const at = node.textContent.indexOf(glyph);
                                if (at < 0) continue;
                                const range = document.createRange();
                                range.setStart(node, at);
                                range.setEnd(node, at + 1);
                                positions.push(range.getBoundingClientRect().x);
                            }
                        }
                        return positions;
                    }""")
                    assert len(positions) == 2 and abs(positions[0] - positions[1]) <= 1, "Emoji fallback fonts shifted the compact column"
                    fills = preview.locator('pre.terminal span[style*="background-color"]').evaluate_all("nodes => nodes.map(el => el.getBoundingClientRect().height)")
                    assert fills and max(fills) - min(fills) <= 1, "Emoji cell wrappers changed the pill fill height"
            selector.select_option("collection/status-matrix")
            selected = page.locator("[data-component]:visible")
            selected.get_by_role("region", name="Built from", exact=True).get_by_role("link", name="Data grid", exact=True).click()
            assert page.locator('[data-component="collection/data-grid"]').is_visible()
            selected = page.locator("[data-component]:visible")
            selected.get_by_role("region", name="Used by", exact=True).get_by_role("link", name="Grid inspector", exact=True).click()
            assert page.locator('[data-component="collection/grid-inspector"]').is_visible()
            page.go_back()
            page.locator('[data-component="collection/data-grid"]').wait_for(state="visible")
            assert page.locator(":focus").inner_text() == "Data grid"
            assert page.locator(":focus").evaluate("el => el.getBoundingClientRect().top >= 0 && el.getBoundingClientRect().bottom <= innerHeight")
            selector.select_option("input/radio-group")
            usage = page.locator("[data-component]:visible").get_by_role("region", name="Example usage", exact=True)
            usage.get_by_role("link", name="Trace replay", exact=True).last.click()
            assert page.locator(":focus").inner_text() == "Trace replay"
            assert page.locator(":focus").evaluate("el => el.getBoundingClientRect().top >= 0 && el.getBoundingClientRect().bottom <= innerHeight")
            page.go_back()
            page.locator('[data-component="input/radio-group"]').wait_for(state="visible")
            assert page.locator(":focus").inner_text() == "Radio group"
            selector.select_option("diagram/trace-replay")
            selected = page.locator("[data-component]:visible")
            built = selected.get_by_role("region", name="Built from", exact=True)
            assert built.get_by_role("link", name="State diagram", exact=True).count() == 1
            assert built.get_by_role("link", name="Radio group", exact=True).count() == 0
            assert selected.get_by_role("region", name="Example usage", exact=True).get_by_role("link", name="Radio group", exact=True).count() == 2
            selector.focus()
            page.keyboard.press("Home")
            page.keyboard.press("ArrowDown")
            page.keyboard.press("Enter")
            page.keyboard.press("Escape")
            assert page.locator("[data-component]:visible").get_attribute("data-component") == selector.input_value()
            assert page.locator(":focus").get_attribute("id") == "graph-component"
            page.goto(base + "dependencies.html#component-missing")
            assert "Unknown component" in page.locator("[data-graph-status]").inner_text()
            assert page.locator('[data-component="collection/status-matrix"]').is_visible()
            assert not errors, errors
            page.screenshot(path=str(output / f"dependencies-{width}.png"), full_page=True)
        finally:
            context.close()

    context = browser.new_context(java_script_enabled=False)
    try:
        context.route("https://guide.test/**", serve)
        page = context.new_page()
        page.goto(base + "dependencies.html#component-collection%2Fdata-grid")
        assert page.locator("[data-component]:visible").count() == len(document["components"])
        assert page.locator(".graph-picker").is_hidden()
        assert page.locator('[id="component-collection/data-grid"]').is_visible()
    finally:
        context.close()


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--site", type=Path, default=Path("docs/dist"))
    parser.add_argument("--output", type=Path, default=Path("artifacts/mobile-guide"))
    parser.add_argument("--chromium-path")
    args = parser.parse_args()
    args.output.mkdir(parents=True, exist_ok=True)
    pages = sorted(args.site.glob("*.html"))
    assert len(pages) > 12, "build the complete guide before running browser checks"
    sources = {file.name: document(args.site, file.name) for file in pages}
    results = []
    config = json.loads((Path(__file__).resolve().parents[1] / 'docs/standards.json').read_text())
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
                page.wait_for_function("() => Array.from(document.images).every(image => image.complete && image.naturalWidth > 0)")
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
                if filename in ("atoms.html", "atoms-chart.html", "molecules.html", "screens-pattern.html", "development.html") and width in (390, 1440):
                    page.screenshot(path=str(args.output / f"{filename.removesuffix('.html')}-{width}.png"))
                expected_levels = ["Foundations", "Atoms", "Molecules", "Organisms", "Templates", "Screens"]
                actual_levels = page.locator(".desktop-nav .nav-level").evaluate_all("els => els.map(el => el.firstChild.textContent.trim())")
                assert actual_levels == expected_levels, (filename, actual_levels)
                if filename in [level.lower() + ".html" for level in expected_levels]:
                    assert page.locator(".component").count() == 0, "category pages must stay lightweight"
                    assert page.locator(".group-card").count() > 0
                    assert page.locator(".group-card details, .group-card summary, .group-card button").count() == 0
                    specimens = page.locator(".group-specimen pre")
                    assert specimens.count() == page.locator(".group-card").count()
                    assert all(text.strip() for text in specimens.all_text_contents()), "every group needs a native example"
                    assert specimens.evaluate_all("""els => els.every(el => {
                        const sample = el.getBoundingClientRect();
                        const card = el.parentElement.getBoundingClientRect();
                        return el.scrollWidth <= el.clientWidth + 1 &&
                            sample.left >= card.left && sample.right <= card.right &&
                            sample.top >= card.top && sample.bottom <= card.bottom;
                    })"""), (filename, width, "clipped specimen")
                    assert page.locator('.desktop-nav [aria-current="page"]').count() == 1
                if filename.startswith("molecules-"):
                    molecules = page.locator(".component")
                    assert molecules.count() > 0, "molecule groups must not be empty"
                    assert molecules.locator(".component-source h3").all_text_contents() == ["Go usage#"] * molecules.count()
                for component in page.locator(".component").all():
                    example = component.locator('[role="tab"]').first
                    assert example.inner_text() == "Example"
                    assert example.get_attribute("aria-selected") == "true"
                    assert component.locator('[id$="-example"][role="tabpanel"]').is_visible()
                    code = component.locator(".source-code").inner_text()
                    assert not any(internal in code for internal in ("e.Config", "sampleSeries(", "package snippets", "func Example"))
                if filename == "development.html":
                    tasks = page.locator(".dev-paths a")
                    assert tasks.count() == 3, "developer guide needs three task paths"
                    for task in tasks.all():
                        target = task.get_attribute("href")
                        assert target.startswith("#") and page.locator(f'[id="{target[1:]}"]').count() == 1, target
                if filename == "screens-pattern.html" and width == 390:
                    check_menu(page, args.output)
                rhythm = check_rhythm(page, filename, config)
                results.append({"page": filename, **values, "rhythm": rhythm})
                context.close()
            print(f"Verified {width}px layouts", flush=True)
        context = browser.new_context(viewport={"width": 390, "height": 844}, java_script_enabled=False)
        page = context.new_page()
        page.set_content(sources["screens-pattern.html"][0])
        menu = page.locator("[data-mobile-nav]")
        assert menu.get_attribute("open") is None
        page.locator(".mobile-nav summary").click()
        assert menu.get_attribute("open") is not None, "native menu must work without JavaScript"
        context.close()
        check_catalog_navigation(browser, args.site, args.output)
        check_component_graph(browser, args.site, args.output)
        browser.close()
    (args.output / "results.json").write_text(json.dumps(results, indent=2), encoding="utf-8")
    print(f"Verified {len(results)} page/width layouts, touch, keyboard, fragments, resize, and no-JS menu.")


if __name__ == "__main__":
    main()
