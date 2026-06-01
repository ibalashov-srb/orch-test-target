# Plan

The repository already contains a compliant CONTRIBUTING.md (33 lines, plain Markdown, no badges). Since the prior Code stage could not finish without a commit on a pure no-op, this plan reframes the work as a single minimal, spec-compliant edit to CONTRIBUTING.md — appending a short "Questions" pointer line — that produces a real commit while preserving all three constraints (plain Markdown, no badges, ≤60 lines).

## Work items

### WI-1 — docs: append minimal Questions pointer to CONTRIBUTING.md

**Complexity:** trivial

Make a minimal, spec-compliant edit to the existing root-level CONTRIBUTING.md so the Code stage can produce exactly one valid commit while keeping the file plain Markdown, badge-free, and well under 60 lines. The spec's three structural constraints (plain Markdown, no badges, ≤60 lines) must continue to hold after the edit. Per the replan context, this resolves the Code stage's contradiction between the no-op spec and the mandatory-commit rule by choosing option (b): a trivial but real addition. Scope is strictly limited to CONTRIBUTING.md — no other file may be touched (see spec §Out of Scope).

**Acceptance:**

- CONTRIBUTING.md at the repository root exists and is the only file changed by the commit.
- `wc -l CONTRIBUTING.md` reports a line count ≤ 60 after the change.
- `grep -E '\[!\[' CONTRIBUTING.md` returns no matches (no badge syntax present).
- The file contains no raw HTML tags and no non-standard Markdown extensions — only headings, paragraphs, and bullet lists.

