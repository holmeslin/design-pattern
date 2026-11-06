//! 把狀態放進型別裡:`Order<Created>`、`Order<Paid>`、`Order<Shipped>` 是三個不同的型別,
//! 每個狀態只有自己合法的方法,非法的轉換在編譯期就被擋下。
//!
//! 合法:先付款再出貨。
//!
//! ```
//! use typestate::Order;
//!
//! Order::new(1).pay().ship();
//! ```
//!
//! 非法:`Order<Created>` 沒有 `ship` 方法,沒付款不能出貨。
//! 與上面的範例只差少了 `.pay()`,證明失敗來自這一處。
//!
//! ```compile_fail
//! use typestate::Order;
//!
//! Order::new(1).ship();
//! ```
//!
//! 合法:`pay` 取走 `self` 並回傳新狀態,新的變數可以繼續用。
//!
//! ```
//! use typestate::Order;
//!
//! let o = Order::new(1);
//! let p = o.pay();
//! ```
//!
//! 非法:`pay(self)` 已經把 `o` move 走,舊狀態不能再用。
//! 與上面的範例只差最後多的 `o.pay();`。
//!
//! ```compile_fail
//! use typestate::Order;
//!
//! let o = Order::new(1);
//! let p = o.pay();
//! o.pay();
//! ```

use std::marker::PhantomData;

// 狀態只是型別層級的標記,不帶資料。
pub struct Created;
pub struct Paid;
pub struct Shipped;

pub struct Order<S> {
    id: u32,
    _state: PhantomData<S>,
}

// id 與狀態無關,所有狀態共用。
impl<S> Order<S> {
    pub fn id(&self) -> u32 {
        self.id
    }
}

impl Order<Created> {
    pub fn new(id: u32) -> Self {
        Order {
            id,
            _state: PhantomData,
        }
    }

    pub fn status(&self) -> &'static str {
        "created"
    }

    // 取走 self:舊的 Order<Created> 在呼叫後就不能再用。
    pub fn pay(self) -> Order<Paid> {
        Order {
            id: self.id,
            _state: PhantomData,
        }
    }
}

impl Order<Paid> {
    pub fn status(&self) -> &'static str {
        "paid"
    }

    pub fn ship(self) -> Order<Shipped> {
        Order {
            id: self.id,
            _state: PhantomData,
        }
    }
}

impl Order<Shipped> {
    pub fn status(&self) -> &'static str {
        "shipped"
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 依序 pay -> ship,每一步的型別都不同;id 在所有狀態都能讀
    #[test]
    fn order_walks_through_states() {
        let created = Order::new(42);
        assert_eq!(created.status(), "created");
        let paid = created.pay();
        assert_eq!(paid.status(), "paid");
        let shipped = paid.ship();
        assert_eq!(shipped.status(), "shipped");
        assert_eq!(shipped.id(), 42);
    }

    // 一氣呵成的鏈式呼叫
    #[test]
    fn chained_calls_end_shipped() {
        let order = Order::new(42).pay().ship();
        assert_eq!(order.status(), "shipped");
        assert_eq!(order.id(), 42);
    }
}
