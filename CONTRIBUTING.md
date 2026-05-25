# Contributing

Welcome! We're glad you want to contribute to this Go project.

## Prerequisites

- Go 1.22 or newer installed on your system.
- Clone the repository: `git clone https://github.com/ibalashov-srb/orch-test-target.git`
- Verify your setup: `go build ./...`

## Running Tests

Run the full test suite with:

```
go test ./...
```

Tests live alongside the packages they cover. Key locations include
`cmd/grep/` for the grep command and `doc_test.go` at the repo root.

## Reporting Issues

- Search existing issues before filing a new one to avoid duplicates.
- Include clear reproduction steps so maintainers can confirm the problem.
- State your Go version (`go version`) and operating system version.

## Submitting Changes

- Fork the repository to your own account.
- Create a feature branch off `main` for your work.
- Open a pull request against `main` when your branch is ready for review.
- Prefer small, focused PRs — they are easier to review and merge.

## Commit Messages

- Write subjects in the imperative mood ("Add", "Fix", "Remove").
- Keep the subject line to 72 characters or fewer.
- Reference related issues with `Fixes #N` (e.g., `Fixes #42`) when applicable.

## Code Style

- Format all code with `go fmt ./...` before submitting.
- Run `go vet ./...` and resolve any reported issues.
- Do not include unrelated reformatting in the same commit as functional changes.
