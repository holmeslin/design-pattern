import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Ring } from './iterator.js';

test('展開運算子依序取出所有元素', () => {
  assert.deepEqual([...new Ring(10, 20, 30)], [10, 20, 30]);
});

test('Iterator helpers 直接可用:map 後 toArray', () => {
  const doubled = new Ring(1, 2, 3)[Symbol.iterator]().map((x) => x * 2).toArray();
  assert.deepEqual(doubled, [2, 4, 6]);
});
