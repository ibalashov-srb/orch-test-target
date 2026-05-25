# Contributing

Thanks for your interest in improving this project! Please read `README.md`
first to get familiar with the project's purpose and setup before contributing.

## Reporting Issues

- Search existing issues first to avoid filing duplicates.
- Provide clear reproduction steps so maintainers can confirm the behavior.
- Include environment details: your OS and Go version (`go version`).

## Submitting Changes

- Fork the repository to your own account.
- Create a feature branch off `main` for your work.
- Write descriptive commits that explain the intent of each change.
- Open a pull request against `main` when your branch is ready for review.
- Prefer small, focused PRs over large, sprawling ones — they are easier to
  review and merge.

## Commit Messages

- Use the imperative mood in the subject line (e.g., "Add", "Fix", "Update").
- Keep the subject line to 72 characters or fewer.
- Reference related issues using `#N` (e.g., `Fixes #42`) when applicable.

## Code Style

- Run `go fmt ./...` before submitting to keep formatting consistent.
- Run `go vet ./...` before submitting to catch common correctness issues.
- Match the style of the surrounding code; consistency matters more than
  personal preference.
- Minimize unrelated changes — avoid reformatting code outside your diff.
