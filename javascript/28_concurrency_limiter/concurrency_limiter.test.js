import { test } from 'node:test';
import assert from 'node:assert/strict';
import { createLimiter } from './concurrency_limiter.js';

test('對照 Go worker pool:5 個任務由 3 個名額輪流處理,平方和為 55', async () => {
  const limit = createLimiter(3);
  const results = await Promise.all([1, 2, 3, 4, 5].map((n) => limit(async () => n * n)));

  assert.equal(results.reduce((a, b) => a + b, 0), 55);
});

test('同時執行的任務數不超過上限,結果順序與輸入一致', async () => {
  const limit = createLimiter(3);
  let active = 0;
  let peak = 0;

  const results = await Promise.all(
    [1, 2, 3, 4, 5, 6].map((n) =>
      limit(async () => {
        active++;
        peak = Math.max(peak, active);
        await new Promise((resolve) => setTimeout(resolve, 10));
        active--;
        return n;
      }),
    ),
  );

  assert.equal(peak, 3);
  assert.deepEqual(results, [1, 2, 3, 4, 5, 6]);
});

test('任務 reject 也會釋放名額,後面的任務照常執行', async () => {
  const limit = createLimiter(1);
  const first = limit(async () => {
    throw new Error('x');
  });
  const second = limit(async () => 'ok');

  await assert.rejects(first, /x/);
  assert.equal(await second, 'ok');
});

test('max 不是正整數就 throw RangeError,不會讓任務永遠卡住', () => {
  assert.throws(() => createLimiter(0), RangeError);
  assert.throws(() => createLimiter(), RangeError);
  assert.throws(() => createLimiter(NaN), RangeError);
});
