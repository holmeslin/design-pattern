import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Station, display } from './observer.js';

test('溫度改變時,依訂閱順序通知所有觀察者', () => {
  const station = new Station();
  const out = [];
  station.addEventListener('temp', display('phone', out));
  station.addEventListener('temp', display('window', out));
  station.setTemp(25);
  assert.deepEqual(out, ['phone: 25', 'window: 25']);
});

test('取消訂閱後不再收到通知', () => {
  const station = new Station();
  const out = [];
  const phone = display('phone', out);
  station.addEventListener('temp', phone);
  station.addEventListener('temp', display('window', out));
  station.setTemp(25);
  station.removeEventListener('temp', phone);
  station.setTemp(30);
  assert.deepEqual(out.slice(2), ['window: 30']);
});
