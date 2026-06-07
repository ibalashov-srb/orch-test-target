# Plan

Add a `multiply(a, b)` named export to `src/math.js` mirroring the style of the existing `add` export, and extend `test/math.test.js` with a matching `node:test` case. Single coarse work item — one production file plus its co-located test form one cohesive deliverable.

## Work items

### WI-1 — feat(math): add multiply(a, b) export with test

**Complexity:** trivial

Extend the `src/math.js` utility module with a new named export `multiply(a, b)` that returns `a * b`, matching the implementation style of the existing `add` export (no coercion, no guards). Add a corresponding `node:test` case in `test/math.test.js` by appending `multiply` to the existing destructured import and adding one new `test(...)` block, leaving the existing `add` test untouched. Covers the entirety of the spec's User-Visible Surface section.

**Acceptance:**

- `src/math.js` exports `multiply` such that `import { multiply } from './src/math.js'` followed by `multiply(3, 4)` returns `12`, and the existing `add` export is unchanged.
- `test/math.test.js` imports both `add` and `multiply` from `../src/math.js` on a single import line and contains a `test("multiply returns the product", ...)` block asserting `assert.equal(multiply(3, 4), 12)`.
- Running `node --test` from the repo root exits 0 with both the pre-existing `add` test and the new `multiply` test reported as passing.

