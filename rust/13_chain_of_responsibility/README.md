# Chain of Responsibility(責任鏈)

把多個處理者**串成一條鏈**,請求沿著鏈依序傳遞,任何一環都可以拒絕;發送端不需要知道是哪一環處理的。

本範例:請求驗證鏈 Auth → RateLimit → Body,各自檢查一個面向,通過就交給下一棒。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Handler** | `trait Handler` | 每一環共同的介面:`handle(&self, &Request) -> Result<(), ChainError>` |
| **ConcreteHandler** | `AuthHandler`、`RateLimitHandler`、`BodyHandler` | 檢查自己負責的面向,失敗回 `Err`,通過就轉給下一棒 |
| **Request** | `Request` | 在鏈上流動的資料 |

## Rust 慣用寫法

每一環以 `Option<Box<dyn Handler>>` **擁有**下一環,所有權沿著鏈往下流,不需要 `Rc` 或參考;沒有下一環時直接回 `Ok(())`。`then(self, next)` 取走 `self` 再還回來,是 builder 風格的串接。它是「設定」而不是「附加」,所以多環要巢狀寫成 `a.then(b.then(c))`;`a.then(b).then(c)` 會把 `b` 蓋掉。

拒絕以 `Result<(), ChainError>` 回傳(`ChainError` 實作 `Display` 與 `Error`),不靠 panic 也不靠字串比對。`RateLimitHandler` 用 `Cell<u32>` 計數,讓 `handle` 只需要 `&self`,同一個 `Handler` trait 不必為了一個有狀態的環改成 `&mut self`。`Cell` 不是 `Sync`,要跨執行緒共用請改用 `AtomicU32`。

## 執行

```sh
cd rust && cargo test -p chain_of_responsibility
```
