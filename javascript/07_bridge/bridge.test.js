import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Radio, RemoteControl, TV } from './bridge.js';

test('同一支遙控器換 TV 實作,行為不變', () => {
  const remote = new RemoteControl(new TV());
  remote.volumeUp();
  assert.equal(remote.status(), 'TV volume: 10');
});

test('同一支遙控器換 Radio 實作,行為不變', () => {
  const remote = new RemoteControl(new Radio());
  remote.volumeUp();
  assert.equal(remote.status(), 'Radio volume: 10');
});
