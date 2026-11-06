use std::collections::HashMap;
use std::sync::{LazyLock, Mutex};

// 欄位私有且沒有公開建構子,外部無法自己 new 出第二個 Config。
pub struct Config {
    values: Mutex<HashMap<String, String>>,
}

impl Config {
    // Rust 不允許未同步的可變全域變數,要能寫入就得包 Mutex(編譯器強制)。
    // 中毒時取回內部資料繼續用:這裡的操作不會讓 map 處於不一致狀態。
    fn lock(&self) -> std::sync::MutexGuard<'_, HashMap<String, String>> {
        self.values.lock().unwrap_or_else(|e| e.into_inner())
    }

    pub fn set(&self, key: &str, value: &str) {
        self.lock().insert(key.to_string(), value.to_string());
    }

    pub fn get(&self, key: &str) -> Option<String> {
        self.lock().get(key).cloned()
    }
}

// LazyLock 第一次存取時才初始化,且保證只初始化一次(對應 Go 的 sync.Once)。
static INSTANCE: LazyLock<Config> = LazyLock::new(|| Config {
    values: Mutex::new(HashMap::new()),
});

pub fn instance() -> &'static Config {
    &INSTANCE
}

#[cfg(test)]
mod tests {
    use super::*;

    // 每次取得的都是同一個實例(位址相同)
    #[test]
    fn instance_is_unique() {
        assert!(std::ptr::eq(instance(), instance()));
    }

    // 寫入後能讀回;沒設過的 key 回 None。
    // 測試平行執行,所以這裡的 key 與其他測試不重疊。
    #[test]
    fn set_then_get() {
        instance().set("env", "prod");
        assert_eq!(instance().get("env"), Some("prod".to_string()));
        assert_eq!(instance().get("missing"), None);
    }
}
