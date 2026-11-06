// 文法的每條規則就是一個 variant;遞迴結構靠 Box 提供固定大小。
pub enum Expr {
    Num(i64),
    Add(Box<Expr>, Box<Expr>),
    Mul(Box<Expr>, Box<Expr>),
}

impl Expr {
    // shortcut: wrapping 運算,溢位時回繞而不是 panic;輸入不可信時改用 checked_* 並回傳 Result。
    pub fn interpret(&self) -> i64 {
        match self {
            Expr::Num(n) => *n,
            Expr::Add(l, r) => l.interpret().wrapping_add(r.interpret()),
            Expr::Mul(l, r) => l.interpret().wrapping_mul(r.interpret()),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::Expr::{Add, Mul, Num};

    // (1 + 2) * 3:遞迴求值要先算括號內的加法
    #[test]
    fn evaluates_nested_expression() {
        let expr = Mul(
            Box::new(Add(Box::new(Num(1)), Box::new(Num(2)))),
            Box::new(Num(3)),
        );
        assert_eq!(expr.interpret(), 9);
    }
}
