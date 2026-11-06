use std::cell::OnceCell;

// Proxy 與真實物件實作同一個介面,呼叫端分不出差別。
pub trait Image {
    fn load(&self) -> String;
}

pub struct RealImage {
    filename: String,
}

impl Image for RealImage {
    fn load(&self) -> String {
        format!("loading {}", self.filename)
    }
}

// virtual proxy:真正的 RealImage 直到第一次 load 才建立。
pub struct ProxyImage {
    filename: String,
    real: OnceCell<RealImage>,
}

impl ProxyImage {
    pub fn new(filename: &str) -> Self {
        Self {
            filename: filename.to_string(),
            real: OnceCell::new(),
        }
    }
}

impl Image for ProxyImage {
    // OnceCell 透過 interior mutability 在 &self 下延遲初始化,介面不必改成 &mut self。
    fn load(&self) -> String {
        match self.real.get() {
            Some(_) => format!("cached {}", self.filename),
            None => self
                .real
                .get_or_init(|| RealImage {
                    filename: self.filename.clone(),
                })
                .load(),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 第一次才真的載入,之後走快取;load 只需要 &self
    #[test]
    fn proxy_loads_lazily_then_caches() {
        let image = ProxyImage::new("cat.png");
        assert_eq!(image.load(), "loading cat.png");
        assert_eq!(image.load(), "cached cat.png");
    }
}
