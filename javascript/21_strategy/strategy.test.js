import { test } from 'node:test';
import assert from 'node:assert/strict';
import { ExportService } from './export_service.js';
import { CSVExporter, JSONExporter } from './strategy_class.js';
import { csvExport } from './strategy_function.js';

const rows = [{ id: '1', name: 'Alice' }, { id: '2', name: 'Bob' }];
const classService = () => new ExportService(new Map([
  ['csv', new CSVExporter()],
  ['json', new JSONExporter()],
]));

test('class 版:csv 輸出表頭加每列資料,結尾不加換行', () => {
  assert.equal(classService().export(rows, 'csv'), 'id,name\n1,Alice\n2,Bob');
});

test('class 版:json 輸出縮排 2 格的 JSON', () => {
  assert.equal(classService().export(rows, 'json'), JSON.stringify(rows, null, 2));
});

test('未註冊的格式 throw,而且不呼叫任何策略', () => {
  let called = 0;
  const spy = { export: () => { called++; return ''; } };
  const service = new ExportService(new Map([['csv', spy]]));
  assert.throws(() => service.export(rows, 'xml'), { message: 'unsupported: "xml"' });
  assert.equal(called, 0);
});

test('函式版:裸函式包成 { export } 後,輸出與 class 版 csv 相同', () => {
  const service = new ExportService(new Map([['csv', { export: csvExport }]]));
  assert.equal(service.export(rows, 'csv'), classService().export(rows, 'csv'));
});

test('標準庫式:比較函式本身就是策略', () => {
  assert.deepEqual([3, 1, 2].toSorted((a, b) => a - b), [1, 2, 3]);
});
