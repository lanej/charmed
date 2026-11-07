import { createRoot } from "react-dom/client";
import { App } from "./App";
import "./tokens.css";
import "../components/workbench.css";
import "./catalog.css";

createRoot(document.getElementById("root")!).render(<App />);
