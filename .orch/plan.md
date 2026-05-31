# Plan

Introduce a new self-contained `sliceutil` Go package at `sliceutil/` exposing three pure int-slice helpers (`Sum`, `Max`, `Contains`) — each in its own file — backed by a single table-driven test file. No module or existing-package changes; `go test ./...` must continue to pass alongside the existing `cmd/grep` tests.

## Work items

### WI-1 — feat(sliceutil): add Sum, Max, Contains package with table-driven tests

**Complexity:** small

Create the new `sliceutil` package under `sliceutil/` in module `github.com/ibalashov-srb/orch-test-target`, providing the three int-slice helpers specified in the spec's 'Exported API' section: `Sum(xs []int) int`, `Max(xs []int) (int, bool)` (returns `0, false` on empty input), and `Contains(xs []int, v int) bool`. Each function lives in its own file (`sum.go`, `max.go`, `contains.go`) per the spec's 'File Layout' section, with a single white-box table-driven test file `sliceutil_test.go` covering the cases enumerated under 'Implementation Notes → Test file'. No imports beyond the standard library `testing`; `go.mod` and `go.sum` remain untouched (spec 'go.mod / go.sum' and 'Out of Scope').

**Acceptance:**

- `go test ./sliceutil/...` passes, exercising the documented Sum/Max/Contains cases including empty-slice behaviour (Sum→0, Max→(0,false), Contains→false) and the present/absent/first/last/single-element/mixed-sign scenarios from the spec.
- `go test ./...` passes for the whole module, with the pre-existing `cmd/grep` tests unaffected and `go.mod` / `go.sum` byte-identical to before.
- `gofmt -l sliceutil/` produces no output; the directory contains exactly four files — `sum.go`, `max.go`, `contains.go`, `sliceutil_test.go` — all declaring `package sliceutil`, with each of the three exported functions defined in its correspondingly-named file.
- Importing `github.com/ibalashov-srb/orch-test-target/sliceutil` and calling `Sum`, `Max`, `Contains` with the signatures from the spec compiles successfully (verified by the test file itself, which is white-box `package sliceutil`).

