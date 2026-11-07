import { useEffect, useState } from "react";
import manifest from "../generated/manifest.json";

export function GoSource({ storyId }: { storyId: string }) {
  const entry = manifest.find((entry) => entry.id === storyId);
  const [html, setHtml] = useState("");
  const [copied, setCopied] = useState(false);
  useEffect(() => {
    let cancelled = false;
    if (entry)
      void import("./highlight").then(async ({ highlightGo }) => {
        const output = await highlightGo(entry.code);
        if (!cancelled) setHtml(output);
      });
    return () => {
      cancelled = true;
    };
  }, [entry]);
  if (!entry) return null;
  return (
    <div className="go-source">
      <div className="source-heading">
        <a
          href={`https://github.com/lanej/charmed/blob/main/${entry.source}`}
          target="_blank"
          rel="noreferrer"
        >
          {entry.source}
        </a>
        <button
          onClick={() =>
            void navigator.clipboard
              .writeText(entry.code)
              .then(() => setCopied(true))
          }
        >
          {copied ? "Copied" : "Copy Go"}
        </button>
      </div>
      {html ? (
        <div dangerouslySetInnerHTML={{ __html: html }} />
      ) : (
        <pre>
          <code>{entry.code}</code>
        </pre>
      )}
      <p>
        <a href={entry.upstream} target="_blank" rel="noreferrer">
          Upstream library ↗
        </a>
      </p>
    </div>
  );
}
