set shell := ["bash", "-eu", "-o", "pipefail", "-c"]
set positional-arguments

export GOTOOLCHAIN := env_var_or_default("GOTOOLCHAIN", "go" + `awk '$1 == "go" { print $2; exit }' go.mod`)

# List available targets.
default:
    @just --list

# Run ShellCheck and the full Go quality gate used in CI.
check: shellcheck
    bash scripts/check-go.sh

# Check shell scripts, formatting, vet, module tidiness and strict lint.
lint: shellcheck
    bash scripts/check-go.sh lint

# Check repository shell scripts, including portable profile assets.
shellcheck:
    bash scripts/check-shell.sh

# Run uncached race tests.
test:
    bash scripts/check-go.sh test

# Check for reachable vulnerabilities.
vuln:
    bash scripts/check-go.sh vuln

# Format explicitly selected Go files.
fmt +files:
    bash scripts/check-go.sh fmt "$@"

# Apply available whitespace fixes to selected packages.
fix +packages:
    bash scripts/check-go.sh fix "$@"

# Remove resolved findings from the existing baseline.
ratchet:
    bash scripts/check-go.sh ratchet

# Install the pinned lint and vulnerability tools.
tools:
    bash scripts/check-go.sh tools

# Generate the static guide in docs/dist.
docs:
    go run ./cmd/workbench generate --out docs

# Run a native fixture, optionally followed by workbench flags.
run story="screen/services" *args:
    go run ./cmd/workbench run "$@"
