# Plan

Add a new self-contained `mathx` Go package exposing `GCD`, `LCM`, `IsPrime`, and `Factorial` with a single table-driven test file. The change is cohesive and additive — four small source files plus one test file under `mathx/`, no edits to existing files — so it ships as one work_item per TD-134 (single small package, code + tests stay together).

## Work items

### WI-1 — feat(mathx): add GCD, LCM, IsPrime, Factorial package

**Complexity:** small

Create a new `mathx` package at module path `github.com/ibalashov-srb/orch-test-target/mathx` with the four functions defined in the spec's Public API Surface section, each in its own source file (`gcd.go`, `lcm.go`, `isprime.go`, `factorial.go`), plus a single white-box table-driven `mathx_test.go` covering the edge-case tables enumerated in the spec. The package is purely additive — no existing file (`go.mod`, `doc_test.go`, `cmd/grep/*`, README, CONTRIBUTING) is modified, and only Go standard-library imports (`math`, `errors`, `testing`) are used. Implementations follow the spec's Implementation Notes (iterative Euclidean GCD, zero-guarded LCM, sqrt-bounded trial-division IsPrime, iterative Factorial returning a non-nil error for negative inputs).

**Acceptance:**

- `go test ./mathx/...` passes on a clean checkout, exercising every row in the spec's TestGCD / TestLCM / TestIsPrime / TestFactorial tables (including GCD(0,0)=0, LCM(x,0)=0, IsPrime(2)=true, IsPrime(n<2)=false, Factorial(0)=1, and Factorial(-1) returning a non-nil error).
- `gofmt -l ./mathx/` produces no output and `go vet ./mathx/...` is clean.
- The four source files `mathx/gcd.go`, `mathx/lcm.go`, `mathx/isprime.go`, `mathx/factorial.go` each declare `package mathx` and export exactly the signature given in the spec (`GCD(a, b int) int`, `LCM(a, b int) int`, `IsPrime(n int) bool`, `Factorial(n int) (int, error)`).
- `git diff --name-only main` shows only files under `mathx/` added; `go.mod`, `go.sum`, `doc_test.go`, and `cmd/grep/**` are unchanged.

