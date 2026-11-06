use std::collections::HashMap;
use std::rc::Rc;

// 內部狀態(可共享):同樣的字型與大小只需要存一份。
#[derive(Debug, PartialEq, Eq, Hash)]
pub struct CharFormat {
    pub font: String,
    pub size: u32,
}

// 工廠負責快取:相同 key 回傳同一個 Rc,而不是新建。
#[derive(Default)]
pub struct FormatFactory {
    cache: HashMap<(String, u32), Rc<CharFormat>>,
}

impl FormatFactory {
    pub fn get(&mut self, font: &str, size: u32) -> Rc<CharFormat> {
        let format = self
            .cache
            .entry((font.to_string(), size))
            .or_insert_with(|| {
                Rc::new(CharFormat {
                    font: font.to_string(),
                    size,
                })
            });
        Rc::clone(format)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 相同的 (font, size) 共用同一份物件,不同的才另外建立
    #[test]
    fn same_key_shares_one_instance() {
        let mut factory = FormatFactory::default();
        let a = factory.get("Arial", 12);
        let b = factory.get("Arial", 12);
        let c = factory.get("Arial", 14);

        assert!(Rc::ptr_eq(&a, &b));
        assert!(!Rc::ptr_eq(&a, &c));
        // cache、a、b 各持有一份
        assert_eq!(Rc::strong_count(&a), 3);
    }
}
