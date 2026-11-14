# Agent skills and Markdown

Charmed publishes portable agent guidance and a Markdown edition of every guide page. Download the relevant package for your project, or fetch individual reference pages when an agent needs current API examples and lifecycle contracts.

## Choose a skill

Packages contain a `SKILL.md` entrypoint plus all bundled references, scripts and assets. The Charmed skill helps application authors use the library; Builder and Go provide contributor guidance without installing dependencies or replacing project policy.

| Skill | Scope | Downloads |
| --- | --- | --- |
| Charmed | Components, layouts, native lifecycle and retained rendering | [Entrypoint](skills/charmed/SKILL.md) · [Complete package](skills/charmed.tar.gz) |
| Builder | Software contracts, architecture and verification | [Entrypoint](skills/builder/SKILL.md) · [Complete package](skills/builder.tar.gz) |
| Go | Correctness checks, enums, errors and portable quality assets | [Entrypoint](skills/go/SKILL.md) · [Complete package](skills/go.tar.gz) |

## Install a package

Run from your application's root and select `charmed`, `builder` or `go`. This example installs into the repository's `.agents/skills` directory. Choose the skill directory supported by your agent when it differs. Inspect the downloaded entrypoint and package before extraction; an existing skill directory is left untouched.

```sh
(
  set -eu
  skill=charmed
  destination=.agents/skills
  archive=$(mktemp)
  trap 'rm -f "$archive"' EXIT
  curl --fail --show-error --location \
    "https://lanej.io/charmed/skills/$skill.tar.gz" -o "$archive"
  tar -tzf "$archive"
  mkdir -p "$destination"
  test ! -e "$destination/$skill"
  test ! -L "$destination/$skill"
  tar -xzf "$archive" -C "$destination"
)
```

Load or refresh your agent's skills using its normal discovery workflow. Installing the package does not make its guidance authoritative over project instructions. Builder's example checker and Go's optional portable quality assets have their own prerequisites documented inside each package.

## Read the website as Markdown

Replace a page's `.html` suffix with `.md`, or use its Markdown footer link. Exports contain the page's complete main content, including generated component usage, lifecycle guidance, tables and plain-text terminal previews; navigation chrome and color styling are omitted. Links resolve to the corresponding web references.

```sh
curl --fail --show-error --location https://lanej.io/charmed/llms.txt
curl --fail --show-error --location https://lanej.io/charmed/getting-started.md
curl --fail --show-error --location https://lanej.io/charmed/atoms-chart.md
curl --fail --show-error --location https://lanej.io/charmed/skills/charmed/SKILL.md
```

The [Markdown index](llms.txt) lists every generated page and the raw skill entrypoints. HTML, Markdown and skill packages are published together from the same source revision; the files describe the current supported library rather than its implementation history.
