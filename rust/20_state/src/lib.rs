use std::fmt;

#[derive(Debug, PartialEq)]
pub enum OrderError {
    AlreadyShipped,
}

impl fmt::Display for OrderError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            OrderError::AlreadyShipped => write!(f, "order already shipped"),
        }
    }
}

impl std::error::Error for OrderError {}

// 每個狀態自己決定下一個狀態是誰;Order 不寫任何 if/match 分支。
pub trait State {
    fn name(&self) -> &'static str;
    fn next(&self) -> Result<Box<dyn State>, OrderError>;
}

struct Created;
struct Paid;
struct Shipped;

impl State for Created {
    fn name(&self) -> &'static str {
        "created"
    }
    fn next(&self) -> Result<Box<dyn State>, OrderError> {
        Ok(Box::new(Paid))
    }
}

impl State for Paid {
    fn name(&self) -> &'static str {
        "paid"
    }
    fn next(&self) -> Result<Box<dyn State>, OrderError> {
        Ok(Box::new(Shipped))
    }
}

impl State for Shipped {
    fn name(&self) -> &'static str {
        "shipped"
    }
    fn next(&self) -> Result<Box<dyn State>, OrderError> {
        Err(OrderError::AlreadyShipped)
    }
}

pub struct Order {
    state: Box<dyn State>,
}

impl Order {
    pub fn new() -> Self {
        Order {
            state: Box::new(Created),
        }
    }

    // `?` 在賦值之前就返回,所以推進失敗時狀態不變。
    pub fn advance(&mut self) -> Result<String, OrderError> {
        let next = self.state.next()?;
        let transition = format!("{} -> {}", self.state.name(), next.name());
        self.state = next;
        Ok(transition)
    }

    pub fn state(&self) -> &'static str {
        self.state.name()
    }
}

// Box<dyn State> 無法 derive Default,所以手寫並委派給 new()。
impl Default for Order {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 依序走完 created -> paid -> shipped
    #[test]
    fn advances_through_states() {
        let mut order = Order::new();
        assert_eq!(order.advance(), Ok("created -> paid".to_string()));
        assert_eq!(order.advance(), Ok("paid -> shipped".to_string()));
    }

    // 已出貨再推進會失敗,而且狀態不變
    #[test]
    fn shipped_order_cannot_advance() {
        let mut order = Order::default();
        order.advance().unwrap();
        order.advance().unwrap();
        let err = order.advance().unwrap_err();
        assert_eq!(err, OrderError::AlreadyShipped);
        assert_eq!(err.to_string(), "order already shipped");
        assert_eq!(order.state(), "shipped");
    }
}
