import { test } from 'node:test';
import assert from 'node:assert/strict';
import { slowApi, withTimeout } from './cancellation.js';

test('在時限內完成,回傳結果', async () => {
  assert.equal(await withTimeout((s) => slowApi(10, { signal: s }), 200), 'done');
});

test('逾時:reject TimeoutError', async () => {
  await assert.rejects(
    withTimeout((s) => slowApi(200, { signal: s }), 10),
    { name: 'TimeoutError' },
  );
});

test('使用者主動取消:reject AbortError,與逾時可區分', async () => {
  const controller = new AbortController();
  setTimeout(() => controller.abort(), 10);

  await assert.rejects(
    withTimeout((s) => slowApi(200, { signal: s }), 1000, { signal: controller.signal }),
    { name: 'AbortError' },
  );
});

test('傳入已取消的 signal:立刻 reject AbortError', async () => {
  const signal = AbortSignal.abort();

  await assert.rejects(slowApi(200, { signal }), { name: 'AbortError' });
});
