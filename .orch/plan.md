# Plan

Add a named-export `subtract(a, b)` to `src/math.js` returning `a - b`, and a matching `node:test` block in `test/math.test.js` that asserts `subtract(5, 3) === 2`. Both edits are tiny, land in two existing files, and form a single cohesive deliverable, so they are grouped as one work item per TD-134.

## Work items

### WI-1 — feat(math): add subtract function with node:test coverage

**Complexity:** trivial

Append a named-export `subtract(a, b)` to `src/math.js` that returns `a - b`, mirroring the existing `add` export pattern (ES modules, named export, no input validation). Extend the existing import line in `test/math.test.js` to also pull in `subtract`, and append one `test(...)` block using `assert.equal` to verify `subtract(5, 3) === 2`. This brings arithmetic parity with `add` as called out in the spec's 'What and Why' and 'User-visible / Consumer-visible Surface' sections; the production change and its test are a single cohesive contract surface and ship together.

**Acceptance:**

- `src/math.js` exports a named function `subtract` such that `import { subtract } from './src/math.js'` followed by `subtract(5, 3)` returns `2`, and `subtract(-1, -1)` returns `0`.
- `test/math.test.js` contains a single `import { add, subtract } from "../src/math.js";` line (no duplicate import statement) and a new `test("subtract returns the difference", ...)` block that asserts `subtract(5, 3) === 2` via `assert.equal`.
- Running `npm test` (i.e., `node --test`) on a clean checkout exits 0 and reports both the pre-existing `add returns the sum` test and the new `subtract returns the difference` test as passing.
- No files other than `src/math.js` and `test/math.test.js` are modified; `package.json`, `README.md`, `CONTRIBUTING.md`, and `go.mod` are unchanged.

