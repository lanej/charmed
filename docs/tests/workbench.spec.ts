import { test, expect } from "@playwright/test";
import { mkdirSync } from "node:fs";
import { join } from "node:path";

test("native input accepts q as text and emits a Go submission", async ({
  page,
}) => {
  await page.goto("/#input/text");
  await expect(page.getByRole("status")).toContainText("isolated PTY");
  await page
    .getByRole("button", { name: "Focus terminal", exact: true })
    .click();
  await page.keyboard.press("Control+u");
  await page.keyboard.type("q");
  await page.keyboard.press("Enter");
  await expect(page.getByTestId("terminal-events")).toContainText(
    '"value":"q"',
  );
  await expect(page.getByRole("status")).toContainText("isolated PTY");
  await page.getByRole("button", { name: "Reset", exact: true }).click();
  await expect(page.getByTestId("terminal-events")).not.toContainText(
    '"value":"q"',
  );
});

test("the service screen coordinates selection and background steps", async ({
  page,
}) => {
  await page.goto("/#screen/services");
  await expect(page.getByRole("status")).toContainText("isolated PTY");
  await expect(page.getByTestId("terminal-text")).toContainText("API gateway");
  await page
    .getByRole("button", { name: "Focus terminal", exact: true })
    .click();
  await page.keyboard.press("Tab");
  await page.keyboard.press("ArrowDown");
  await expect(page.getByTestId("terminal-events")).toContainText(
    '"id":"events"',
  );
  await expect(page.getByTestId("terminal-text")).toContainText(
    "Event stream / throughput",
  );
  await page.getByRole("button", { name: "Step", exact: true }).click();
  await expect(page.getByTestId("terminal-events")).toContainText(
    '"selected":"events"',
  );
  await expect(page.getByTestId("terminal-text")).toContainText(
    "Background samples: 1",
  );
  await page
    .getByRole("spinbutton", { name: "Columns", exact: true })
    .fill("40");
  await page
    .getByRole("spinbutton", { name: "Columns", exact: true })
    .press("Enter");
  await expect(page.getByRole("status")).toContainText("40 × 34");
  await expect(page.getByTestId("terminal-text")).toContainText(
    "Event stream / throughput",
  );
  await page
    .getByRole("spinbutton", { name: "Columns", exact: true })
    .fill("120");
  await page
    .getByRole("spinbutton", { name: "Columns", exact: true })
    .press("Enter");
  await expect(page.getByTestId("terminal-text")).toContainText(
    "Stable selection: events",
  );
  if (process.env.CHARMED_CAPTURE_DIR) {
    mkdirSync(process.env.CHARMED_CAPTURE_DIR, { recursive: true });
    await page.screenshot({
      path: join(process.env.CHARMED_CAPTURE_DIR, "workbench.png"),
    });
  }
  await page
    .getByRole("button", { name: "Focus terminal", exact: true })
    .click();
  await page.keyboard.press("Control+]");
  await expect(
    page.getByRole("button", { name: "Release focus", exact: true }),
  ).toBeFocused();
});

test("MDX pages include highlighted Go and file context", async ({ page }) => {
  await page.goto("/#chart/timeseries");
  await expect(
    page.getByRole("heading", { name: /Time series$/, level: 1 }),
  ).toBeVisible();
  await expect(page.locator(".go-source .shiki")).toBeVisible();
  await expect(page.locator(".source-heading a")).toHaveText(
    "examples/timeseries/main.go",
  );
  await expect(
    page.locator(".go-source .shiki span[style]").first(),
  ).toBeVisible();
});

test("deployment comparison links host selection, windows, bands and sparse fixtures", async ({ page }) => {
  await page.setViewportSize({ width: 1760, height: 1250 });
  await page.goto("/#screen/deployment");
  await expect(page.getByRole("status")).toContainText("isolated PTY");
  const text = page.getByTestId("terminal-text");
  await expect(text).toContainText("method=blue/green");
  await expect(text).toContainText("R a13c9f2");
  await expect(text).toContainText("C f4d7e81");
  await expect(text).toContainText("host-096 / C");
  await expect(text).toContainText("ID: f4d7e81");
  for (const pill of ["2xx Success", "3xx Redirects", "4xx Client error", "5xx Server error"]) {
    await expect(text).toContainText(pill);
  }
  await expect(text).toContainText("Fleet R 96.093%  C 95.904%");
  await expect(text).toContainText("Fleet R 0.200%  C 0.403%");
  const viewport = await page.locator(".terminal-scroller").evaluate((element) => ({
    width: element.clientWidth, contentWidth: element.scrollWidth,
  }));
  expect(viewport.contentWidth).toBeLessThanOrEqual(viewport.width);
  await page.getByRole("button", { name: "Focus terminal", exact: true }).click();
  await page.keyboard.press("ArrowDown");
  await expect(page.getByTestId("terminal-events")).toContainText('"member":"host-095"');
  await expect(page.getByTestId("terminal-events")).toContainText('"cohort":"f4d7e81"');
  await expect(text).toContainText("host-095 / C");
  await page.keyboard.press("ArrowLeft");
  await page.keyboard.press("2");
  await expect(page.getByTestId("terminal-events")).toContainText('"window":10');
  await expect(text).toContainText("+/- 2s host spread");
  await page.keyboard.press("Control+]");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).fill("40");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).press("Enter");
  await expect(page.getByRole("status")).toContainText("40 × 40");
  await expect(text).toContainText("host-095 / C");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).fill("132");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).press("Enter");
  await expect(text).toContainText("host-095 / C");
  await expect(text).toContainText("+/- 2s host spread");
  await page.getByRole("combobox", { name: "Scenario", exact: true }).selectOption("insufficient-samples");
  await expect(text).toContainText("baseline s unavailable");
  await expect(text).toContainText("95 excluded");
  await expect(text).not.toContainText("NaN");
  await page.getByRole("combobox", { name: "Scenario", exact: true }).selectOption("isolated-hosts");
  await page.getByRole("combobox", { name: "Host spread bands", exact: true }).selectOption("3");
  await expect(text).toContainText("5 beyond 3s");
  const reference = page.getByRole("textbox", { name: "Reference version", exact: true });
  const candidate = page.getByRole("textbox", { name: "Candidate version", exact: true });
  await reference.fill("release/2026.11.07+build.42");
  await reference.press("Enter");
  await expect(text).toContainText("R release/2026.11.07+build.42");
  await candidate.fill("sha256:731a9e4");
  await candidate.press("Enter");
  await expect(text).toContainText("ID: sha256:731a9e4");
  await page.getByRole("button", { name: "Focus terminal", exact: true }).click();
  await page.keyboard.press("ArrowDown");
  await expect(page.getByTestId("terminal-events")).toContainText('"cohort":"sha256:731a9e4"');
  await expect(page.getByTestId("terminal-events")).toContainText('"reference":"release/2026.11.07+build.42"');
  await page.keyboard.press("Control+]");
  await page.getByRole("combobox", { name: "Preset", exact: true }).selectOption("Default");
  await expect(text).toContainText("ID: f4d7e81");
  if (process.env.CHARMED_CAPTURE_DIR) {
    mkdirSync(process.env.CHARMED_CAPTURE_DIR, { recursive: true });
    await page.locator(".story-canvas").screenshot({ path: join(process.env.CHARMED_CAPTURE_DIR, "deployment.png") });
  }
});

test("static deployment variants retain the traffic and sampling scenarios", async ({ page }) => {
  await page.goto("http://127.0.0.1:5174/#screen/deployment");
  await expect(page.getByRole("status")).toContainText("Static snapshot: Default");
  await page.getByRole("combobox", { name: "Saved variant" }).selectOption("Insufficient samples");
  await expect(page.getByRole("status")).toContainText("Static snapshot: Insufficient samples");
  await expect(page.getByTestId("terminal-text")).toContainText("baseline s unavailable");
  await page.getByRole("combobox", { name: "Saved variant" }).selectOption("Unequal traffic");
  await expect(page.getByTestId("terminal-text")).toContainText("unequal-traffic");
});

test("static export labels snapshots and offers saved variants", async ({
  page,
}) => {
  await page.goto("http://127.0.0.1:5174/#screen/services");
  await expect(page.getByRole("status")).toContainText(
    "Static snapshot: Default",
  );
  await expect(
    page.getByRole("button", { name: "Step", exact: true }),
  ).toBeDisabled();
  await page
    .getByRole("combobox", { name: "Saved variant" })
    .selectOption("Narrow");
  await expect(page.getByRole("status")).toContainText(
    "Static snapshot: Narrow",
  );
  await expect(page.getByTestId("terminal-text")).toContainText("CHARMED");
});

test("catalog navigation supports search, capability groups and mobile", async ({
  page,
}) => {
  await page.goto("/");
  await expect(page.getByRole("heading", { level: 1 })).toContainText(
    "Charmed",
  );
  const navigation = page.getByRole("navigation", {
    name: "Components",
    exact: true,
  });
  await expect(navigation.getByRole("link")).toHaveCount(38);
  await page.getByPlaceholder("Search components…").fill("heatmap");
  await expect(navigation.getByRole("link")).toHaveCount(1);
  await navigation.getByRole("link").click();
  await expect(
    page.getByRole("heading", { name: "Heatmap", exact: true }),
  ).toBeVisible();
  await page.getByRole("button", { name: "Capability", exact: true }).click();
  await expect(
    page.getByRole("button", { name: "Capability", exact: true }),
  ).toHaveAttribute("aria-pressed", "true");
  await page.setViewportSize({ width: 390, height: 844 });
  await page.getByRole("button", { name: "Components", exact: true }).click();
  await expect(page.getByPlaceholder("Search components…")).toBeVisible();
});

test("radar links smooth profiles, host selection and the detail view", async ({ page }) => {
  await page.setViewportSize({ width: 1760, height: 1400 });
  await page.goto("/#screen/deployment-radar");
  const text = page.getByTestId("terminal-text");
  const events = page.getByTestId("terminal-events");
  await expect(page.getByRole("status")).toContainText("isolated PTY");
  await expect(text).toContainText("Version profile / smooth");
  await expect(text).toContainText("0 = reference mean");
  await expect(text).toContainText("Fleet R 0.200%  C 0.403%");
  await expect(page.getByRole("combobox", { name: "Profile curve", exact: true })).toHaveValue("smooth");
  await page.getByRole("button", { name: "Focus terminal", exact: true }).click();
  await page.keyboard.press("ArrowDown");
  await expect(events).toContainText('"member":"host-095"');
  await expect(events).toContainText('"radar":true');
  await page.keyboard.press("s");
  await expect(text).toContainText("Version profile / straight");
  await expect(events).toContainText('"smooth":false');
  await page.keyboard.press("r");
  await expect(text).not.toContainText("Version profile");
  await expect(text).toContainText("host-095 / C");
  await page.keyboard.press("r");
  await expect(text).toContainText("Version profile / straight");
  await page.keyboard.press("]");
  await expect(events).toContainText('"metric":"2xx"');
  await page.keyboard.press("ArrowLeft");
  await page.keyboard.press("1");
  await expect(events).toContainText('"window":10');
  await page.keyboard.press("Control+]");
  const columns = page.getByRole("spinbutton", { name: "Columns", exact: true });
  await columns.fill("40");
  await columns.press("Enter");
  await expect(page.getByRole("status")).toContainText("40 × 39");
  await columns.fill("132");
  await columns.press("Enter");
  await expect(text).toContainText("Version profile / straight");
  await expect(text).toContainText("host-095 / C");
  await expect(text).toContainText("+/- 1s host spread");
  await page.getByRole("combobox", { name: "Scenario", exact: true }).selectOption("insufficient-samples");
  await expect(text).toContainText("Reference spread unavailable");
  await expect(text).not.toContainText("NaN");
  await page.getByRole("combobox", { name: "Preset", exact: true }).selectOption("Default");
  await expect(text).toContainText("Version profile / smooth");
  const viewport = await page.locator(".terminal-scroller").evaluate((element) => ({
    width: element.clientWidth, contentWidth: element.scrollWidth,
    height: element.clientHeight, contentHeight: element.scrollHeight,
  }));
  expect(viewport.contentWidth).toBeLessThanOrEqual(viewport.width);
  expect(viewport.contentHeight).toBeLessThanOrEqual(viewport.height);
  if (process.env.CHARMED_CAPTURE_DIR) {
    mkdirSync(process.env.CHARMED_CAPTURE_DIR, { recursive: true });
    await page.locator(".story-canvas").screenshot({ path: join(process.env.CHARMED_CAPTURE_DIR, "deployment-radar.png") });
  }
});

test("static radar includes straight and sparse variants", async ({ page }) => {
  await page.goto("http://127.0.0.1:5174/#screen/deployment-radar");
  const text = page.getByTestId("terminal-text");
  await expect(page.getByRole("status")).toContainText("Static snapshot: Default");
  await expect(text).toContainText("Version profile / smooth");
  await page.getByRole("combobox", { name: "Saved variant" }).selectOption("Straight");
  await expect(text).toContainText("Version profile / straight");
  await page.getByRole("combobox", { name: "Saved variant" }).selectOption("Insufficient samples");
  await expect(text).toContainText("Reference spread unavailable");
});
