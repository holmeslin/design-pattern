import { test } from 'node:test';
import assert from 'node:assert/strict';
import { ChatRoom, User } from './mediator.js';

test('透過聊天室廣播,其他人收到、發送者不會收到自己的訊息', () => {
  const room = new ChatRoom();
  const alice = new User('Alice', room);
  const bob = new User('Bob', room);
  const carol = new User('Carol', room);
  alice.send('hi');
  assert.deepEqual(bob.inbox, ['Bob received from Alice: hi']);
  assert.deepEqual(carol.inbox, ['Carol received from Alice: hi']);
  assert.deepEqual(alice.inbox, []);
});
