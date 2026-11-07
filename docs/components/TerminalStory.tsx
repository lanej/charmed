import { useEffect, useRef, useState } from "react";
import type { Terminal } from "@xterm/xterm";
import "@xterm/xterm/css/xterm.css";
import manifest from "../generated/manifest.json";

export type StoryArgs = (typeof manifest)[number]["defaults"] & {
  storyId: string;
};
type PreviewMode = "connecting" | "live" | "static" | "stopped";

function encode(value: unknown) {
  return btoa(
    String.fromCharCode(...new TextEncoder().encode(JSON.stringify(value))),
  )
    .replaceAll("+", "-")
    .replaceAll("/", "_")
    .replaceAll("=", "");
}
function colors(theme: string) {
  return theme === "light"
    ? {
        background: "#f8fafc",
        foreground: "#172033",
        cursor: "#075985",
        selectionBackground: "#cbd5e1",
      }
    : {
        background: theme === "contrast" ? "#000000" : "#101722",
        foreground: "#e2e8f0",
        cursor: "#67d4c0",
        selectionBackground: "#334155",
      };
}

export function TerminalStory(props: StoryArgs) {
  const entry = manifest.find((entry) => entry.id === props.storyId);
  const host = useRef<HTMLDivElement>(null);
  const shell = useRef<HTMLDivElement>(null);
  const release = useRef<HTMLButtonElement>(null);
  const terminal = useRef<Terminal | null>(null);
  const socket = useRef<WebSocket | null>(null);
  const paintSnapshot = useRef<
    ((snapshot: (typeof manifest)[number]["snapshots"][number]) => void) | null
  >(null);
  const [mode, setMode] = useState<PreviewMode>("connecting");
  const [reset, setReset] = useState(0);
  const [savedVariant, setSavedVariant] = useState<string | null>(null);
  const [readable, setReadable] = useState("");
  const [events, setEvents] = useState<unknown[]>([]);
  const { storyId, columns, rows, ...rest } = props;
  const configurationKey = JSON.stringify(rest);
  const automaticSnapshot = entry?.snapshots.find(
    (snapshot) =>
      snapshot.config.theme === props.theme &&
      snapshot.config.columns === columns &&
      snapshot.config.rows === rows,
  );
  const snapshot =
    entry?.snapshots.find((snapshot) => snapshot.name === savedVariant) ??
    automaticSnapshot ??
    entry?.snapshots[0];
  const validSize =
    Number.isInteger(columns) &&
    columns >= 20 &&
    columns <= 240 &&
    Number.isInteger(rows) &&
    rows >= 8 &&
    rows <= 80;

  useEffect(
    () => setSavedVariant(null),
    [storyId, configurationKey, columns, rows],
  );

  useEffect(() => {
    if (!host.current || !entry || !snapshot || !validSize) return;
    let disposed = false;
    let term: Terminal | undefined;
    let ws: WebSocket | undefined;
    let refresh = 0;
    const controller = new AbortController();
    setEvents([]);
    setMode("connecting");

    const start = async () => {
      const { Terminal } = await import("@xterm/xterm");
      if (disposed || !host.current) return;
      term = new Terminal({
        cols: snapshot.config.columns,
        rows: snapshot.config.rows,
        fontFamily: '"SFMono-Regular", Consolas, "Liberation Mono", monospace',
        fontSize: 13,
        lineHeight: 1.2,
        cursorBlink: false,
        convertEol: true,
        scrollback: 0,
        theme: colors(snapshot.config.theme),
      });
      terminal.current = term;
      term.open(host.current);
      host.current
        .querySelector("textarea")
        ?.setAttribute("aria-label", "Interactive terminal input");
      const readBuffer = () => {
        if (!term || disposed) return;
        const buffer = term.buffer.active;
        const lines = Array.from(
          { length: term.rows },
          (_, row) =>
            buffer.getLine(buffer.baseY + row)?.translateToString(true) ?? "",
        );
        setReadable(lines.join("\n").trimEnd());
      };
      const render = (data: string | Uint8Array) =>
        term?.write(data, () => {
          cancelAnimationFrame(refresh);
          refresh = requestAnimationFrame(readBuffer);
        });
      term.attachCustomKeyEventHandler((event) => {
        if (event.ctrlKey && event.key === "]") {
          term?.blur();
          release.current?.focus();
          return false;
        }
        return true;
      });
      paintSnapshot.current = (saved) => {
        term?.reset();
        term?.resize(saved.config.columns, saved.config.rows);
        if (term) term.options.theme = colors(saved.config.theme);
        render(saved.frame);
      };
      paintSnapshot.current(snapshot);
      let available = false;
      let socketURL = "/api/session";
      try {
        const response = await fetch("/api/health", {
          signal: AbortSignal.any([
            controller.signal,
            AbortSignal.timeout(1500),
          ]),
        });
        if (response.ok) {
          const health = await response.json();
          available = health.status === "ok";
          socketURL = health.socketURL ?? socketURL;
        }
      } catch {
        /* A static export intentionally has no preview service. */
      }
      if (disposed) return;
      if (!available) {
        setMode("static");
        return;
      }

      const config = { ...rest, columns, rows };
      term.resize(columns, rows);
      term.options.theme = colors(props.theme);
      term.clear();
      term.reset();
      const url = new URL(socketURL, window.location.href);
      url.protocol =
        url.protocol === "https:" || url.protocol === "wss:" ? "wss:" : "ws:";
      url.searchParams.set("id", storyId);
      url.searchParams.set("config", encode(config));
      ws = new WebSocket(url);
      ws.binaryType = "arraybuffer";
      socket.current = ws;
      ws.onopen = () => {
        if (!disposed) setMode("live");
      };
      ws.onmessage = (message) => {
        if (disposed) return;
        if (message.data instanceof ArrayBuffer)
          render(new Uint8Array(message.data));
        else {
          try {
            const value = JSON.parse(message.data);
            if (value.type === "event")
              setEvents((events) => [...events.slice(-11), value.event]);
          } catch {
            /* Terminal bytes use binary frames; control messages are JSON. */
          }
        }
      };
      ws.onclose = () => {
        if (!disposed) setMode("stopped");
      };
      ws.onerror = () => {
        if (!disposed) setMode("stopped");
      };
      term.onData((data) => {
        if (ws?.readyState === WebSocket.OPEN)
          ws.send(JSON.stringify({ type: "input", data }));
      });
    };
    void start();
    return () => {
      disposed = true;
      controller.abort();
      cancelAnimationFrame(refresh);
      ws?.close();
      term?.dispose();
      if (terminal.current === term) terminal.current = null;
      if (socket.current === ws) socket.current = null;
    };
    // Size changes use the existing live session, preserving selection and focus.
    // Other fixture changes and Reset create a fresh Go model.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [storyId, configurationKey, reset, savedVariant, validSize]);

  useEffect(() => {
    if (mode === "static" && snapshot) paintSnapshot.current?.(snapshot);
  }, [mode, snapshot]);

  useEffect(() => {
    if (mode !== "live" || !validSize) return;
    terminal.current?.resize(columns, rows);
    if (socket.current?.readyState === WebSocket.OPEN)
      socket.current.send(JSON.stringify({ type: "resize", columns, rows }));
  }, [columns, rows, mode, validSize]);

  if (!entry || !snapshot)
    return <p role="alert">Unknown Go story: {storyId}</p>;
  if (!validSize) return <p role="alert">Use 20–240 columns and 8–80 rows.</p>;
  const releaseFocus = () => {
    terminal.current?.blur();
    release.current?.focus();
  };
  return (
    <section ref={shell} className="terminal-story" aria-label={entry.title}>
      <div className="terminal-toolbar">
        <div className="terminal-identity">
          <span className={`mode-dot mode-${mode}`} />
          <strong>
            {mode === "live"
              ? "Live Go"
              : mode === "static"
                ? "Saved Go snapshot"
                : mode === "stopped"
                  ? "Session ended"
                  : "Connecting"}
          </strong>
          <span>{storyId}</span>
        </div>
        <div className="terminal-actions">
          {mode === "static" && (
            <select
              aria-label="Saved variant"
              value={snapshot.name}
              onChange={(event) => setSavedVariant(event.target.value)}
            >
              {entry.snapshots.map((value) => (
                <option key={value.name} value={value.name}>
                  {value.name} · {value.config.columns}×{value.config.rows}
                </option>
              ))}
            </select>
          )}
          <button onClick={() => terminal.current?.focus()}>
            Focus terminal
          </button>
          <button ref={release} onClick={releaseFocus}>
            Release focus
          </button>
          <button
            disabled={mode !== "live" || !entry.step}
            onClick={() =>
              socket.current?.send(JSON.stringify({ type: "step" }))
            }
          >
            Step
          </button>
          <button onClick={() => setReset((value) => value + 1)}>Reset</button>
          <button onClick={() => shell.current?.requestFullscreen()}>
            Fullscreen
          </button>
        </div>
      </div>
      <div
        className="terminal-scroller"
        style={{
          background: colors(
            mode === "static" ? snapshot.config.theme : props.theme,
          ).background,
        }}
      >
        <div ref={host} className="terminal-host" />
      </div>
      <div className="terminal-caption" role="status">
        {mode === "live"
          ? `${columns} × ${rows} cells · upstream Bubble Tea in an isolated PTY · Ctrl+] releases keyboard focus`
          : mode === "static"
            ? `Static snapshot: ${snapshot.name}. Start npm run dev for live keyboard input and controls.`
            : mode === "stopped"
              ? "The Go session ended. Reset starts a fresh fixture."
              : "Checking for the native preview service…"}
      </div>
      <details className="terminal-readable">
        <summary>Readable terminal output</summary>
        <pre data-testid="terminal-text">{readable}</pre>
      </details>
      <details className="terminal-events" open={events.length > 0}>
        <summary>
          Go events <span>{events.length}</span>
        </summary>
        <pre data-testid="terminal-events">
          {events.length
            ? events.map((event) => JSON.stringify(event)).join("\n")
            : "Interact with the live fixture to inspect emitted events."}
        </pre>
      </details>
    </section>
  );
}
