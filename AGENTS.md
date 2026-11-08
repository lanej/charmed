# Development

- Follow [Builder](.agents/skills/builder/SKILL.md) for software work and the [Go skill](.agents/skills/go/SKILL.md) for Go changes.
- Preserve the deliberate upstream API contract in [ARCHITECTURE.md](ARCHITECTURE.md). Document wrapper lifecycle effects on native customization.
- Run `bash scripts/check-go.sh lint` while editing and `bash scripts/check-go.sh` before handoff.
- Keep the lint baseline shrinking: run `bash scripts/check-go.sh ratchet` after fixes. Never enlarge allowances, add blanket suppressions or filter checks to changed lines.
- For guide or fixture changes, run `go run ./cmd/workbench generate --out docs`. Keep Markdown paragraphs on one source line.
