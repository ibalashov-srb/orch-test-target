# Plan

Add a single CONTRIBUTING.md file at the repository root containing the conventional sections (Reporting Issues, Submitting Changes, Commit Messages, Code Style) in plain Markdown, no badges, strictly under 60 lines. Since this is a one-file, no-dependency change, the plan consists of a single bead.

## Beads

### b1 — docs: add CONTRIBUTING.md at repository root

**Complexity:** trivial

Create /CONTRIBUTING.md (alongside README.md) with a concise contributor guide in plain GitHub-Flavored Markdown. Includes a top-level '# Contributing' heading, a brief welcome paragraph referencing the README, and the sections '## Reporting Issues', '## Submitting Changes', '## Commit Messages', and '## Code Style' in that order, using ATX headings and hyphenated bullet lists. The file must contain no badges or inline images, no front-matter, and no HTML, ending in a single trailing newline. This is the only file added or modified; no other repo files are touched.

**Acceptance:**

- File `CONTRIBUTING.md` exists at the repository root.
- `wc -l CONTRIBUTING.md` returns a value of 59 or less (strictly under 60 lines).
- File contains, in order: a top-level `# Contributing` heading, a welcome paragraph mentioning the README, and `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, and `## Code Style` sections using ATX-style headings.
- `## Submitting Changes` covers fork → feature branch → descriptive commit → PR against `main`, and mentions that small, focused PRs are preferred.
- `## Commit Messages` includes 2–3 bullets covering imperative mood, 72-character subject limit, and referencing issue numbers.
- File contains no badge or inline-image Markdown (no `![...](...)` syntax) and no HTML blocks or YAML front-matter; unordered lists use hyphens.
- No other files in the repository are created, modified, or deleted (README.md, .gitignore, .orch/, .smoke-probe unchanged).

