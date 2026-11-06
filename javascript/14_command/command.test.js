import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Account, History, deposit } from './command.js';

test('執行存款指令後餘額增加,undo 後還原', () => {
  const acc = new Account();
  const history = new History();
  history.run(deposit(acc, 100));
  assert.equal(acc.balance, 100);
  history.undoLast();
  assert.equal(acc.balance, 0);
});

test('空歷史呼叫 undoLast 不動作', () => {
  const acc = new Account();
  const history = new History();
  assert.doesNotThrow(() => history.undoLast());
  assert.equal(acc.balance, 0);
});
