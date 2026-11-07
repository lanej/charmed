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
  await expect(text).toContainText("Reference a13c9f2");
  await expect(text).toContainText("Candidate f4d7e81");
  await expect(text).toContainText("Member host-096");
  await expect(text).toContainText("Version f4d7e81");
  for (const pill of ["2xx", "3xx", "4xx", "5xx"]) {
    await expect(text).toContainText(pill);
  }
  await expect(text).not.toContainText("Fleet ");
  await expect(text).toContainText("+0.203pp");
  await expect(text).toContainText("delta +/-0.5pp");
  await expect(text).toContainText("CODE");
  await expect(text).toContainText("RATE");
  for (const legend of ["History %:", "3=member", "X = this member", "W/E =", "R=reference", "C=candidate", "1/2/3 bands"]) {
    await expect(text).not.toContainText(legend);
  }
  const viewport = await page.locator(".terminal-scroller").evaluate((element) => ({
    width: element.clientWidth, contentWidth: element.scrollWidth,
  }));
  expect(viewport.contentWidth).toBeLessThanOrEqual(viewport.width);
  await page.getByRole("button", { name: "Focus terminal", exact: true }).click();
  await page.keyboard.press("ArrowDown");
  await expect(page.getByTestId("terminal-events")).toContainText('"member":"host-095"');
  await expect(page.getByTestId("terminal-events")).toContainText('"cohort":"f4d7e81"');
  await expect(text).toContainText("Member host-095");
  await page.keyboard.press("ArrowLeft");
  await page.keyboard.press("2");
  await expect(page.getByTestId("terminal-events")).toContainText('"window":10');
  await expect(text).toContainText("+/- 2s host spread");
  await page.keyboard.press("Tab");
  await page.keyboard.press("Tab");
  await page.keyboard.press("ArrowUp");
  await expect(page.getByTestId("terminal-events")).toContainText('"metric":"4xx"');
  await expect(text).toContainText("Member host-095");
  await page.keyboard.press("ArrowDown");
  await page.keyboard.press("Control+]");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).fill("40");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).press("Enter");
  await expect(page.getByRole("status")).toContainText("40 × 40");
  await expect(text).toContainText("Member host-095");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).fill("132");
  await page.getByRole("spinbutton", { name: "Columns", exact: true }).press("Enter");
  await expect(text).toContainText("Member host-095");
  await expect(text).toContainText("+/- 2s host spread");
  await page.getByRole("combobox", { name: "Scenario", exact: true }).selectOption("insufficient-samples");
  await expect(text).toContainText("baseline s unavailable");
  await expect(text).toContainText("95 excluded");
  await expect(text).not.toContainText("NaN");
  await page.getByRole("combobox", { name: "Scenario", exact: true }).selectOption("isolated-hosts");
  await page.getByRole("combobox", { name: "Host spread bands", exact: true }).selectOption("3");
  await expect(text).toContainText("5 beyond 3s");
  await expect(text).toContainText("ERROR");
  await expect(text).toContainText("+0.203pp");
  await expect(text).toContainText("Warn>=+0.1 Error>=+0.3 pp");
  await page.getByRole("combobox", { name: "Deviation boundaries", exact: true }).selectOption("off");
  await expect(text).not.toContainText("ERROR");
  await expect(text).not.toContainText("ATTN");
  await expect(text).toContainText("5xx");
  await page.getByRole("combobox", { name: "Deviation boundaries", exact: true }).selectOption("example");
  await expect(text).toContainText("ERROR");
  await page.getByRole("button", { name: "Step", exact: true }).click();
  await expect(page.getByTestId("terminal-events")).toContainText('"type":"attention"');
  await expect(page.getByTestId("terminal-events")).toContainText('"previous":"error"');
  await expect(page.getByTestId("terminal-events")).toContainText('"level":"normal"');
  await page.getByRole("combobox", { name: "Preset", exact: true }).selectOption("Default");
  await expect(text).toContainText("ERROR");
  const reference = page.getByRole("textbox", { name: "Reference version", exact: true });
  const candidate = page.getByRole("textbox", { name: "Candidate version", exact: true });
  await reference.fill("release/2026.11.07+build.42");
  await reference.press("Enter");
  await expect(text).toContainText("Reference release/2026.11.07+build.42");
  await candidate.fill("sha256:731a9e4");
  await candidate.press("Enter");
  await expect(text).toContainText("Version sha256:731a9e4");
  await page.getByRole("button", { name: "Focus terminal", exact: true }).click();
  await page.keyboard.press("ArrowDown");
  await expect(page.getByTestId("terminal-events")).toContainText('"cohort":"sha256:731a9e4"');
  await expect(page.getByTestId("terminal-events")).toContainText('"reference":"release/2026.11.07+build.42"');
  await page.keyboard.press("Control+]");
  await page.getByRole("combobox", { name: "Preset", exact: true }).selectOption("Default");
  await expect(text).toContainText("Version f4d7e81");
  if (process.env.CHARMED_CAPTURE_DIR) {
    mkdirSync(process.env.CHARMED_CAPTURE_DIR, { recursive: true });
    await page.locator(".story-canvas").screenshot({ path: join(process.env.CHARMED_CAPTURE_DIR, "deployment.png") });
  }
});

test("static deployment variants retain the traffic and sampling scenarios", async ({ page }) => {
  await page.goto("http://127.0.0.1:5174/#screen/deployment");
  await expect(page.getByRole("status")).toContainText("Static snapshot: Default");
  await page.getByRole("combobox", { name: "Saved variant" }).selectOption("No boundaries");
  await expect(page.getByTestId("terminal-text")).not.toContainText("ERROR");
  await expect(page.getByTestId("terminal-text")).toContainText("5xx");
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

test("the reusable attention fixture distinguishes unknown and disabled rules", async ({ page }) => {
  await page.goto("/#feedback/attention");
  const text = page.getByTestId("terminal-text");
  await expect(text).toContainText(/^\s*WARN\s*$/m);
  await page.getByRole("slider", { name: "Amount", exact: true }).fill("85");
  await expect(text).toContainText(/^\s*ERROR\s*$/m);
  await page.getByRole("combobox", { name: "State", exact: true }).selectOption("empty");
  await expect(text).toContainText(/^\s*UNKNOWN\s*$/m);
  await page.getByRole("combobox", { name: "Deviation boundaries", exact: true }).selectOption("off");
  await expect(text).toContainText("OFF / no boundaries");
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
  await expect(navigation.getByRole("link")).toHaveCount(41);
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

test("nested static pages expose composition links in both directions", async ({ page }) => {
  await page.goto("http://127.0.0.1:5174/#screen/deployment");
  let composition = page.getByRole("region", { name: "Component composition", exact: true });
  await expect(composition.getByRole("heading", { name: "Composed from", exact: true })).toBeVisible();
  await composition.getByRole("link", { name: "Comparison panel", exact: true }).click();
  await expect(page).toHaveURL(/#chart\/comparison-panel$/);
  await expect(page.getByRole("heading", { name: "Comparison panel", exact: true })).toBeVisible();
  await expect(page.getByRole("status")).toContainText("Static snapshot: Default");
  composition = page.getByRole("region", { name: "Component composition", exact: true });
  await expect(composition.getByRole("link", { name: "Deployment comparison", exact: true })).toHaveAttribute("href", "#screen/deployment");
  await composition.getByRole("link", { name: "Cohort distributions", exact: true }).click();
  await expect(page).toHaveURL(/#chart\/distributions$/);
  composition = page.getByRole("region", { name: "Component composition", exact: true });
  await expect(composition.getByRole("link", { name: "Comparison panel", exact: true })).toBeVisible();
  await composition.getByRole("link", { name: "Interval band", exact: true }).click();
  await expect(page).toHaveURL(/#chart\/interval-band$/);
  await expect(page.getByRole("status")).toContainText("Static snapshot: Default");
  composition = page.getByRole("region", { name: "Component composition", exact: true });
  await expect(composition.getByRole("heading", { name: "Composed from", exact: true })).toHaveCount(0);
  await expect(composition.getByRole("link", { name: "Cohort distributions", exact: true })).toBeVisible();
  await expect(composition.getByRole("link", { name: "Merged delta property", exact: true })).toBeVisible();
  if (process.env.CHARMED_CAPTURE_DIR) {
    mkdirSync(process.env.CHARMED_CAPTURE_DIR, { recursive: true });
    await page.goto("http://127.0.0.1:5174/#chart/distributions");
    await expect(page.getByRole("region", { name: "Component composition", exact: true })).toBeVisible();
    await page.getByRole("region", { name: "Component composition", exact: true }).screenshot({ path: join(process.env.CHARMED_CAPTURE_DIR, "composition.png") });
  }
});
