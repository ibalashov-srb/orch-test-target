# Plan

Introduce a minimal Go module at the repo root and a self-contained `cmd/grep` CLI that scans a file or stdin for regexp matches, plus in-process tests and a README usage section. Work is split into four coarse contract surfaces: module scaffolding, the grep CLI/`runGrep` behavior, the table-driven test suite, and the README documentation update.

## Work items

### WI-1 — chore(gomod): initialize module github.com/ibalashov-srb/orch-test-target on Go 1.22

**Complexity:** trivial

Create a top-level go.mod declaring the module path and Go 1.22 toolchain so the repository becomes a valid Go module that `go test ./...` can resolve. The spec's 'go.mod' section mandates this exact module path with no require stanzas (standard library only, no go.sum). This is the foundation that makes every other work_item buildable and testable.

**Acceptance:**

- `go.mod` exists at the repo root with first line `module github.com/ibalashov-srb/orch-test-target`
- `go.mod` declares `go 1.22`
- `go.mod` contains no `require` directives and no `go.sum` is added to the repo
- `go mod verify` succeeds against the committed `go.mod`
- `go list ./...` runs without error from the repo root

### WI-2 — feat(grep): add cmd/grep CLI with testable runGrep entry point

**Complexity:** small

Create `cmd/grep/main.go` implementing the grep-style scanner described in the spec's 'cmd/grep/main.go' section. Expose a `runGrep(pattern, filePath string, in io.Reader, out io.Writer) int` function that returns the documented exit codes, with `main()` as a thin wrapper around `flag.Parse()` and `os.Exit(runGrep(...))`. Uses only the standard library (`bufio`, `fmt`, `io`, `os`, `regexp`, `flag`).

**Acceptance:**

- `go build ./cmd/grep` produces a binary with no external dependencies (only stdlib imports in the source)
- Invoking the binary as `grep <pattern> <file>` prints every matching line as `<filename>:<lineno>:<line>` with 1-based line numbers
- Invoking the binary with no file argument, or with `-` as the file argument, reads from stdin and prints `<lineno>:<line>` with no filename prefix
- Process exits 0 when at least one line matched, exits 1 when input was read successfully but no line matched
- Process exits 2 when the pattern fails `regexp.Compile`, when the file path cannot be opened, or when no pattern argument is supplied; an explanatory message is written to stderr in each error case
- `runGrep` is callable from another file in `package main` with the signature `(pattern, filePath string, in io.Reader, out io.Writer) int`

### WI-3 — test(grep): add table-driven runGrep tests covering match, no-match, errors, stdin

**Complexity:** small

Create `cmd/grep/main_test.go` in `package main` with a `TestRunGrep` table-driven test that exercises `runGrep` in-process via `strings.NewReader` and `strings.Builder`, plus a `tmpFile` helper that registers `t.Cleanup`. The five required scenarios from the spec's test table must each be a named sub-test asserting both the integer exit code and the expected stdout content (or emptiness).

**Acceptance:**

- `go test ./cmd/grep` passes with at least the five sub-tests: `single_file_match`, `no_match`, `invalid_regex`, `missing_file`, `stdin_match`
- `single_file_match` writes a 3-line temp file, runs `runGrep` against it, asserts exit code 0 and that captured stdout contains `<tmpfile-path>:<lineno>:<matching-line>`
- `no_match` runs against a temp file with no matching line, asserts exit code 1 and empty captured stdout
- `invalid_regex` calls `runGrep` with pattern `"[invalid"` and asserts exit code 2
- `missing_file` calls `runGrep` with a non-existent path and asserts exit code 2 with empty captured stdout
- `stdin_match` calls `runGrep` with `filePath==""` and `in=strings.NewReader(...)`, asserts exit code 0 and that captured stdout contains `1:<matching-line>` with no filename prefix
- `go test ./...` from the repo root passes

### WI-4 — docs(readme): document grep tool usage and exit codes

**Complexity:** trivial

Append a `## grep` section to `README.md` describing how to invoke the new tool and what its exit codes mean, as specified in the spec's 'README.md update' section. The section must include a fenced code block with the `go run ./cmd/grep <pattern> [file|-]` invocation form and an explicit enumeration of the 0/1/2 exit-code contract.

**Acceptance:**

- `README.md` contains a new `## grep` heading appended after the existing content (existing lines are preserved verbatim)
- The section contains a fenced code block showing `go run ./cmd/grep <pattern> [file|-]`
- The section explains the output format: `filename:lineno:` for file input, `lineno:` for stdin input
- The section enumerates all three exit codes: 0 = match found, 1 = no match, 2 = usage/file error

## Dependencies

- WI-2 depends on WI-1
- WI-3 depends on WI-1
- WI-3 depends on WI-2
