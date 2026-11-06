import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Order } from './state.js';

test('訂單依序從 created 推進到 paid、shipped', () => {
  const order = new Order();
  assert.equal(order.advance(), 'created -> paid');
  assert.equal(order.advance(), 'paid -> shipped');
});

test('已出貨的訂單不能再推進', () => {
  const order = new Order();
  order.advance();
  order.advance();
  assert.throws(() => order.advance(), { message: 'order already shipped' });
});
