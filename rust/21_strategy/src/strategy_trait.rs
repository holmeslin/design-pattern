// 寫法 A:一種演算法一個型別,適合策略有狀態或需要多個方法時。
use crate::{Exporter, Row};

pub struct CsvExporter;
pub struct JsonExporter;

impl Exporter for CsvExporter {
    fn export(&self, rows: &[Row]) -> String {
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
}

impl Exporter for JsonExporter {
    fn export(&self, rows: &[Row]) -> String {
        if rows.is_empty() {
            return "[]".to_string();
        }
        let objects = rows.iter().map(|row| {
            // shortcut: 不處理字串跳脫,需要時改用 serde_json
            let fields = row
                .iter()
                .map(|(k, v)| format!("    \"{k}\": \"{v}\""))
                .collect::<Vec<_>>()
                .join(",\n");
            format!("  {{\n{fields}\n  }}")
        });
        format!("[\n{}\n]", objects.collect::<Vec<_>>().join(",\n"))
    }
}
