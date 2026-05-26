# Contributing

Thanks for helping improve `orch-test-target`, a small Go CLI project that
currently provides a minimal `grep`-like command. Start with `README.md` for
the project overview, usage, output format, and exit-code behavior.

## Prerequisites

- Go 1.22 or newer.
- A working shell for running the validation commands below.

## Making changes

- Keep changes small and focused.
- Match the style of the surrounding Go code and documentation.
- Update `README.md` when user-facing behavior changes.
- Prefer clear commit messages that describe the intent of the change.

## Validate before submitting

Run these commands from the repository root:

```sh
go test ./...
go run ./cmd/grep <pattern> [file|-]
```

Use the `go run` command to manually confirm CLI behavior when your change
affects input handling, output, or exit codes.
