import { test } from 'node:test';
import assert from 'node:assert/strict';
import { CharFormatFactory } from './flyweight.js';

test('相同 font + size 回傳同一個共享物件', () => {
  const factory = new CharFormatFactory();
  const a = factory.get('Arial', 12);
  const b = factory.get('Arial', 12);
  assert.equal(a, b);
});

test('不同 size 回傳不同的物件', () => {
  const factory = new CharFormatFactory();
  const a = factory.get('Arial', 12);
  assert.notEqual(factory.get('Arial', 14), a);
});

test('共享物件必須不可變,否則一處修改會影響所有使用者', () => {
  const factory = new CharFormatFactory();
  assert.equal(Object.isFrozen(factory.get('Arial', 12)), true);
});
