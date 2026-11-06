// 寫法 B:策略就是一個純函式;註冊時包成 { export: csvExport },物件字面值就是轉接器。
export function csvExport(rows) {
  const cols = Object.keys(rows[0]);
  // shortcut: 不處理含逗號/引號/換行的欄位跳脫,需要時改用 RFC 4180 跳脫
  return [cols.join(','), ...rows.map((row) => cols.map((c) => row[c]).join(','))].join('\n');
}
