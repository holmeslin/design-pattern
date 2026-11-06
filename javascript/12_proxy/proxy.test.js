import { test } from 'node:test';
import assert from 'node:assert/strict';
import { RealImage, withCache } from './proxy.js';

test('代理第一次真的載入,第二次直接回傳快取', () => {
  const image = withCache(new RealImage('cat.png'));
  assert.equal(image.load(), 'loading cat.png');
  assert.equal(image.load(), 'cached cat.png');
});
