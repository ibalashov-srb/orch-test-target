# Plan

The spec resolves the request as already satisfied: a conformant `CONTRIBUTING.md` (29 lines, plain Markdown, no badges) already exists at the repository root. The plan contains a single verification work_item that confirms the existing file meets every stated constraint and records the no-op outcome; no files will be created or modified.

## Work items

### WI-1 — chore(docs): verify existing CONTRIBUTING.md satisfies spec

**Complexity:** trivial

The spec's `What and Why` section concludes that the existing `CONTRIBUTING.md` at the repo root already meets every constraint (exists, <60 lines, plain Markdown, no badges) and that the correct resolution is zero file changes. This work_item records that verification as the deliverable: confirm the file is present and conformant, and close the task with an empty changeset as called for under `Implementation Notes` (`Code stage action: output an empty changeset; mark task done`). No file in the working tree is created, modified, or deleted.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root and `wc -l CONTRIBUTING.md` reports 29 lines (under the 60-line cap from the spec's constraints).
- `grep -iE 'badge|shields\.io|!\[' CONTRIBUTING.md` returns no matches, confirming the plain-Markdown / no-badges constraint.
- `git status` after the plan executes shows no modifications, additions, or deletions under the repository working tree (empty changeset, per `Implementation Notes`).
- The file content is byte-identical to the pre-plan snapshot — no overwrite-with-equivalent-content is performed (rejected alternative in the spec).

