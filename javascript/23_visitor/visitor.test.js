import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Circle, Rectangle, areaVisitor } from './visitor.js';

test('同一個 areaVisitor 對不同圖形算出各自的面積', () => {
  assert.equal(new Rectangle(2, 3).accept(areaVisitor), 6);
  assert.equal(new Circle(1).accept(areaVisitor), Math.PI);
});
