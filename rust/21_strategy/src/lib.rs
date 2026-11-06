use std::collections::{BTreeMap, HashMap};
use std::fmt;

mod strategy_closure;
mod strategy_trait;

pub use strategy_closure::csv_export;
pub use strategy_trait::{CsvExporter, JsonExporter};

// BTreeMap 讓欄位依 key 排序,輸出固定,測試不會因 HashMap 的隨機順序而不穩。
pub type Row = BTreeMap<String, String>;

// 「會變的行為」的形狀;ExportService 只認這個 trait。
pub trait Exporter {
    fn export(&self, rows: &[Row]) -> String;
}

// 函式與閉包自動成為策略,不必像 Go 的 ExportFunc 那樣手動包一層轉接器。
impl<F: Fn(&[Row]) -> String> Exporter for F {
    fn export(&self, rows: &[Row]) -> String {
        self(rows)
    }
}

#[derive(Debug, PartialEq)]
pub enum ExportError {
    Unsupported(String),
}

impl fmt::Display for ExportError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            ExportError::Unsupported(format) => write!(f, "unsupported: {format:?}"),
        }
    }
}

impl std::error::Error for ExportError {}

// Context:只查表 + 轉交,不碰任何格式細節。
#[derive(Default)]
pub struct ExportService {
    strategies: HashMap<String, Box<dyn Exporter>>,
}

impl ExportService {
    // 'static:策略被收進 Box<dyn Exporter>,不能借用呼叫端的區域變數。
    pub fn register(&mut self, format: &str, exporter: impl Exporter + 'static) {
        self.strategies
            .insert(format.to_string(), Box::new(exporter));
    }

    pub fn export(&self, rows: &[Row], format: &str) -> Result<String, ExportError> {
        let strategy = self
            .strategies
            .get(format)
            .ok_or_else(|| ExportError::Unsupported(format.to_string()))?;
        Ok(strategy.export(rows))
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::cell::Cell;
    use std::rc::Rc;

    fn row(id: &str, name: &str) -> Row {
        Row::from([
            ("id".to_string(), id.to_string()),
            ("name".to_string(), name.to_string()),
        ])
    }

    fn rows() -> Vec<Row> {
        vec![row("1", "Alice"), row("2", "Bob")]
    }

    fn service() -> ExportService {
        let mut svc = ExportService::default();
        svc.register("csv", CsvExporter);
        svc.register("json", JsonExporter);
        svc
    }

    // 寫法 A:struct 策略,依 format 查表委派,CSV 輸出固定
    #[test]
    fn trait_csv() {
        assert_eq!(
            service().export(&rows(), "csv").unwrap(),
            "id,name\n1,Alice\n2,Bob"
        );
    }

    // 寫法 A:JSON 手寫,2 格縮排
    #[test]
    fn trait_json() {
        assert_eq!(
            service().export(&rows(), "json").unwrap(),
            "[\n  {\n    \"id\": \"1\",\n    \"name\": \"Alice\"\n  },\n  {\n    \"id\": \"2\",\n    \"name\": \"Bob\"\n  }\n]"
        );
    }

    // 未註冊的格式短路回錯,而且不會呼叫任何已註冊的策略
    #[test]
    fn unsupported_format_short_circuits() {
        let calls = Rc::new(Cell::new(0u32));
        let spy = Rc::clone(&calls);
        let mut svc = ExportService::default();
        svc.register("csv", move |_: &[Row]| {
            spy.set(spy.get() + 1);
            String::new()
        });

        let err = svc.export(&rows(), "xml").unwrap_err();
        assert_eq!(err, ExportError::Unsupported("xml".to_string()));
        assert_eq!(err.to_string(), "unsupported: \"xml\"");
        assert_eq!(calls.get(), 0);

        // 對照組:查得到的格式才會呼叫策略,證明上面的 0 不是 spy 沒接上
        svc.export(&rows(), "csv").unwrap();
        assert_eq!(calls.get(), 1);
    }

    // 寫法 B:裸函式直接註冊,輸出與 struct 版相同
    #[test]
    fn function_is_a_strategy() {
        let mut svc = ExportService::default();
        svc.register("csv", csv_export);
        assert_eq!(
            svc.export(&rows(), "csv").unwrap(),
            service().export(&rows(), "csv").unwrap()
        );
    }

    // 空資料:CSV 回空字串(JSON 回 "[]")
    #[test]
    fn empty_rows() {
        assert_eq!(CsvExporter.export(&[]), "");
        assert_eq!(csv_export(&[]), "");
        assert_eq!(JsonExporter.export(&[]), "[]");
    }

    // 缺 key 的列依表頭查值,該欄輸出空字串,其餘欄位不會錯位
    #[test]
    fn missing_key_yields_empty_field() {
        let sparse = vec![
            row("1", "Alice"),
            Row::from([("name".to_string(), "Bob".to_string())]),
        ];
        assert_eq!(CsvExporter.export(&sparse), "id,name\n1,Alice\n,Bob");
        assert_eq!(csv_export(&sparse), "id,name\n1,Alice\n,Bob");
    }

    // 標準庫式:sort_by_key 的 key 閉包本身就是策略
    #[test]
    fn std_closure_is_a_strategy() {
        let mut v = vec![3, 1, 2];
        v.sort_by_key(|&x| std::cmp::Reverse(x));
        assert_eq!(v, [3, 2, 1]);
    }
}
