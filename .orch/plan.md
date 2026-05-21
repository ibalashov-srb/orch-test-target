# Plan

Rewrite the repository-root `CONTRIBUTING.md` as a project-specific Go contributing guide covering prerequisites, getting started, code style, PR workflow, commit messages, and issue reporting — all under 60 physical lines, plain Markdown, no badges, with no other files touched.

## Work items

### WI-1 — docs: rewrite CONTRIBUTING.md as Go-specific contributing guide

**Complexity:** trivial

Replace the existing root `CONTRIBUTING.md` with a freshly authored guide for this Go 1.22 grep-like utility, per the spec's 'Target file content (normative)' section. The file must include, in order: an opening welcome naming `orch-test-target`; prerequisites (Go 1.22+); getting started with clone/fork plus `go build ./...` and `go test ./...`; code style with `go fmt ./...` and `go vet ./...`; PR workflow (branch from and target `main`, keep PRs small, document what and why); commit message rules (imperative mood, ≤72-char subject, `#N` issue refs); and issue reporting (search first, include OS, Go version, and repro steps). No other file in the repo is modified.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root and its content differs from the pre-change version (full rewrite, not an append).
- The file has fewer than 60 physical lines (a `wc -l` count returns ≤ 59).
- The file contains no badge/image syntax — i.e., no occurrence of the regex `!\[[^\]]*\]\([^)]*\)` — and uses only plain GitHub-Flavored Markdown.
- The file contains literal references to `go fmt`, `go vet`, `go test`, and `go build`, and names Go `1.22` as the minimum version.
- The file contains distinct sections (Markdown headings) for prerequisites, getting started, code style, pull-request workflow, commit messages, and reporting issues, appearing in that order.
- No file in the repository other than `CONTRIBUTING.md` is added, modified, or deleted by this change (verified via `git diff --name-only` returning exactly `CONTRIBUTING.md`).

