# Development

- Follow [Builder](.agents/skills/builder/SKILL.md) for software work and the [Go skill](.agents/skills/go/SKILL.md) for Go changes.
- Preserve the deliberate upstream API contract in [ARCHITECTURE.md](ARCHITECTURE.md). Document wrapper lifecycle effects on native customization.
- Use the targets in [justfile](justfile): `just lint` while editing and `just check` before handoff. Both include ShellCheck; use `just shellcheck` for shell changes. Run `just` to list targets.
- Keep the lint baseline shrinking: run `just ratchet` after fixes. Never enlarge allowances, add blanket suppressions or filter checks to changed lines.
- For guide or fixture changes, run `just docs`. Keep Markdown paragraphs on one source line.
