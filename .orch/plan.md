# Plan

Add an exported `HelloWorld() string` function to the `cmd/grep` package in a new `cmd/grep/hello.go`, accompanied by a focused unit test in a new `cmd/grep/hello_test.go`. No existing files (`main.go`, `main_test.go`, `go.mod`, `go.sum`) are touched, and the function is not wired into the CLI — both per the spec's Out-of-Scope section.

## Work items

### WI-1 — feat(cmd/grep): add HelloWorld function with unit test

**Complexity:** trivial

Introduce a new `cmd/grep/hello.go` containing an exported `HelloWorld() string` function in `package main` that returns the literal `"hello, world"`, and a co-located `cmd/grep/hello_test.go` with `TestHelloWorld` that asserts the return value using `t.Errorf` (matching the existing `TestRunGrep` style). The function is intentionally isolated from `main.go`'s grep-scanner logic so it can be reviewed, tested, and removed cleanly — see spec sections 'User-Visible Surface' and 'Why a new file rather than appending to main.go'. No CLI wiring, no changes to `main.go` / `main_test.go`, and no new module dependencies.

**Acceptance:**

- `cmd/grep/hello.go` exists, declares `package main`, and defines exported `func HelloWorld() string` that returns the literal `"hello, world"`.
- `cmd/grep/hello_test.go` exists in `package main` and defines `TestHelloWorld(t *testing.T)` which fails via `t.Errorf` when `HelloWorld()` does not equal `"hello, world"`.
- `go test ./cmd/grep/...` passes, with both `TestRunGrep` and the new `TestHelloWorld` reported as PASS and no other test files modified.
- `cmd/grep/main.go`, `cmd/grep/main_test.go`, `go.mod`, and `go.sum` are unchanged from their pre-change contents (verified by diff).

