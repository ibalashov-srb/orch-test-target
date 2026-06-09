# Plan

Replace the existing repository-root `CONTRIBUTING.md` with independently authored, plain-Markdown content (~25 lines) that covers the spec's required sections — intro/purpose, reporting issues, submitting changes, commit style, and Go-specific code style — while staying well under the 60-line cap and containing no badge syntax. Single-file scope; no other files are touched.

## Work items

### WI-1 — docs: rewrite CONTRIBUTING.md with concise contributor guide

**Complexity:** trivial

Overwrite the existing `CONTRIBUTING.md` at the repository root with freshly authored, plain GitHub-flavored Markdown content satisfying the spec's constraints (see SpecDoc §User-Visible Surface). The new file must cover intro/purpose, reporting issues, submitting changes, commit messages, and code style (with Go-specific `go fmt` / `go vet` hints, per SpecDoc §Implementation Notes). The pipeline requires an observable diff, so the file is replaced even though the prior version is mechanically compliant.

**Acceptance:**

- `CONTRIBUTING.md` at the repository root has fewer than 60 lines total (including blank lines), and its content differs from the prior version committed at `main`.
- The file contains no badge syntax — `grep -E '\[!\[.*\]\(.*\)\]\(.*\)' CONTRIBUTING.md` returns no matches — and uses only plain GitHub-flavored Markdown (headings, lists, inline code).
- The file includes distinct sections covering, at minimum: an intro/purpose paragraph, reporting issues, submitting changes, and commit message style; a code-style section mentions `go fmt` and `go vet`.
- No file other than `CONTRIBUTING.md` is added, modified, or deleted by this change.

