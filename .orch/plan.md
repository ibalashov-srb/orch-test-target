# Plan

Single-file deliverable: overwrite `CONTRIBUTING.md` at the repo root with a clean, compliant version that preserves the existing four sections (Reporting Issues, Submitting Changes, Commit Messages, Code Style) plus the README cross-reference, while honoring the spec's constraints (≤60 lines, plain Markdown, no badges, no HTML, no front-matter). One work item is sufficient — there is no second contract surface and no other file in the repo references CONTRIBUTING.md.

## Work items

### WI-1 — docs: add CONTRIBUTING.md at repo root

**Complexity:** trivial

Author `CONTRIBUTING.md` at the repository root meeting the spec's format contract (see SpecDoc § User-Visible Surface and § Required Sections). The file is a clean overwrite of the existing 35-line file, retaining the four-section structure (Reporting Issues, Submitting Changes, Commit Messages, Code Style) and the introductory pointer to `README.md`. Constraints from the spec — ≤60 lines, plain GitHub-Flavored Markdown, no badges, no HTML, no front-matter — must hold for the committed file.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repo root, UTF-8 encoded with Unix line endings, and `wc -l CONTRIBUTING.md` reports a value ≤ 60.
- The file contains exactly these four H2 section headers in order: `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style`, and an intro paragraph referencing `README.md` precedes them.
- `grep -F '![' CONTRIBUTING.md` returns no matches and `grep -E '<[a-zA-Z/!]' CONTRIBUTING.md` returns no matches (no badge image syntax, no HTML tags, no YAML/TOML front-matter delimiters at the top).
- Each of the four sections contains the substantive guidance called out in the spec: search-first/repro-steps/environment for Reporting Issues; fork → feature branch off `main` → small focused PR for Submitting Changes; imperative mood, ≤72-char subject, `#N` issue references for Commit Messages; match surrounding style, minimal diffs, run linters for Code Style.

