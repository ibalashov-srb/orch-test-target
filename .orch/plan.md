# Plan

Replace the existing root `CONTRIBUTING.md` with a concise, repository-specific Markdown guide for this Go CLI project. The plan is documentation-only and intentionally excludes source, test, module, README, CI, or template changes.

## Work items

### WI-1 — docs(contributing): replace root guide

**Complexity:** trivial

Replace the existing generic root contributor guide with the concise content described in the spec sections “Required content” and “Implementation notes.” This work item delivers the only requested user-visible surface: a self-contained `CONTRIBUTING.md` tailored to the repository’s Go 1.22 CLI conventions. It preserves the documented out-of-scope boundaries by avoiding changes to code, tests, module metadata, README content, CI, or additional templates.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root as a plain Markdown file with fewer than 60 physical lines.
- The guide identifies the project as a small Go command-line project and directs contributors to `README.md` for `grep` usage and exit-code context.
- The guide includes the repo-specific prerequisites and commands: Go 1.22 or newer, `gofmt -w .`, `go test ./...`, and `go run ./cmd/grep <pattern> [file|-]`.
- The file contains no badges, images, HTML, generated table of contents, or references to unsupported infrastructure such as CI dashboards, security policy files, code of conduct files, or issue/PR templates.
- No files other than root-level `CONTRIBUTING.md` are changed.


## Quality warnings

- WI-1 has 5 acceptance bullets — 5+ usually means two WIs sharing a file; consider splitting at the Plan gate
