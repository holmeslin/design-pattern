import { test } from 'node:test';
import assert from 'node:assert/strict';
import { config } from './singleton.js';

test('重複 import 拿到的是同一個物件', async () => {
  const again = await import('./singleton.js');
  assert.equal(config, again.config);
});

test('一處設定,另一處讀得到', async () => {
  const again = await import('./singleton.js');
  config.set('env', 'prod');
  assert.equal(again.config.get('env'), 'prod');
});
