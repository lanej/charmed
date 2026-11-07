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
  await expect(navigation.getByRole("link")).toHaveCount(34);
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
