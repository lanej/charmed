# Development

Use the repository-local Builder core in `.agents/skills/builder/SKILL.md` for software work and `.agents/skills/go/SKILL.md` for Go changes.

Run `bash scripts/check-go.sh lint` while editing and `bash scripts/check-go.sh` before handoff. The runner selects the module's Go version by default. Preserve the strict profile and its shrinking baseline: never add baseline allowances, suppress new findings broadly, or filter checks to changed lines. After fixing existing findings, run `bash scripts/check-go.sh ratchet` and include the baseline reduction.

Run `go run ./cmd/workbench generate --out docs` when changing the static guide or its fixtures. Keep Markdown paragraphs on one source line.
