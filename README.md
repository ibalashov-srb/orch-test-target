# orch-test-target
Smoke-test target for the orchestrator pipeline

## grep

A minimal `grep`-like utility that searches for a pattern in a file or in
standard input and prints matching lines.

```
go run ./cmd/grep <pattern> [file|-]
```

If `[file|-]` is omitted or given as `-`, input is read from standard input.

### Output format

- When the input is a file, each matching line is printed as
  `filename:lineno:<line>`, where `filename` is the path passed on the
  command line and `lineno` is the 1-based line number within that file.
- When the input is standard input, each matching line is printed as
  `lineno:<line>` (no filename prefix).

### Exit codes

- `0` — at least one matching line was found.
- `1` — the input was read successfully but no line matched the pattern.
- `2` — a usage error (e.g. missing pattern, too many arguments) or a
  file error (e.g. the named file could not be opened or read).
