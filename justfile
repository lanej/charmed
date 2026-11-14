set shell := ["bash", "-eu", "-o", "pipefail", "-c"]
set positional-arguments

export GOTOOLCHAIN := env_var_or_default("GOTOOLCHAIN", "go" + `awk '$1 == "go" { print $2; exit }' go.mod`)

# List available targets.
default:
    @just --list

# Run ShellCheck and the full Go quality gate used in CI.
check: shellcheck docs-check perf-check
    bash scripts/check-go.sh

# Check shell scripts, formatting, vet, module tidiness and strict lint.
lint: shellcheck docs-check
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
docs: docs-check
    # Preview builds check out the PR head even when the event SHA points to main.
    GITHUB_SHA="${GITHUB_SHA:+$(git rev-parse HEAD)}" go run ./cmd/workbench generate --out docs
    python3 scripts/check_docs.py --rendered docs/dist

# Run a native fixture, optionally followed by workbench flags.
run story="screen/services" *args:
    go run ./cmd/workbench run "$@"

# Check content bounds without compiling or opening a browser.
docs-check:
    python3 scripts/check_docs.py
    python3 -m unittest discover -s scripts -p 'test_docs_standard.py'

# Install the optional staged-tree source gate for this checkout.
hooks:
    git config core.hooksPath .githooks

# Record native performance reports; use --suite hotspots or profiles to focus.
perf *args:
    uv run --no-project scripts/perf.py "$@"

# Sample native PTY latency, idle resources and resize memory.
perf-native *args:
    uv run --no-project scripts/perf_native.py "$@"

# Validate performance reporting contracts; Go structural gates run in check.
perf-check:
    python3 -m unittest discover -s scripts -p 'test_perf*.py'
