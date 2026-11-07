import { useEffect, useState } from "react";
import manifest from "../generated/manifest.json";
import { TerminalStory } from "./TerminalStory";

type Config = (typeof manifest)[number]["defaults"];
function Dimension({
  label,
  value,
  min,
  max,
  onChange,
}: {
  label: string;
  value: number;
  min: number;
  max: number;
  onChange: (value: number) => void;
}) {
  const [draft, setDraft] = useState(String(value));
  useEffect(() => setDraft(String(value)), [value]);
  const commit = () => {
    const number = Number(draft);
    if (Number.isInteger(number) && number >= min && number <= max)
      onChange(number);
    else setDraft(String(value));
  };
  return (
    <label>
      {label}
      <input
        aria-label={label}
        type="number"
        min={min}
        max={max}
        value={draft}
        onChange={(event) => setDraft(event.target.value)}
        onBlur={commit}
        onKeyDown={(event) => {
          if (event.key === "Enter") event.currentTarget.blur();
        }}
      />
    </label>
  );
}

function Identifier({ label, value, other, onChange }: {
  label: string; value: string; other: string; onChange: (value: string) => void;
}) {
  const [draft, setDraft] = useState(value);
  useEffect(() => setDraft(value), [value]);
  return (
    <label>
      {label}
      <input
        aria-label={label}
        value={draft}
        maxLength={128}
        onChange={(event) => {
          event.currentTarget.setCustomValidity("");
          setDraft(event.target.value);
        }}
        onBlur={(event) => {
          const invalid = !draft.trim() || new TextEncoder().encode(draft).length > 128 || /\p{Cc}/u.test(draft) || draft === other;
          event.currentTarget.setCustomValidity(invalid ? "Use a distinct, printable version identifier of 1–128 bytes." : "");
          if (invalid) event.currentTarget.reportValidity();
          else onChange(draft);
        }}
        onKeyDown={(event) => { if (event.key === "Enter") event.currentTarget.blur(); }}
      />
    </label>
  );
}

export function StoryCanvas({
  storyId,
  variant = "Default",
}: {
  storyId: string;
  variant?: string;
}) {
  const entry = manifest.find((entry) => entry.id === storyId);
  const [config, setConfig] = useState<Config>(
    () =>
      entry?.snapshots.find((snapshot) => snapshot.name === variant)?.config ??
      entry!.defaults,
  );
  const [presetRevision, setPresetRevision] = useState(0);
  useEffect(() => {
    if (entry)
      setConfig(
        entry.snapshots.find((snapshot) => snapshot.name === variant)?.config ??
          entry.defaults,
      );
  }, [entry, variant]);
  if (!entry) return <p role="alert">Unknown Go fixture: {storyId}</p>;
  const set = <K extends keyof Config>(key: K, value: Config[K]) =>
    setConfig((config) => ({ ...config, [key]: value }));
  const states = entry.states ?? [];
  return (
    <div className="story-canvas">
      <TerminalStory key={`${storyId}:${presetRevision}`} storyId={storyId} {...config} />
      <details className="fixture-controls" open>
        <summary>
          Fixture controls<span>Changes run through the Go model</span>
        </summary>
        <div className="control-grid">
          <label>
            Preset
            <select
              aria-label="Preset"
              value=""
              onChange={(event) => {
                const snapshot = entry.snapshots.find(
                  (snapshot) => snapshot.name === event.target.value,
                );
                if (snapshot) {
                  setConfig(snapshot.config);
                  setPresetRevision((revision) => revision + 1);
                }
              }}
            >
              <option value="" disabled>
                Choose a variant
              </option>
              {entry.snapshots.map((snapshot) => (
                <option key={snapshot.name} value={snapshot.name}>
                  {snapshot.name}
                </option>
              ))}
            </select>
          </label>
          <Dimension
            label="Columns"
            value={config.columns}
            min={20}
            max={240}
            onChange={(value) => set("columns", value)}
          />
          <Dimension
            label="Rows"
            value={config.rows}
            min={8}
            max={80}
            onChange={(value) => set("rows", value)}
          />
          <label>
            Theme
            <select
              aria-label="Theme"
              value={config.theme}
              onChange={(event) => set("theme", event.target.value)}
            >
              {["dark", "light", "contrast", "mono"].map((theme) => (
                <option key={theme}>{theme}</option>
              ))}
            </select>
          </label>
          <label>
            Color profile
            <select
              aria-label="Color profile"
              value={config.profile}
              onChange={(event) => set("profile", event.target.value)}
            >
              {["truecolor", "ansi256", "ansi16", "none"].map((profile) => (
                <option key={profile}>{profile}</option>
              ))}
            </select>
          </label>
          {entry.controls?.includes("state") && (
            <label>
              State
              <select
                aria-label="State"
                value={config.state}
                onChange={(event) => set("state", event.target.value)}
              >
                {states.map((state) => (
                  <option key={state}>{state}</option>
                ))}
              </select>
            </label>
          )}
          {entry.controls?.includes("value") && (
            <label>
              Value
              <input
                aria-label="Value"
                value={config.value}
                onChange={(event) => set("value", event.target.value)}
              />
            </label>
          )}
          {entry.controls?.includes("scenario") && (
            <label>
              Scenario
              <select
                aria-label="Scenario"
                value={config.scenario}
                onChange={(event) => set("scenario", event.target.value)}
              >
                {(entry.scenarios ?? []).map((scenario) => (
                  <option key={scenario}>{scenario}</option>
                ))}
              </select>
            </label>
          )}
          {entry.controls?.includes("referenceID") && (
            <Identifier label="Reference version" value={config.referenceID} other={config.candidateID} onChange={(value) => set("referenceID", value)} />
          )}
          {entry.controls?.includes("candidateID") && (
            <Identifier label="Candidate version" value={config.candidateID} other={config.referenceID} onChange={(value) => set("candidateID", value)} />
          )}
          {entry.controls?.includes("sigma") && (
            <label>
              Host spread bands
              <select
                aria-label="Host spread bands"
                value={config.sigma}
                onChange={(event) => set("sigma", Number(event.target.value))}
              >
                {[1, 2, 3].map((sigma) => (
                  <option key={sigma} value={sigma}>±{sigma}σ</option>
                ))}
              </select>
            </label>
          )}
          {entry.controls?.includes("boundaries") && (
            <label>
              Deviation boundaries
              <select aria-label="Deviation boundaries" value={config.boundaries ? "example" : "off"} onChange={(event)=>set("boundaries",event.target.value==="example")}>
                <option value="example">Example rules</option>
                <option value="off">No boundaries</option>
              </select>
            </label>
          )}
          {entry.controls?.includes("series") && (
            <Dimension
              label="Series"
              value={config.series}
              min={1}
              max={4}
              onChange={(value) => set("series", value)}
            />
          )}
          {entry.controls?.includes("amount") && (
            <label>
              Amount <span>{config.amount}</span>
              <input
                aria-label="Amount"
                type="range"
                min="0"
                max="100"
                step="1"
                value={config.amount}
                onChange={(event) => set("amount", Number(event.target.value))}
              />
            </label>
          )}
          <div className="boolean-controls">
            {(["ascii", "reducedMotion", "live"] as const)
              .filter((key) => key !== "live" || entry.clock)
              .map((key) => (
                <label key={key}>
                  <input
                    type="checkbox"
                    checked={config[key]}
                    onChange={(event) => set(key, event.target.checked)}
                  />
                  {key === "ascii"
                    ? "ASCII"
                    : key === "live"
                      ? "Live clock"
                      : "Reduced motion"}
                </label>
              ))}
            {entry.controls?.includes("horizontal") && (
              <label>
                <input
                  type="checkbox"
                  checked={config.horizontal}
                  onChange={(event) => set("horizontal", event.target.checked)}
                />
                Horizontal bars
              </label>
            )}
          </div>
        </div>
      </details>
    </div>
  );
}
