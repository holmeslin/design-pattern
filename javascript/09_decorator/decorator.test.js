import { test } from 'node:test';
import assert from 'node:assert/strict';
import { auth, base, logging } from './decorator.js';

test('裝飾者層層包住 handler,由外而內疊加行為', () => {
  assert.equal(logging(auth(base))('req'), '[log] [auth] ok');
});
