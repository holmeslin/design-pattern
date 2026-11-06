import { test } from 'node:test';
import assert from 'node:assert/strict';
import { createNotifier } from './factory_method.js';

test('依管道名稱建立對應的通知器', () => {
  assert.equal(createNotifier('email').send('hello'), 'email: hello');
  assert.equal(createNotifier('sms').send('hello'), 'sms: hello');
});

test('未知的管道丟出錯誤', () => {
  assert.throws(() => createNotifier('fax'), { message: 'unknown channel: "fax"' });
});
