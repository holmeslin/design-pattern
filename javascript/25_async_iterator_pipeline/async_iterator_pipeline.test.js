import { test } from 'node:test';
import assert from 'node:assert/strict';
import { filter, map, take } from './async_iterator_pipeline.js';

test('日誌管線:只留 ERROR、去掉前綴、取前 2 筆,並一路關閉上游', async () => {
  let closed = false;
  async function* source() {
    try {
      yield* ['INFO a', 'ERROR b', 'INFO c', 'ERROR d', 'ERROR e'];
    } finally {
      closed = true;
    }
  }

  const result = await Array.fromAsync(
    take(map(filter(source(), (l) => l.startsWith('ERROR')), (l) => l.slice(6)), 2),
  );

  assert.deepEqual(result, ['b', 'd']);
  assert.equal(closed, true);
});

test('filter 會 await 非同步 predicate,只留符合的項目', async () => {
  const result = await Array.fromAsync(filter([1, 2, 3, 4], async (n) => n % 2 === 0));

  assert.deepEqual(result, [2, 4]);
});
