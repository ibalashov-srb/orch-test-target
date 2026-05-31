# Plan

Add a new stdlib-only Go package `sliceutil` (import path `github.com/ibalashov-srb/orch-test-target/sliceutil`) containing three `[]int` helpers — `Sum`, `Max`, and `Contains` — each in its own file (`sum.go`, `max.go`, `contains.go`), accompanied by one white-box table-driven test file (`sliceutil_test.go`). All files must be `gofmt`-clean and `go test ./...` must pass without modifying any existing file (go.mod, doc_test.go, cmd/grep, etc.).

## Work items

### WI-1 — feat(sliceutil): add Sum, Max, Contains package with table-driven tests

**Complexity:** small

Create the new `sliceutil/` directory with three production files (`sum.go`, `max.go`, `contains.go`) each declaring `package sliceutil` and one exported function with the signatures fixed in the spec's 'User-Visible Surface' table, plus the white-box table-driven test file `sliceutil_test.go` covering all three. The package uses only the standard library and ships as a single cohesive deliverable (one production-file-per-function + its tests) per the spec's 'File Layout' and 'Implementation Notes' sections. No existing file — go.mod, doc_test.go, README.md, cmd/grep/* — is modified.

**Acceptance:**

- `sliceutil/sum.go`, `sliceutil/max.go`, `sliceutil/contains.go`, and `sliceutil/sliceutil_test.go` exist; each declares `package sliceutil`; signatures exactly match `Sum(xs []int) int`, `Max(xs []int) (int, bool)`, and `Contains(xs []int, v int) bool`.
- Behaviour matches the spec: `Sum(nil)==0`, `Sum([]int{-1,-2,3})==0`, `Sum([]int{2,2,2})==6`; `Max(nil)==(0,false)`, `Max([]int{-3,-1,-2})==(-1,true)`; `Contains(nil,1)==false`, `Contains([]int{-1,0,1},-1)==true`, `Contains([]int{1,2,3},9)==false`.
- `sliceutil_test.go` defines `TestSum`, `TestMax`, `TestContains` as table-driven tests using `t.Run` sub-tests and `t.Errorf` (not `t.Fatalf`) on mismatch, with the case coverage listed in the spec's 'Test file' section (nil, empty, single, mixed signs, duplicates / first-max / last-max / all-equal / negatives / present / absent / first / last).
- `gofmt -l sliceutil/` produces no output, and `go test ./...` from the repo root exits 0 with the existing `cmd/grep` and root `doc_test.go` tests still passing alongside the new `sliceutil` tests.
- No file outside `sliceutil/` is created or modified — `git diff --name-only main` lists only paths under `sliceutil/`.


## Quality warnings

- WI-1 has 5 acceptance bullets — 5+ usually means two WIs sharing a file; consider splitting at the Plan gate
