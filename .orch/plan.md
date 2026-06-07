# Plan

The spec resolves the task as already complete: `CONTRIBUTING.md` exists at the repo root, is 30 lines of plain Markdown with no badges, and covers the four expected sections. The plan therefore contains a single verification work_item that confirms the existing file satisfies every constraint from the brief; no file writes are required, and the Decompose stage should produce a no-op or verification-only bead set.

## Work items

### WI-1 — docs: verify CONTRIBUTING.md meets brief constraints

**Complexity:** trivial

Confirm that the existing repo-root `CONTRIBUTING.md` already satisfies every constraint from the brief — plain Markdown, no badges, under 60 lines — so the task can be closed without churn. The spec's resolution of open question q1 explicitly chooses 'no action' over overwriting a correct file (see spec sections 'Open Question Resolution (q1)' and 'Implementation Notes'). This work_item exists to make that verification an auditable deliverable rather than an implicit decision.

**Acceptance:**

- `CONTRIBUTING.md` is present at the repository root (path is exactly `CONTRIBUTING.md`, not nested under `docs/` or `.github/`).
- The file's line count is at most 60 lines, verifiable via `wc -l CONTRIBUTING.md`.
- The file contains no Markdown image syntax (`![`) and no shields.io / badge URLs, verifiable via `grep -E '!\[|shields\.io|badge' CONTRIBUTING.md` returning no matches.
- No tracked files in the repository are modified by this work_item — `git status` is clean after the plan executes.

