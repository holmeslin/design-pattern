import { test } from 'node:test';
import assert from 'node:assert/strict';
import { num, add, mul } from './interpreter.js';

test('文法樹 (1 + 2) * 3 直譯後等於 9', () => {
  assert.equal(mul(add(num(1), num(2)), num(3)).interpret(), 9);
});
