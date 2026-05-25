# Plan

Rewrite the root CONTRIBUTING.md as a Go-specific contributor guide tailored to this Go 1.22 CLI repo. The new file replaces the existing generic content with the seven sections specified (Title, Prerequisites, Running Tests, Reporting Issues, Submitting Changes, Commit Messages, Code Style), staying under 60 lines of plain Markdown with no badges or images.

## Work items

### WI-1 — docs: rewrite CONTRIBUTING.md as Go-specific contributor guide

**Complexity:** trivial

Overwrite the root `CONTRIBUTING.md` with a Go-tailored contributor guide as defined in the spec's 'Canonical content structure'. The file must include exactly the seven sections in order (Title + welcome, Prerequisites, Running Tests, Reporting Issues, Submitting Changes, Commit Messages, Code Style) and reference Go 1.22+, `go build ./...`, `go test ./...`, `go fmt ./...`, and `go vet ./...`. No other files in the repository are modified.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repo root and contains fewer than 60 lines total (including blank lines), ends with a single trailing newline, uses LF line endings, and has no trailing whitespace.
- The file contains exactly these seven top-level sections in this order: an H1 title with a one-line welcome, Prerequisites (mentioning Go 1.22+ and `go build ./...`), Running Tests (mentioning `go test ./...` and the test locations `cmd/grep/` and `doc_test.go`), Reporting Issues, Submitting Changes (fork → feature branch off `main` → PR against `main`), Commit Messages (imperative mood, ≤72-char subject, `Fixes #N`), and Code Style (`go fmt ./...`, `go vet ./...`, no unrelated reformatting).
- The file contains no Markdown image syntax (`![…](…)`), no `shields.io` or other badge URLs, and no references to CI/CD, CODEOWNERS, CHANGELOG, LICENSE, or `.github/` templates.
- No file other than `CONTRIBUTING.md` is added, modified, or deleted by the change (verified via `git diff --name-only` showing only `CONTRIBUTING.md`).

