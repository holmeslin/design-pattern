import { test } from 'node:test';
import assert from 'node:assert/strict';
import { darkFactory, lightFactory, renderUI } from './abstract_factory.js';

test('換工廠就換整族產品,renderUI 不用改', () => {
  assert.equal(renderUI(darkFactory), '[dark button] [dark checkbox]');
  assert.equal(renderUI(lightFactory), '[light button] [light checkbox]');
});
