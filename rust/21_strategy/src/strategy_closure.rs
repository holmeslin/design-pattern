// 寫法 B:策略只是一個函式時,不必造型別;lib.rs 的 blanket impl 讓它直接成為 Exporter。
use crate::Row;

pub fn csv_export(rows: &[Row]) -> String {
    let Some(first) = rows.first() else {
        return String::new();
    };
    let header_keys: Vec<&str> = first.keys().map(String::as_str).collect();
    let header = header_keys.join(",");
    // shortcut: 欄位以第一列的 key 為準,不處理含逗號/引號/換行的跳脫,需要時改用 RFC 4180 跳脫
    let lines = rows.iter().map(|row| {
        header_keys
            .iter()
            .map(|k| row.get(*k).map_or("", String::as_str))
            .collect::<Vec<_>>()
            .join(",")
    });
    std::iter::once(header)
        .chain(lines)
        .collect::<Vec<_>>()
        .join("\n")
}
