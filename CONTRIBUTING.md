# Contributing

Welcome, and thanks for your interest in improving this project! Before diving in, please skim the [README](README.md) to get a feel for the project's purpose and scope. This guide collects the conventions we use so your contributions can land smoothly.

## Reporting Issues

- Search existing issues first to avoid duplicates.
- Open a new issue with a clear, descriptive title and a minimal reproduction or concrete example.
- Include relevant environment details (OS, runtime version, commit SHA) and the expected vs. actual behavior.
- For security concerns, prefer a private channel over a public issue.

## Submitting Changes

- Fork the repository to your own account.
- Create a feature branch off `main` (for example, `git checkout -b fix/typo-in-readme`).
- Make focused commits with descriptive messages as you go.
- Push your branch and open a pull request against `main` in this repository.
- Small, focused PRs are strongly preferred over large, sweeping changes — they are easier to review and more likely to be merged quickly.
- Ensure any tests or checks pass locally before requesting review.

## Commit Messages

- Write the subject in the imperative mood (for example, "Add foo" rather than "Added foo" or "Adds foo").
- Keep the subject line under 72 characters; wrap longer explanations in the body.
- Reference relevant issue numbers in the body or footer (for example, `Refs #123` or `Fixes #123`).

## Code Style

- Match the style of the surrounding code; do not reformat unrelated lines.
- Prefer clear names and small functions over clever one-liners.
- Keep changes minimal and on-topic for the PR; refactors belong in their own commits.
- Run any formatters or linters configured in the repo before submitting.
