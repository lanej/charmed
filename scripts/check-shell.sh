#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
command -v shellcheck >/dev/null || {
    printf 'Install ShellCheck, then rerun just shellcheck.\n' >&2
    exit 1
}

git ls-files --cached --others --exclude-standard -z -- '*.sh' |
    xargs -0 shellcheck --external-sources --
