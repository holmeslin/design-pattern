import { test } from 'node:test';
import assert from 'node:assert/strict';
import { StripeAdapter, StripeSDK } from './adapter.js';

test('Adapter 把「分」轉成 SDK 要的「元」,呼叫端不用知道 SDK 的單位', () => {
  assert.equal(new StripeAdapter(new StripeSDK()).pay(1234), 'charged $12.34');
});
