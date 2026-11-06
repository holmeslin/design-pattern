import { test } from 'node:test';
import assert from 'node:assert/strict';
import { signal, computed, effect } from './signals.js';

test('購物車:相依自動追蹤,值沒變就不通知', () => {
  const price = signal(100);
  const qty = signal(1);
  const total = computed(() => price.get() * qty.get());
  const log = [];
  effect(() => log.push(total.get()));

  qty.set(2);
  price.set(50);
  assert.deepEqual(log, [100, 200, 100]);

  qty.set(2);
  assert.equal(log.length, 3);
});
