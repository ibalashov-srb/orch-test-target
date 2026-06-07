# Plan

Add a `multiply(a, b)` named ESM export to `src/math.js` mirroring the existing `add` function's style, and extend `test/math.test.js` with a matching `node:test` assertion. This is a single cohesive deliverable touching one production file and its test, so it ships as one work item.

## Work items

### WI-1 — feat(math): add multiply function with test

**Complexity:** trivial

Append a `multiply(a, b)` function declaration to `src/math.js` as a named ESM export returning `a * b`, matching the style of the existing `add` function. Extend the destructured import in `test/math.test.js` to include `multiply` and append a `node:test` block asserting `multiply(3, 4) === 12` via `assert.equal`. Spec sections covered: 'User-Visible Surface' (both files) and 'Implementation Notes' (ESM, node --test, function-declaration style).

**Acceptance:**

- `src/math.js` exports a named `multiply` function declared with the `function` keyword whose body is `return a * b;`, appearing after the existing `add` export.
- `test/math.test.js` imports `multiply` alongside `add` from `../src/math.js` in a single named-import statement and contains a `test(...)` block asserting `multiply(3, 4)` equals `12` using `assert.equal`.
- Running `npm test` exits 0 with both the existing `add` test and the new `multiply` test reported as passing by `node --test`.
- No files outside `src/math.js` and `test/math.test.js` are modified.

