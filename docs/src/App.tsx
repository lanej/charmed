import {
  lazy,
  Suspense,
  useEffect,
  useMemo,
  useState,
  type ComponentType,
} from "react";
import manifest from "../generated/manifest.json";

const generated = import.meta.glob<{ default: ComponentType }>(
  "../generated/*.mdx",
);
const pages = import.meta.glob<{ default: ComponentType }>("../pages/*.mdx");
const levels = [
  "Foundations",
  "Atoms",
  "Molecules",
  "Organisms",
  "Templates",
  "Screens",
];
const fixedPages: Record<string, { title: string; file: string }> = {
  overview: { title: "Charmed", file: "../pages/Welcome.mdx" },
  composition: { title: "Composition", file: "../pages/Composition.mdx" },
  development: {
    title: "Develop and publish",
    file: "../pages/Development.mdx",
  },
};
function route() {
  return decodeURIComponent(location.hash.slice(1)) || "overview";
}

export function App() {
  const [current, setCurrent] = useState(route);
  const [query, setQuery] = useState("");
  const [groupBy, setGroupBy] = useState<"level" | "family">("level");
  const [menuOpen, setMenuOpen] = useState(false);
  useEffect(() => {
    const changed = () => {
      setCurrent(route());
      setMenuOpen(false);
      window.scrollTo(0, 0);
    };
    window.addEventListener("hashchange", changed);
    return () => window.removeEventListener("hashchange", changed);
  }, []);
  const entry = manifest.find((entry) => entry.id === current);
  const title =
    entry?.title.split("/").at(-1) ?? fixedPages[current]?.title ?? "Not found";
  useEffect(() => {
    document.title = `${title} · Charmed`;
  }, [title]);
  const Page = useMemo(() => {
    const file = fixedPages[current]?.file;
    const loader = file
      ? pages[file]
      : generated[`../generated/${current.replaceAll("/", "_")}.mdx`];
    return loader ? lazy(loader) : null;
  }, [current]);
  const visible = manifest.filter((entry) =>
    `${entry.title} ${entry.family} ${entry.description}`
      .toLowerCase()
      .includes(query.toLowerCase()),
  );
  const groups =
    groupBy === "level"
      ? levels
      : Array.from(new Set(manifest.map((entry) => entry.family))).sort();
  return (
    <div className="catalog-shell">
      <header className="catalog-header">
        <button
          className="menu-button"
          onClick={() => setMenuOpen((value) => !value)}
          aria-expanded={menuOpen}
          aria-controls="catalog-navigation"
        >
          Components
        </button>
        <a className="catalog-brand" href="#overview">
          charmed<span>terminal design system</span>
        </a>
        <a
          className="repository-link"
          href="https://github.com/lanej/charmed"
          target="_blank"
          rel="noreferrer"
        >
          GitHub ↗
        </a>
      </header>
      <aside
        id="catalog-navigation"
        className={`catalog-sidebar ${menuOpen ? "sidebar-open" : ""}`}
      >
        <label className="component-search">
          <span>Find a component</span>
          <input
            type="search"
            placeholder="Search components…"
            value={query}
            onChange={(event) => setQuery(event.target.value)}
          />
        </label>
        <nav aria-label="Documentation">
          {Object.entries(fixedPages).map(([id, page]) => (
            <a
              key={id}
              href={`#${id}`}
              className={current === id ? "active" : ""}
              aria-current={current === id ? "page" : undefined}
            >
              {page.title}
            </a>
          ))}
        </nav>
        <div
          className="catalog-grouping"
          role="group"
          aria-label="Group components"
        >
          <button
            aria-pressed={groupBy === "level"}
            onClick={() => setGroupBy("level")}
          >
            Atomic
          </button>
          <button
            aria-pressed={groupBy === "family"}
            onClick={() => setGroupBy("family")}
          >
            Capability
          </button>
        </div>
        <nav aria-label="Components">
          {groups.map((group) => {
            const entries = visible.filter((entry) => entry[groupBy] === group);
            return entries.length ? (
              <section key={group}>
                <h2>
                  {group}
                  <span>{entries.length}</span>
                </h2>
                {entries.map((entry) => (
                  <a
                    key={entry.id}
                    href={`#${entry.id}`}
                    className={current === entry.id ? "active" : ""}
                    aria-current={current === entry.id ? "page" : undefined}
                  >
                    <span>{entry.title.split("/").at(-1)}</span>
                    <small>{entry.family}</small>
                  </a>
                ))}
              </section>
            ) : null;
          })}
          {!visible.length && (
            <p className="no-results">No components match “{query}”.</p>
          )}
        </nav>
        <p className="sidebar-footer">
          {manifest.length} Go fixtures
          <br />
          One registry. Native terminal output.
        </p>
      </aside>
      <main id="main-content" className="catalog-main">
        <div className="catalog-breadcrumb">
          {entry ? `${entry.level} / ${entry.family}` : "Documentation"}
          <span>MDX</span>
        </div>
        <article className="mdx-content" key={current}>
          <Suspense fallback={<p>Loading component…</p>}>
            {Page ? (
              <Page />
            ) : (
              <>
                <h1>Component not found</h1>
                <p>Choose a component from the catalog.</p>
              </>
            )}
          </Suspense>
        </article>
        <footer className="catalog-footer">
          Charmed · built with Go, Charm and ntcharts
        </footer>
      </main>
    </div>
  );
}
