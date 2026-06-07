import test from "node:test";
import assert from "node:assert/strict";
import { add, multiply } from "../src/math.js";

test("add returns the sum", () => {
  assert.equal(add(2, 3), 5);
});

test("multiply returns the product", () => {
  assert.equal(multiply(3, 4), 12);
});
