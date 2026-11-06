pub trait Handler {
    fn handle(&self, req: &str) -> String;
}

pub struct Base;

impl Handler for Base {
    fn handle(&self, _req: &str) -> String {
        "ok".to_string()
    }
}

// 裝飾器本身也實作 Handler,所以可以無限層層包裝;泛型讓包裝在編譯期展開。
pub struct Logging<H: Handler>(pub H);
pub struct Auth<H: Handler>(pub H);

impl<H: Handler> Handler for Logging<H> {
    fn handle(&self, req: &str) -> String {
        format!("[log] {}", self.0.handle(req))
    }
}

impl<H: Handler> Handler for Auth<H> {
    fn handle(&self, req: &str) -> String {
        format!("[auth] {}", self.0.handle(req))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 裝飾器可以任意層層包裝,外層先處理
    #[test]
    fn decorators_stack_outermost_first() {
        assert_eq!(Logging(Auth(Base)).handle("req"), "[log] [auth] ok");
    }
}
