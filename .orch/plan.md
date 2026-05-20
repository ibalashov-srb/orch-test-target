# Plan

Add a single new `CONTRIBUTING.md` at the repository root following the spec's content and formatting contract (H1 + four ordered H2 sections, hyphen lists, no badges/HTML, ≤59 lines, single trailing newline). No other files are touched. The scope is narrow enough that one coarse work item covers the entire spec surface.

## Work items

### WI-1 — docs: add CONTRIBUTING.md at repository root

**Complexity:** trivial

Create `/CONTRIBUTING.md` covering issue reporting, PR workflow, commit-message conventions, and code style, as required by the spec's 'User-Visible Surface' and 'Section content requirements' sections. The file must conform to the spec's strict formatting contract (ATX headings, hyphen lists, no badges/HTML/front-matter, ≤59 lines, single trailing LF newline, UTF-8). No other files in the repo (`README.md`, `.gitignore`, `.smoke-probe`, `.orch/`) may be modified, per 'Implementation Notes' and 'Out of Scope'.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root; no other tracked files are added, modified, or deleted.
- File begins with `# Contributing` followed by a welcome paragraph that mentions `README.md` by name.
- File contains exactly four H2 sections in this order: `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style`.
- `## Reporting Issues` lists searching existing issues, providing reproduction steps, and including environment details; `## Submitting Changes` describes the fork → feature branch → descriptive commit → PR-against-`main` workflow and notes a preference for small, focused PRs; `## Commit Messages` contains three bullets covering imperative mood, ≤72-char subject, and `#N` issue references; `## Code Style` is a brief note to follow existing conventions.
- All unordered list items use `-` markers; no `*` or `+` list markers appear; no `![…](…)` image/badge syntax, raw HTML blocks, or YAML front-matter are present.
- `wc -l CONTRIBUTING.md` reports a value ≤ 59, the file is valid UTF-8 with LF line endings, and it ends with exactly one trailing newline (no blank trailing lines).

