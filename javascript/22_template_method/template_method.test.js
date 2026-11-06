import { test } from 'node:test';
import assert from 'node:assert/strict';
import { DataJob, EvenFilterJob } from './template_method.js';

test('子類別只填三個步驟,固定流程由 run 負責', () => {
  assert.equal(new EvenFilterJob().run(), '[2 4]');
});

test('沒覆寫步驟的基底類別執行會 throw', () => {
  assert.throws(() => new DataJob().run(), { message: 'not implemented' });
});
