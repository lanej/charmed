#!/usr/bin/env bash
# Publish a prebuilt PR snapshot from an unprivileged job. This script runs
# only from the trusted main checkout; it never checks out or executes PR code.
set -euo pipefail

pr="$1"
action="$2"
sha="$3"
case "$pr" in *[!0-9]*|"") echo "Invalid PR number" >&2; exit 1 ;; esac
if [[ "$action" != "closed" && ! "$sha" =~ ^[0-9a-f]{40}$ ]]; then
  echo "Expected 40-character PR commit SHA" >&2
  exit 1
fi
: "$GITHUB_REPOSITORY"
: "$GH_TOKEN"

repo="$GITHUB_REPOSITORY"
workspace="$PWD"
current_repo="$(gh api "repos/$repo/pulls/$pr" --jq '.head.repo.full_name')"
if [[ "$current_repo" != "$repo" ]]; then
  echo "Forked PR previews require a separate trusted artifact flow; refusing publication" >&2
  exit 1
fi

if [[ "$action" != "closed" ]]; then
  current_sha="$(gh api "repos/$repo/pulls/$pr" --jq '.head.sha')"
  if [[ "$current_sha" != "$sha" ]]; then
    echo "PR #$pr advanced to $current_sha; skip stale build $sha"
    exit 0
  fi
  test -f "$workspace/preview-artifact/index.html"
  test "$(cat "$workspace/preview-artifact/revision.txt")" = "$sha"
fi

# The preview branch contains *only generated assets*, not executable sources.
# Do not checkout a PR ref with a token that can push to the repository.
store="$(mktemp -d)"
trap 'rm -rf "$store"' EXIT
cd "$store"
git init -q
git remote add origin "https://github.com/$repo.git"
gh auth setup-git
if git ls-remote --exit-code origin refs/heads/pr-preview-cache >/dev/null 2>&1; then
  git fetch -q --depth=1 origin refs/heads/pr-preview-cache
  git switch -q -c pr-preview-cache FETCH_HEAD
else
  git switch -q --orphan pr-preview-cache
fi
target="pr-preview/pr-$pr"
rm -rf "$target"
if [[ "$action" != "closed" ]]; then
  mkdir -p "$target"
  cp -a "$workspace/preview-artifact/." "$target/"
fi
git add -A
if ! git diff --cached --quiet; then
  git -c user.name='github-actions[bot]' \
      -c user.email='41898282+github-actions[bot]@users.noreply.github.com' \
      commit -qm "docs(preview): $action PR #$pr"
  git push -q origin HEAD:refs/heads/pr-preview-cache
fi
cd "$workspace"

# GitHub Actions deliberately suppresses push events caused by GITHUB_TOKEN.
# Dispatch the trusted main publisher rather than relying on that push.
gh api --method POST "repos/$repo/actions/workflows/pages.yml/dispatches" \
  -f ref=main >/dev/null

if [[ "$action" == "closed" ]]; then
  python3 scripts/preview_pr.py --repo "$repo" --pr "$pr" --remove
  exit 0
fi

# Only advertise a URL after the real custom-domain response proves that Pages
# contains this PR's actual commit. Otherwise CI fails and no false link appears.
revision_url="https://lanej.io/charmed/$target/revision.txt"
ready=false
for _ in $(seq 1 50); do
  published="$(curl --connect-timeout 5 --max-time 10 -fsSL "$revision_url" 2>/dev/null || true)"
  if [[ "$published" == "$sha" ]]; then
    ready=true
    break
  fi
  sleep 8
done
if [[ "$ready" != true ]]; then
  echo "Preview did not publish expected revision $sha at $revision_url" >&2
  exit 1
fi
python3 scripts/preview_pr.py --repo "$repo" --pr "$pr" --sha "$sha"
echo "Verified preview: https://lanej.io/charmed/$target/components.html"
