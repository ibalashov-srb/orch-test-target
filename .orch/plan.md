# Plan

Ensure `CONTRIBUTING.md` exists at the repository root containing the canonical 32-line plain-Markdown content (Introduction, Reporting Issues, Submitting Changes, Commit Messages, Code Style sections) and satisfying the structural constraints (≤60 lines, no image/badge syntax). The file already matches the spec's canonical content verbatim, so the Code stage's role is to verify and commit/no-op as required by the pipeline. Single coarse work item — one file, one contract surface.

## Work items

### WI-1 — docs: add CONTRIBUTING.md at repository root

**Complexity:** trivial

Provide a `CONTRIBUTING.md` at the repo root matching the canonical content defined in the spec's "Canonical File Content" section. The file documents how to report issues, submit pull requests, write commit messages, and follow code style. It must remain plain GitHub-flavored Markdown with no image or shield/badge syntax and stay within the 60-line cap defined under "Acceptance Criteria". Since the file already exists with compliant content, the implementation verifies it is present and byte-identical to the canonical text, writing it only if missing or divergent.

**Acceptance:**

- `CONTRIBUTING.md` exists at the repository root and its line count is ≤ 60.
- The file contains the five required H1/H2 sections in this order: top-level `# Contributing` introduction, `## Reporting Issues`, `## Submitting Changes`, `## Commit Messages`, `## Code Style`.
- The file contains no Markdown image syntax (`![...](...)`) and no raw HTML `<img>` or shield/badge elements (verifiable by grep).
- File content is byte-for-byte identical to the canonical Markdown block in the spec's "Canonical File Content" section.

