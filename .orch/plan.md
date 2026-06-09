# Plan

Overwrite the repo-root CONTRIBUTING.md with fresh, project-aware content for this Go grep utility. The file must stay under 60 lines, contain no badge markup, and cover the five required sections (Getting Started, Reporting Issues, Submitting Changes, Commit Messages, Code Style) with Go-toolchain references (go build, go test, go fmt, go vet). Single coarse work_item — one file, one contract surface.

## Work items

### WI-1 — docs: rewrite CONTRIBUTING.md with project-aware Go guidance

**Complexity:** trivial

Replace the existing repo-root CONTRIBUTING.md with fresh, project-specific contributor guidance for the orch-test-target Go grep utility. Per the spec's 'User-Visible Surface' and 'Normative content template' sections, the file must satisfy hard constraints on location, length, format, and content sections, and should reference the actual Go toolchain commands contributors will use. This is the sole observable deliverable for the spec; no other files are touched.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root and `wc -l CONTRIBUTING.md` reports a value ≤ 59.
- The file contains exactly these top-level `##` section headings, in order: `Getting Started`, `Reporting Issues`, `Submitting Changes`, `Commit Messages`, `Code Style`; and it references `go fmt`, `go vet`, and `go test` somewhere in the body.
- `grep -E '!\[.*\]\(http' CONTRIBUTING.md` produces no output (no badge/shield image markup present).
- `git diff --name-only` against the base branch lists only `CONTRIBUTING.md` — no Go sources, `go.mod`, `README.md`, `.github/`, or `.orch/` files are modified.

