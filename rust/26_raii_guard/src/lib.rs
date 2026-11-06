use std::cell::RefCell;

// 資源綁在值的生命週期上:持有 guard 就代表連線開著。
// 用 RefCell 共用 log,讓 Drop 在只拿到 &self 的情況下也能寫入。
pub struct ConnectionGuard<'a> {
    name: String,
    log: &'a RefCell<Vec<String>>,
}

pub fn open<'a>(name: &str, log: &'a RefCell<Vec<String>>) -> ConnectionGuard<'a> {
    log.borrow_mut().push(format!("open {name}"));
    ConnectionGuard {
        name: name.to_string(),
        log,
    }
}

impl ConnectionGuard<'_> {
    pub fn query(&self, sql: &str) -> String {
        format!("{}: {sql}", self.name)
    }
}

// 離開 scope、提早 drop、panic 展開都會走到這裡,不必手動呼叫 close。
impl Drop for ConnectionGuard<'_> {
    fn drop(&mut self) {
        self.log.borrow_mut().push(format!("close {}", self.name));
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::cell::RefCell;
    use std::panic::{self, AssertUnwindSafe};

    // 離開 scope 時依宣告的相反順序釋放:b 先關,a 後關
    #[test]
    fn drops_in_reverse_order() {
        let log = RefCell::new(Vec::new());
        {
            let a = open("a", &log);
            let _b = open("b", &log);
            log.borrow_mut().push(a.query("SELECT 1"));
        }
        assert_eq!(
            *log.borrow(),
            ["open a", "open b", "a: SELECT 1", "close b", "close a"]
        );
    }

    // panic 展開(unwind)時一樣會呼叫 Drop
    #[test]
    fn drops_on_panic() {
        let log = RefCell::new(Vec::new());
        // 暫時靜音 panic 訊息,保持測試輸出乾淨;結束前還原原本的 hook。
        let prev = panic::take_hook();
        panic::set_hook(Box::new(|_| {}));
        let result = panic::catch_unwind(AssertUnwindSafe(|| {
            let _c = open("c", &log);
            panic!("boom");
        }));
        panic::set_hook(prev);

        assert!(result.is_err());
        assert!(log.borrow().contains(&"close c".to_string()));
    }

    // 用 drop() 提早釋放:不必等到 scope 結束
    #[test]
    fn early_drop() {
        let log = RefCell::new(Vec::new());
        let d = open("d", &log);
        drop(d);
        assert_eq!(log.borrow().last().map(String::as_str), Some("close d"));
    }
}
