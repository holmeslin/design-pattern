// 寫法 A:一種演算法一個 class + 一個 export 方法;策略有狀態或多方法時用。
export class CSVExporter {
  export(rows) {
    if (rows.length === 0) return '';
    const cols = Object.keys(rows[0]);
    // shortcut: 不處理含逗號/引號/換行的欄位跳脫,需要時改用 RFC 4180 跳脫
    return [cols.join(','), ...rows.map((row) => cols.map((c) => row[c]).join(','))].join('\n');
  }
}

export class JSONExporter {
  export(rows) {
    return JSON.stringify(rows, null, 2);
  }
}
