import { test } from 'node:test';
import assert from 'node:assert/strict';
import { AuthHandler, RateLimitHandler, BodyHandler } from './chain_of_responsibility.js';

function buildChain() {
  const auth = new AuthHandler();
  auth.setNext(new RateLimitHandler()).setNext(new BodyHandler());
  return auth;
}

test('合法請求通過整條鏈,回傳 ok', () => {
  assert.equal(buildChain().handle({ user: 'alice', body: 'hello' }), 'ok');
});

test('沒有使用者,在 AuthHandler 就被擋下', () => {
  assert.throws(() => buildChain().handle({ user: '', body: 'hello' }), { message: 'auth failed' });
});

test('body 為空,在 BodyHandler 被擋下', () => {
  assert.throws(() => buildChain().handle({ user: 'alice', body: '' }), { message: 'empty body' });
});

test('同一條鏈的合法請求第 4 次,被 RateLimitHandler 擋下', () => {
  const chain = buildChain();
  const req = { user: 'alice', body: 'hello' };
  for (let i = 0; i < 3; i++) {
    assert.equal(chain.handle(req), 'ok');
  }
  assert.throws(() => chain.handle(req), { message: 'rate limited' });
});
