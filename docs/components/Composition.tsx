import manifest from "../generated/manifest.json";

export function Composition({ storyId }: { storyId: string }) {
  const entry = manifest.find((entry) => entry.id === storyId);
  if (!entry) return null;
  const children = (entry.composes ?? []).map(
    (id) => manifest.find((item) => item.id === id)!,
  );
  const parents = manifest.filter((item) => (item.composes ?? []).includes(storyId));
  if (!children.length && !parents.length) return null;
  const groups = [
    { title: "Composed from", items: children },
    { title: "Used by", items: parents },
  ].filter((group) => group.items.length);
  return (
    <section className="component-composition" aria-label="Component composition">
      {groups.map((group) => (
        <div key={group.title}>
          <h2>{group.title}</h2>
          <ul>
            {group.items.map((item) => (
              <li key={item.id}>
                <a href={`#${item.id}`}>{item.title.split("/").at(-1)}</a>
                <span>{item.level}</span>
              </li>
            ))}
          </ul>
        </div>
      ))}
    </section>
  );
}
