import { spawn, spawnSync } from "node:child_process";
import { mkdirSync } from "node:fs";
import { resolve } from "node:path";
import { fileURLToPath } from "node:url";

const docs = fileURLToPath(new URL("..", import.meta.url));
const root = resolve(docs, "..");
const binary = resolve(root, ".cache/workbench");
mkdirSync(resolve(root, ".cache"), { recursive: true });
for (const args of [["build", "-o", binary, "./cmd/workbench"]]) {
  const result = spawnSync("go", args, { cwd: root, stdio: "inherit" });
  if (result.status !== 0) process.exit(result.status ?? 1);
}
const generated = spawnSync(binary, ["generate", "--out", docs], {
  cwd: root,
  stdio: "inherit",
});
if (generated.status !== 0) process.exit(generated.status ?? 1);
const backend = spawn(binary, ["serve"], { cwd: root, stdio: "inherit" });
const frontend = spawn(
  process.execPath,
  [
    resolve(docs, "node_modules/vite/bin/vite.js"),
    "--port",
    "5173",
    "--host",
    "127.0.0.1",
  ],
  { cwd: docs, stdio: "inherit" },
);
let closing = false;
function stop(code = 0) {
  if (closing) return;
  closing = true;
  backend.kill("SIGTERM");
  frontend.kill("SIGTERM");
  process.exitCode = code;
}
process.on("SIGINT", () => stop());
process.on("SIGTERM", () => stop());
backend.on("exit", (code) => stop(code ?? 1));
frontend.on("exit", (code) => stop(code ?? 1));
