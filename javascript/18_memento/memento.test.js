import { test } from 'node:test';
import assert from 'node:assert/strict';
import { Editor } from './memento.js';

test('save 之後繼續編輯,restore 能回到存檔時的內容', () => {
  const editor = new Editor();
  editor.type('v1');
  const saved = editor.save();
  editor.type('v2');
  assert.equal(editor.content, 'v2');
  editor.restore(saved);
  assert.equal(editor.content, 'v1');
});

test('memento 是凍結的,外部無法竄改', () => {
  const editor = new Editor();
  editor.type('v1');
  assert.equal(Object.isFrozen(editor.save()), true);
});
