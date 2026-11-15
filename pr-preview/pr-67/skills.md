# Agent skills and Markdown

Charmed publishes portable agent guidance and a Markdown edition of every guide page. Download the relevant package for your project, or fetch individual reference pages when an agent needs current API examples and lifecycle contracts.

## [Choose a skill](https://lanej.io/charmed/skills.html#choose-a-skill)

Choose guidance for the work you are doing: building an application with Charmed or contributing to the library itself. Each package contains its own `SKILL.md` entrypoint and references canonical documentation without installing tools or replacing project policy.

| Skill                | Scope                                                                          | Downloads                                                                                                                                                     |
|----------------------|--------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Charmed              | Components, layouts, native lifecycle and retained rendering                   | [Entrypoint](https://lanej.io/charmed/skills/charmed/SKILL.md) · [Complete package](https://lanej.io/charmed/skills/charmed.tar.gz)                           |
| Charmed contributing | Library architecture, native examples, repository checks and website standards | [Entrypoint](https://lanej.io/charmed/skills/charmed-contributing/SKILL.md) · [Complete package](https://lanej.io/charmed/skills/charmed-contributing.tar.gz) |

## [Install a package](https://lanej.io/charmed/skills.html#install-a-package)

Select `charmed` for application work, or `charmed-contributing` for library contributions. Charmed source checkouts already contain both skills; use those copies. This example installs into the repository's `.agents/skills` directory. Choose the skill directory supported by your agent when it differs. Inspect the downloaded entrypoint and package before extraction; an existing skill directory is left untouched.

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

Load or refresh your agent's skills using its normal discovery workflow. Installing the package does not make its guidance authoritative over project instructions. Application authors use their own project checks; contributors follow the checked-out Charmed instructions and existing repository tooling.

## [Read the website as Markdown](https://lanej.io/charmed/skills.html#read-the-website-as-markdown)

Replace a page's `.html` suffix with `.md`, or use its Markdown footer link. Exports contain the page's complete main content, including generated component usage, lifecycle guidance, tables and plain-text terminal previews; navigation chrome and color styling are omitted. Links resolve to the corresponding web references.

```sh
curl --fail --show-error --location https://lanej.io/charmed/llms.txt
curl --fail --show-error --location https://lanej.io/charmed/getting-started.md
curl --fail --show-error --location https://lanej.io/charmed/atoms-chart.md
curl --fail --show-error --location https://lanej.io/charmed/skills/charmed/SKILL.md
```

The [Markdown index](https://lanej.io/charmed/llms.txt) lists every generated page and the raw skill entrypoints. HTML, Markdown and skill packages are published together from the same source revision; the files describe the current supported library rather than its implementation history.
