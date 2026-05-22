# Plan

This is a no-feature smoke run: bump `.smoke-probe` with a fresh monotonically-increasing probe value so the orchestrator exercises the full Code → gofmt → `go test ./...` → commit → push path and Langfuse records a complete trace. The only repo artifact is the updated `.smoke-probe` line; no Go sources, docs, or `.orch/` config are touched.

## Work items

### WI-1 — chore(smoke): bump .smoke-probe to trigger live Langfuse pipeline run

**Complexity:** trivial

Update the single line in `.smoke-probe` from `probe 1779188493` to `probe <new>` where `<new>` is a fresh Unix-epoch-seconds integer (or any integer strictly greater than 1779188493) captured at Code-stage time. This is the minimal non-empty diff required to drive every pipeline stage to completion and produce a full Langfuse trace, per the spec sections 'What and Why', 'User-Visible Surface', and decision q1/q2. No other file in the repo is modified.

**Acceptance:**

- `.smoke-probe` contains exactly one line matching `^probe [0-9]+\n$` after the change.
- The integer in `.smoke-probe` is strictly greater than `1779188493` (the prior probe value).
- `git diff --name-only` against the prior commit lists exactly `.smoke-probe` and no other path.
- `gofmt -l .` produces empty output (no formatting changes) after the edit.
- `go test ./...` exits 0, running the existing `cmd/grep` sub-tests and `doc_test.go` unchanged.
- A commit containing the `.smoke-probe` change is pushed to `main`, producing a Langfuse trace that spans Code, test, and push stages.

