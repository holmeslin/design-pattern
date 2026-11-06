import { test } from 'node:test';
import assert from 'node:assert/strict';
import { QueryBuilder } from './builder.js';

test('鏈式呼叫逐步組出完整的 SQL', () => {
  const sql = new QueryBuilder().select('id', 'name').from('users').where('age > 18').build();
  assert.equal(sql, 'SELECT id, name FROM users WHERE age > 18');
});

test('不呼叫 where 時,SQL 不帶 WHERE 子句', () => {
  const sql = new QueryBuilder().select('id').from('users').build();
  assert.equal(sql, 'SELECT id FROM users');
});
