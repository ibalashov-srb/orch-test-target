# Plan

Overwrite the existing `CONTRIBUTING.md` at the repo root with freshly authored content per the spec: plain Markdown, no badges, ≤60 lines, five required sections (Reporting Issues, Submitting Changes, Commit Messages, Code Style) plus title and intro, with Go-specific tooling (`go fmt`, `go vet`) called out. Single-file change; no source, build, or CI modifications.

## Work items

### WI-1 — docs: author CONTRIBUTING.md with Go-specific contributor guide

**Complexity:** trivial

Replace the existing root-level `CONTRIBUTING.md` with a freshly authored guide that satisfies every constraint in the spec's User-Visible Surface section. The new file documents how to report issues, submit changes, format commit messages, and follow code style for this Go 1.22 CLI module, explicitly naming `go fmt` and `go vet` as the canonical pre-submission tools (per the spec's Implementation Notes). No other file in the repository is modified.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root, is valid UTF-8 with Unix line endings, contains no badges, shields, or HTML tags, and has a total line count ≤ 60.
- The file begins with an H1 `# Contributing`, followed by a short intro paragraph that points readers to `README.md`, then H2 sections in this order: `Reporting Issues`, `Submitting Changes`, `Commit Messages`, `Code Style`.
- `Reporting Issues` covers searching before filing, reproduction steps, and environment details (OS + Go version); `Submitting Changes` covers fork → feature branch off `main` → descriptive commits → PR against `main` → preference for small focused PRs; `Commit Messages` covers imperative mood, ≤72-char subject, and `#N` issue references.
- `Code Style` explicitly instructs contributors to run `go fmt` and `go vet` before submitting, in addition to matching surrounding style and minimizing unrelated changes.
- No file other than `CONTRIBUTING.md` is added, modified, or deleted by this work item.


## Quality warnings

- WI-1 has 5 acceptance bullets — 5+ usually means two WIs sharing a file; consider splitting at the Plan gate
