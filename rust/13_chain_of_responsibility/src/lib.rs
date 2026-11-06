use std::cell::Cell;
use std::error::Error;
use std::fmt;

pub struct Request {
    pub user: String,
    pub body: String,
}

#[derive(Debug, PartialEq)]
pub enum ChainError {
    AuthFailed,
    RateLimited,
    EmptyBody,
}

impl fmt::Display for ChainError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.write_str(match self {
            ChainError::AuthFailed => "auth failed",
            ChainError::RateLimited => "rate limited",
            ChainError::EmptyBody => "empty body",
        })
    }
}

impl Error for ChainError {}

// 失敗用 Err 往外傳;通過就交給下一棒,沒有下一棒代表整條鏈都通過。
pub trait Handler {
    fn handle(&self, req: &Request) -> Result<(), ChainError>;
}

type Next = Option<Box<dyn Handler>>;

fn forward(next: &Next, req: &Request) -> Result<(), ChainError> {
    next.as_deref().map_or(Ok(()), |h| h.handle(req))
}

// 每一棒擁有下一棒(Box<dyn Handler>),所有權沿著鏈往下,不需要 Rc 或參考。
#[derive(Default)]
pub struct AuthHandler {
    next: Next,
}

impl AuthHandler {
    pub fn new() -> Self {
        Self::default()
    }

    /// 取代(不是附加)下一個處理者;`a.then(b).then(c)` 會丟掉 b,要串三個請用巢狀 `a.then(b.then(c))`。
    pub fn then(mut self, next: impl Handler + 'static) -> Self {
        self.next = Some(Box::new(next));
        self
    }
}

impl Handler for AuthHandler {
    fn handle(&self, req: &Request) -> Result<(), ChainError> {
        if req.user.is_empty() {
            return Err(ChainError::AuthFailed);
        }
        forward(&self.next, req)
    }
}

pub struct RateLimitHandler {
    limit: u32,
    count: Cell<u32>,
    next: Next,
}

impl RateLimitHandler {
    pub fn new(limit: u32) -> Self {
        Self {
            limit,
            count: Cell::new(0),
            next: None,
        }
    }

    /// 取代(不是附加)下一個處理者;`a.then(b).then(c)` 會丟掉 b,要串三個請用巢狀 `a.then(b.then(c))`。
    pub fn then(mut self, next: impl Handler + 'static) -> Self {
        self.next = Some(Box::new(next));
        self
    }
}

impl Handler for RateLimitHandler {
    // Cell 讓計數在 &self 下可變,trait 不必改成 &mut self;saturating_add 避免溢位 panic。
    fn handle(&self, req: &Request) -> Result<(), ChainError> {
        let seen = self.count.get().saturating_add(1);
        self.count.set(seen);
        if seen > self.limit {
            return Err(ChainError::RateLimited);
        }
        forward(&self.next, req)
    }
}

#[derive(Default)]
pub struct BodyHandler {
    next: Next,
}

impl BodyHandler {
    pub fn new() -> Self {
        Self::default()
    }

    /// 取代(不是附加)下一個處理者;`a.then(b).then(c)` 會丟掉 b,要串三個請用巢狀 `a.then(b.then(c))`。
    pub fn then(mut self, next: impl Handler + 'static) -> Self {
        self.next = Some(Box::new(next));
        self
    }
}

impl Handler for BodyHandler {
    fn handle(&self, req: &Request) -> Result<(), ChainError> {
        if req.body.is_empty() {
            return Err(ChainError::EmptyBody);
        }
        forward(&self.next, req)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn chain() -> AuthHandler {
        AuthHandler::new().then(RateLimitHandler::new(3).then(BodyHandler::new()))
    }

    fn req(user: &str, body: &str) -> Request {
        Request {
            user: user.to_string(),
            body: body.to_string(),
        }
    }

    // 合法請求一路通過整條鏈
    #[test]
    fn valid_request_passes_whole_chain() {
        assert_eq!(chain().handle(&req("alice", "hello")), Ok(()));
    }

    // 第一環就拒絕,後面的環不會被執行
    #[test]
    fn auth_rejects_empty_user() {
        let err = chain().handle(&req("", "hello")).unwrap_err();
        assert_eq!(err, ChainError::AuthFailed);
        assert_eq!(err.to_string(), "auth failed");
    }

    // 通過前兩環後,由最後一環擋下空內容
    #[test]
    fn body_rejects_empty_body() {
        assert_eq!(
            chain().handle(&req("alice", "")),
            Err(ChainError::EmptyBody)
        );
    }

    // 限流狀態存在鏈上的實例裡:同一條鏈第 4 次請求被擋
    #[test]
    fn rate_limit_trips_on_fourth_request() {
        let chain = chain();
        for _ in 0..3 {
            assert_eq!(chain.handle(&req("alice", "hello")), Ok(()));
        }
        assert_eq!(
            chain.handle(&req("alice", "hello")),
            Err(ChainError::RateLimited)
        );
    }
}
