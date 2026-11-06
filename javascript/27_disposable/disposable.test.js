import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Connection, FileLock } from './disposable.js';

test('using:離開 block 時依宣告的相反順序關閉', () => {
  const log = [];
  {
    using a = new Connection('a', log);
    using b = new Connection('b', log);
    log.push(a.query('SELECT 1'));
  }
  assert.deepEqual(log, ['a: SELECT 1', 'close b', 'close a']);
});

test('using:block 內丟例外,資源仍會關閉', () => {
  const log = [];
  assert.throws(() => {
    using c = new Connection('c', log);
    throw new Error('boom');
  }, /boom/);
  assert.ok(log.includes('close c'));
});

test('await using:離開 block 時 await 非同步釋放', async () => {
  const log = [];
  {
    await using lock = await FileLock.acquire('f', log);
  }
  assert.deepEqual(log, ['lock f', 'unlock f']);
});
