import { test } from 'node:test';
import assert from 'node:assert/strict';
import { file, folder } from './composite.js';

test('資料夾遞迴加總子節點大小,葉節點與容器用法一致', () => {
  const sub = folder();
  sub.add(file(200));
  const root = folder();
  root.add(file(100));
  root.add(sub);
  assert.equal(root.size(), 300);
});
