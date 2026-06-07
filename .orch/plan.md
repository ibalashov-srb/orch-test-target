# Plan

Add a short "## Building" section to CONTRIBUTING.md between the intro paragraph and the existing "## Reporting Issues" section, citing `go build ./...` and `go test ./...` as inline code spans in a brief prose paragraph. Single-file documentation change, no code or tooling impact.

## Work items

### WI-1 — docs(contributing): add Building section with go build command

**Complexity:** trivial

Insert a new `## Building` section in `CONTRIBUTING.md` immediately after the introductory paragraph and before `## Reporting Issues`, so contributors learn the build command before any other workflow step. The section is a short prose paragraph noting that the project uses standard Go tooling, that `go build ./...` compiles all packages from the repo root, and that `go test ./...` runs the test suite — matching the file's existing inline-code-span style (no fenced code blocks, no bullets). This addresses the spec's 'What and Why' and 'User-Visible Surface' sections; no other files are touched.

**Acceptance:**

- `CONTRIBUTING.md` contains a level-2 heading `## Building` positioned after the intro paragraph and before the `## Reporting Issues` heading.
- The Building section body references the command `go build ./...` as an inline code span (single backticks, no fenced block) and mentions `go test ./...` for running tests.
- Section ordering in `CONTRIBUTING.md` is: intro paragraph, `## Building`, `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style`, with a blank line separating the new section from neighbors.
- No files other than `CONTRIBUTING.md` are modified in the change.

