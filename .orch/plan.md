# Plan

Apply the spec's idempotent canonical write to `CONTRIBUTING.md`: replace the generic "Run any available formatters or linters" bullet in the Code Style section with the Go-specific `go fmt ./...` and `go vet ./...` instruction. All other sections, ordering, and content are preserved; the file remains plain Markdown, badge-free, and well under the 60-line ceiling.

## Work items

### WI-1 — docs(contributing): name go fmt and go vet in Code Style section

**Complexity:** trivial

The spec's Canonical Content Contract requires the Code Style section to explicitly reference `go fmt ./...` and `go vet ./...` since the project is a Go 1.22 module. The existing `CONTRIBUTING.md` is otherwise compliant (34 lines, plain Markdown, no badges, all four required sections in order), so the only change is replacing the final generic Code Style bullet with the Go-specific tooling line. This is an idempotent write: if the file already matches the canonical content, no diff is produced.

**Acceptance:**

- `CONTRIBUTING.md` at the repo root contains a bullet under `## Code Style` reading exactly: `- Run \`go fmt ./...\` and \`go vet ./...\` before submitting your changes.`
- `CONTRIBUTING.md` contains no badge images, no `<img>` tags, and no raw HTML; the file is at most 60 lines long; the four required H2 sections (`Reporting Issues`, `Submitting Changes`, `Commit Messages`, `Code Style`) appear in that order under the top-level `# Contributing` heading.
- The `## Submitting Changes` section continues to name `main` as the target branch, and the introductory paragraph still references `README.md`.
- No files other than `CONTRIBUTING.md` are modified (no changes to `README.md`, `go.mod`, source files, or CI config).

