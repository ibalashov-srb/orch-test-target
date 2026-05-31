# Plan

Add a new stdlib-only `strutil` package at `strutil/` exposing four string utilities (`Reverse`, `IsPalindrome`, `CountVowels`, `Capitalize`), each in its own source file, with a single table-driven `strutil_test.go` covering the edge cases enumerated in the spec (empty strings, multi-byte Unicode, mixed case, spacing). No changes outside `strutil/` and no `go.mod` edits — this is a single cohesive deliverable kept as one work_item per TD-134.

## Work items

### WI-1 — feat(strutil): add Reverse, IsPalindrome, CountVowels, Capitalize package

**Complexity:** small

Create a new package `strutil` under `strutil/` in module `github.com/ibalashov-srb/orch-test-target`, with one file per exported function (`reverse.go`, `ispalindrome.go`, `countvowels.go`, `capitalize.go`) and a single table-driven `strutil_test.go`. Functions follow the signatures and semantics in the spec's 'User-Visible Surface' table: rune-correct reversal, space-stripping case-insensitive palindrome check (reusing `Reverse`), Latin-vowel count, and first-rune-of-each-space-token capitalisation. Implementation uses only the standard library (`strings`, `unicode`); no `go.mod` changes.

**Acceptance:**

- Package `strutil` exists at `strutil/` with files `reverse.go`, `ispalindrome.go`, `countvowels.go`, `capitalize.go`, each exporting the function named in the spec's signature table with the documented behaviour (e.g. `Reverse("日本語") == "語本日"`, `IsPalindrome("A man a plan a canal Panama") == true`, `CountVowels("héllo") == 1`, `Capitalize("hELLO wORLD") == "HELLO WORLD"`).
- `strutil/strutil_test.go` contains one `TestXxx` per exported function, each driven by a `[]struct{...}` table that includes every named row listed in the spec's four test tables (empty, single, ASCII, Unicode multi-rune, mixed case, leading-space, etc.).
- `go test ./...` from the repo root exits 0 with all `strutil` test cases passing, and `gofmt -l strutil/` prints nothing.
- `go.mod` is unchanged (still declares only the module path and `go 1.22`, no `require` block added), and no files outside `strutil/` are modified.

