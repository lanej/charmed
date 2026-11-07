import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import mdx from "@mdx-js/rollup";

export default defineConfig({
  base: "./",
  plugins: [
    { ...mdx(), enforce: "pre" },
    react({ include: /\.(jsx|tsx|mdx)$/ }),
  ],
  server: {
    host: "127.0.0.1",
    port: 5173,
    strictPort: true,
    proxy: { "/api": { target: "http://127.0.0.1:7331", ws: true } },
  },
});
