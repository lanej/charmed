import manifest from "../generated/manifest.json";
import coverage from "../generated/coverage.json";
import { TerminalStory } from "./TerminalStory";

const screen = manifest.find((entry) => entry.id === "screen/services")!;
const layers = [
  ["Foundations", "Themes, cell geometry, focus and motion"],
  ["Atoms", "Inputs, feedback and terminal charts"],
  ["Molecules", "Fields, legends and coordinated controls"],
  ["Organisms", "Collections, forms and document viewers"],
  ["Screens", "Linked components in realistic workflows"],
];

export function Overview() {
  return (
    <div className="charmed-overview">
      <div className="hero-kicker">GO COMPONENTS · TERMINAL STORIES · MDX</div>
      <h1>
        Charmed
        <span>
          One design language.
          <br />
          Every terminal.
        </span>
      </h1>
      <p className="hero-description">
        Compose Charm and ntcharts through shared themes, cell layouts and
        interaction rules. Explore the same Go components in your terminal and
        in this workbench.
      </p>
      <div className="hero-stats">
        <span>
          <strong>{manifest.length}</strong> Go stories
        </span>
        <span>
          <strong>10</strong> chart families
        </span>
        <span>
          <strong>4</strong> color profiles
        </span>
        <span>
          <strong>0</strong> required forks
        </span>
      </div>
      <div className="layer-grid">
        {layers.map(([name, detail]) => (
          <div key={name}>
            <strong>{name}</strong>
            <p>{detail}</p>
          </div>
        ))}
      </div>
      <h2>Start with the whole screen</h2>
      <p>
        Filter the service table, move selection, and inspect its throughput.
        Tab changes focus. Resizing preserves the selected service; background
        updates keep running.
      </p>
      <TerminalStory storyId={screen.id} {...screen.defaults} />
      <h2>Coverage you can inspect</h2>
      <p>
        Core widgets and charts run real Go fixtures. Optional integrations
        retain explicit status.
      </p>
      <div className="coverage-table">
        <table>
          <thead>
            <tr>
              <th>Upstream</th>
              <th>Status</th>
              <th>Integration</th>
            </tr>
          </thead>
          <tbody>
            {coverage.map((row) => (
              <tr key={row.upstream}>
                <td>{row.upstream}</td>
                <td>
                  <span className={`coverage-status coverage-${row.status}`}>
                    {row.status}
                  </span>
                </td>
                <td>{row.notes}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
