# Development

- Read [the design philosophy](DESIGN.md) before component, visual or interaction work. Use it to choose the primary presentation; Unicode is the default and ASCII is an explicit fallback. See [guide design](docs/guide-design.md) for site-specific requirements.

- Follow [Builder](.agents/skills/builder/SKILL.md) for software work and the [Go skill](.agents/skills/go/SKILL.md) for Go changes.
- Preserve the deliberate upstream API contract in [ARCHITECTURE.md](ARCHITECTURE.md). Document wrapper lifecycle effects on native customization.
- Use the targets in [justfile](justfile): `just lint` while editing and `just check` before handoff. Both include ShellCheck; use `just shellcheck` for shell changes. Run `just` to list targets.
- Keep the lint baseline shrinking: run `just ratchet` after fixes. Never enlarge allowances, add blanket suppressions or filter checks to changed lines.
- For guide or fixture changes, run `just docs`. Keep Markdown paragraphs on one source line.

# Documentation

- Follow the page profiles and bounds in `docs/standards.json`; run `just docs-check` for fast feedback. Do not raise bounds to accommodate a content violation.
- Combine related prose before adding headings or artifacts. Headings and decorative images do not reset text-run budgets.
- Component entries require meaningful native output, complete consumer usage, lifecycle guidance and a runnable example. Review chart, input and inspector examples together at mobile and desktop widths when changing templates.
- `just hooks` installs the optional staged-tree pre-commit check. CI and `just docs` enforce the same source rules; browser CI enforces rendered bounds.
