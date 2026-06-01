# Plan

The repository root already contains a CONTRIBUTING.md (33 lines, plain Markdown, no badges) that satisfies all three constraints from the spec. The plan is a single verification work_item: confirm the file is compliant and emit a no-op, or apply a minimal fixup if the file has drifted out of compliance by execution time.

## Work items

### WI-1 — docs: verify CONTRIBUTING.md meets plain-markdown, no-badge, <=60-line contract

**Complexity:** trivial

Confirm that CONTRIBUTING.md at the repository root satisfies the three constraints stated in the spec's Acceptance Criteria table: plain Markdown (no HTML or non-standard extensions), no badge image-link patterns, and at most 60 lines. The spec's primary path is a no-op because the existing 33-line file already complies; the fallback path is a minimal edit (trim, remove badge line, or rewrite a non-standard fragment) only if the file has drifted between spec time and execution. No other files may be touched.

**Acceptance:**

- `wc -l CONTRIBUTING.md` at the repository root reports 60 or fewer lines.
- `grep -E '\[!\[' CONTRIBUTING.md` produces no matches (no badge image-link patterns present).
- CONTRIBUTING.md contains no raw HTML tags and no non-standard Markdown extensions (front-matter, custom directives, etc.); it renders as plain Markdown on GitHub.
- No file other than CONTRIBUTING.md is added, modified, or deleted by this work_item.

