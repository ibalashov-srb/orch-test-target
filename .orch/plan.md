# Plan

Overwrite the root `CONTRIBUTING.md` with a freshly authored, plain-Markdown version (≤60 lines, no badges) that preserves the existing five sections — intro, Reporting Issues, Submitting Changes, Commit Messages, Code Style — to produce a traceable diff while satisfying every stated constraint. This is a single-file change with no secondary edits (no README, Go source, CI, or `.github/` updates), so it is captured as one work item.

## Work items

### WI-1 — docs: rewrite CONTRIBUTING.md preserving five established sections

**Complexity:** trivial

Overwrite `CONTRIBUTING.md` at the repository root with freshly authored plain Markdown content that resolves SpecBrief open question q1 (replace, not no-op or patch). The new file must keep the same five sections currently present — intro pointing to `README.md`, Reporting Issues, Submitting Changes, Commit Messages, Code Style — and obey all three hard constraints from the request: root location, ≤60 lines, no badges. No other files in the repo are modified; per the spec's Out of Scope, `README.md`, Go sources, `go.mod`, CI, and `.github/` stay untouched.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root and `wc -l CONTRIBUTING.md` returns a value ≤ 60.
- The file is plain GitHub-Flavored Markdown with no badge images or shields.io references, and contains, in order, an intro paragraph that points to `README.md`, then H2 sections titled `Reporting Issues`, `Submitting Changes`, `Commit Messages`, and `Code Style` covering the bullet content listed in the spec's Required Sections.
- `git diff --name-only` against the base branch lists exactly one changed path: `CONTRIBUTING.md` (no edits to `README.md`, Go files, `go.mod`, `.gitignore`, CI config, or `.github/`), and the diff shows the file content has actually changed versus the prior version.

