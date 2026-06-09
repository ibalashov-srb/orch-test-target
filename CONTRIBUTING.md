# Contributing

Thank you for your interest in contributing! This guide explains how to report
issues, propose changes, and keep the codebase consistent. Please read
`README.md` first to understand the project's purpose and local setup.

## Reporting Issues

- Search existing issues before opening a new one to avoid duplicates.
- Provide a clear, minimal reproduction case.
- Include your OS, Go version (`go version`), and any relevant error output.

## Submitting Changes

- Fork the repository and create a feature branch off `main`.
- Keep each pull request focused on a single concern.
- Ensure all tests pass locally before opening the PR.
- Open a pull request against `main` and fill in the PR description explaining
  what changed and why.

## Commit Messages

- Write commit subjects in the imperative mood (e.g. "Add", "Fix", "Update").
- Keep the subject line to 72 characters or fewer.
- Reference related issues with `#N` (e.g. `Fixes #42`) when applicable.
- Add a blank line between the subject and any body text.

## Code Style

- Format all Go code with `go fmt ./...` before committing.
- Check for common issues with `go vet ./...` and resolve any findings.
- Match the style of the surrounding code; consistency matters.
- Avoid unrelated reformatting in the same commit as functional changes.
