import test from "node:test";
import assert from "node:assert/strict";
import { add, subtract } from "../src/math.js";

test("add returns the sum", () => {
  assert.equal(add(2, 3), 5);
});

test("subtract returns the difference of two positive integers", () => {
  assert.equal(subtract(5, 3), 2);
});

test("subtract with two negative numbers returns zero", () => {
  assert.equal(subtract(-1, -1), 0);
});

test("subtract with floats", () => {
  assert.ok(Math.abs(subtract(1.5, 0.5) - 1.0) < Number.EPSILON * 10);
});

test("subtract with zero", () => {
  assert.equal(subtract(0, 0), 0);
  assert.equal(subtract(5, 0), 5);
  assert.equal(subtract(0, 5), -5);
});
