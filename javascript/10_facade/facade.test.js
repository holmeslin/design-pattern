import { test } from 'node:test';
import assert from 'node:assert/strict';
import { HomeTheater } from './facade.js';

test('一個 watchMovie() 依序驅動三個子系統,呼叫端不用知道細節', () => {
  assert.deepEqual(new HomeTheater().watchMovie(), [
    'lights dim',
    'projector on',
    'amplifier on',
  ]);
});
