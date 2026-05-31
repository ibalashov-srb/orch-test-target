# Plan

Add a new exported `Multiply(a, b int) int` helper to the root `orchtest` package together with its table-driven `TestMultiply`. Both files are net-new at the repo root; no existing file is touched. Per TD-134, the helper and its colocated test ship as a single cohesive work item.

## Work items

### WI-1 — feat(orchtest): add Multiply integer helper with table-driven test

**Complexity:** trivial

Introduces the exported `Multiply(a, b int) int` function in the root `orchtest` package, returning `a * b` with native Go int semantics, alongside a table-driven `TestMultiply` covering positive, negative, mixed-sign, zero (both positions), identity, and large-value cases. Both files (`multiply.go` and `multiply_test.go`) are new at the repo root and share `package orchtest` for consistency with the existing `doc_test.go`. Together they form one cohesive deliverable per the spec sections 'User-Visible Surface' and 'Files to Create'.

**Acceptance:**

- `multiply.go` exists at repo root declaring `package orchtest` and exporting `func Multiply(a, b int) int` with a doc comment, whose body returns `a * b` and requires no imports.
- `multiply_test.go` exists at repo root declaring `package orchtest` and defining `TestMultiply`, a table-driven test using an anonymous struct slice plus `t.Run` sub-tests that covers all seven named cases from the spec (both_positive=12, both_negative=12, mixed_sign=-12, zero_times_value=0, value_times_zero=0, identity=99, large_values=999000).
- `go test ./...` passes from a clean checkout with all seven `TestMultiply/<name>` sub-tests reported as PASS, and `gofmt -l multiply.go multiply_test.go` prints no output.
- No file other than `multiply.go` and `multiply_test.go` is added or modified (in particular `doc_test.go`, `go.mod`, and everything under `cmd/grep/` are byte-identical to main).

