# Contributing

Thank you for your interest in contributing to this project — a minimal
`grep`-like utility written in Go 1.22. Please read `README.md` first to
understand the project's purpose and usage before getting started.

## Getting Started

1. Clone the repository and ensure Go 1.22 or later is installed.
2. Build the project to confirm everything compiles:
   ```
   go build ./...
   ```
3. Run the test suite to make sure the baseline is green:
   ```
   go test ./...
   ```

## Reporting Issues

- Search existing issues first to avoid filing duplicates.
- Include clear reproduction steps and the exact command you ran.
- State your OS and the output of `go version`.

## Submitting Changes

- Fork the repository and create a feature branch off `main`.
- Keep changes focused; prefer small PRs over large, sprawling ones.
- Ensure `go build ./...` and `go test ./...` both pass before opening
  a pull request against `main`.

## Commit Messages

- Use the imperative mood in the subject line (e.g. "Add", "Fix", "Update").
- Keep the subject line to 72 characters or fewer.
- Reference related issues with `#N` (e.g. `Fixes #42`) when applicable.

## Code Style

- Format all Go code with `go fmt` before committing.
- Run `go vet ./...` and resolve any warnings before submitting.
- Match the style of the surrounding code; consistency matters most.
