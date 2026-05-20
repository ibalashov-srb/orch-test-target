# Plan

Add a new top-level CONTRIBUTING.md that gives contributors a concise guide for issues, PRs, commit style, and code style. The change is purely additive — no existing files are modified, no tooling or CI changes — and the file must satisfy strict formatting constraints (≤59 lines, plain Markdown, four H2 sections in a fixed order, `-` list markers only, no badges/HTML/front-matter).

## Work items

### WI-1 — docs: add CONTRIBUTING.md at repository root

**Complexity:** trivial

Create a new CONTRIBUTING.md at the repo root that documents how to report issues and submit changes to orch-test-target, per the approved spec. The file must follow the mandated structure (H1 welcome paragraph referencing README.md, then exactly four H2 sections in order: Reporting Issues, Submitting Changes, Commit Messages, Code Style) and satisfy the formatting constraints (≤59 lines, `-` list markers only, no badges, no raw HTML, no YAML front-matter, UTF-8 with LF endings and exactly one trailing newline). No existing files are touched.

**Acceptance:**

- File `CONTRIBUTING.md` exists at the repository root, UTF-8 encoded with LF line endings and exactly one trailing newline; `wc -l CONTRIBUTING.md` reports ≤ 59.
- The document begins with an H1 `# Contributing` followed by a welcome paragraph that mentions `README.md` by name.
- The document contains exactly four H2 sections in this order: `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style` (no other H2 headings).
- `## Reporting Issues` instructs contributors to search existing issues first, provide reproduction steps, and include environment details (OS and runtime/language version).
- `## Submitting Changes` describes the workflow: fork the repo, create a feature branch, write descriptive commits, open a PR against `main`, and prefer small focused PRs.
- `## Commit Messages` contains exactly three `-` bullets covering imperative mood, ≤72-character subject line, and issue references using `#N` (e.g. `Fixes #42`).
- The file uses only `-` list markers (no `*` or `+`), contains no badge/image syntax `![…](…)`, no raw HTML tags, and no YAML front-matter; `README.md`, `.gitignore`, and `.smoke-probe` are unchanged.

