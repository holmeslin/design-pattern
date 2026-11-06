import { test } from 'node:test';
import assert from 'node:assert/strict';
import { clone } from './prototype.js';

test('複製出來的物件是深拷貝,改複本不影響原型', () => {
  const orig = { name: 'base', tags: ['a'] };
  const copy = clone(orig);
  copy.tags.push('b');
  assert.deepEqual(orig.tags, ['a']);
  assert.deepEqual(copy.tags, ['a', 'b']);
});
