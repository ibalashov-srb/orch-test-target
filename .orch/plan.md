# Plan

Replace the existing root `CONTRIBUTING.md` with a freshly authored, plain-Markdown contributor guide that satisfies the spec's constraints (under 60 lines, no badges/images) and includes the six required sections in order. Single-file scope — no other repo files are touched.

## Work items

### WI-1 — docs: author root CONTRIBUTING.md per spec

**Complexity:** trivial

Create/overwrite `CONTRIBUTING.md` at the repository root with a freshly authored contributor guide covering the six required sections in order: Title (`# Contributing`), opening paragraph pointing at `README.md`, Reporting Issues, Submitting Changes, Commit Messages, and Code Style (see spec §Required Sections). The file must be plain Markdown with ATX headings, `-` bullets, no badges, no images, no HTML, no fenced code blocks, and must fit under the 60-line ceiling with a single trailing newline (see spec §Implementation Notes). This delivers the single user-visible artifact the spec calls for.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root, begins with `# Contributing`, ends with exactly one trailing newline, and totals fewer than 60 lines when counted by `wc -l`.
- The file contains, in order, the six sections from the spec: title, opening paragraph referencing `README.md`, `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style`; each non-title section uses `-` bullet lists covering the bullet points listed in spec §Required Sections (e.g. imperative mood, 72-char subject, `#N` references under Commit Messages).
- The file contains no badge or image syntax: `grep -E '!\[|<img'` against `CONTRIBUTING.md` returns no matches; it also contains no fenced code blocks (no ``` ``` ```), no HTML tags, and no Markdown tables.
- No file other than `CONTRIBUTING.md` is modified by this work item (e.g. `README.md` is unchanged, no files added under `.github/`).

