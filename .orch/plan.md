# Plan

Append a short `## Running Tests` section to the end of `CONTRIBUTING.md`, after the existing `## Code Style` section, telling contributors to run `go test ./...` from the repo root. Pure append, plain Markdown, no other files touched.

## Work items

### WI-1 — docs(contributing): add Running Tests section

**Complexity:** trivial

Append a new `## Running Tests` section to the end of `CONTRIBUTING.md`, immediately after the existing `## Code Style` section. The section states that `go test ./...` runs the full suite from the repo root and notes that all tests must pass before a PR can be merged, per the spec's 'User-Visible Surface' block. This eliminates the friction of a new contributor having to guess the test invocation.

**Acceptance:**

- `CONTRIBUTING.md` contains a new `## Running Tests` heading placed after the `## Code Style` section, with no existing sections altered.
- The new section includes a fenced ```bash code block containing exactly `go test ./...` as the command.
- The new section includes a closing sentence stating that all tests must pass before a pull request can be merged, and contains no tables, badges, or HTML.
- No files other than `CONTRIBUTING.md` are added, removed, or modified.

