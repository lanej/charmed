#!/usr/bin/env python3
"""Manage per-PR Charmed Pages links without changing the rest of a PR body."""

import argparse
import json
import os
import re
import urllib.request

BEGIN = "<!-- charmed-preview:start -->"
END = "<!-- charmed-preview:end -->"
BLOCK = re.compile(re.escape(BEGIN) + r".*?" + re.escape(END) + r"\n*", re.S)


def preview_url(number: int) -> str:
    if number <= 0:
        raise ValueError("PR number must be positive")
    return f"https://lanej.io/charmed/pr-preview/pr-{number}/"


def description(body: str, number: int, revision: str, remove: bool = False) -> str:
    body = BLOCK.sub("", body or "").strip()
    if remove:
        return body
    base = preview_url(number)
    block = "\n".join(
        [
            BEGIN,
            "## Live preview",
            "",
            f"- [Components]({base}components.html)",
            f"- [Brand sheet]({base}brand.html)",
            f"- [Full catalog]({base})",
            "",
            f"Built from `{revision[:12]}` · updated on every PR push.",
            END,
        ]
    )
    return f"{block}\n\n{body}" if body else block


def github_request(path: str, method: str = "GET", payload: dict | None = None) -> dict:
    token = os.environ["GH_TOKEN"]
    data = json.dumps(payload).encode() if payload is not None else None
    req = urllib.request.Request(
        f"https://api.github.com/{path}",
        data=data,
        method=method,
        headers={
            "Authorization": f"Bearer {token}",
            "Accept": "application/vnd.github+json",
            "X-GitHub-Api-Version": "2022-11-28",
            "Content-Type": "application/json",
            "User-Agent": "charmed-preview-publisher",
        },
    )
    with urllib.request.urlopen(req, timeout=20) as response:
        return json.load(response)


def update(repo: str, number: int, revision: str, remove: bool) -> bool:
    pr = github_request(f"repos/{repo}/pulls/{number}")
    old = pr.get("body") or ""
    new = description(old, number, revision, remove)
    if new == old:
        return False
    github_request(f"repos/{repo}/pulls/{number}", "PATCH", {"body": new})
    return True


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--repo", required=True)
    parser.add_argument("--pr", type=int, required=True)
    parser.add_argument("--sha", default="")
    parser.add_argument("--remove", action="store_true")
    args = parser.parse_args()
    if not args.remove and not re.fullmatch(r"[0-9a-f]{40}", args.sha):
        parser.error("--sha requires the full 40-character commit SHA")
    if update(args.repo, args.pr, args.sha, args.remove):
        print("PR description updated")
    else:
        print("PR description already current")


if __name__ == "__main__":
    main()
