# Factory Method(工廠方法)

把「建立哪個具體物件」的決定集中到一個工廠,呼叫端只依賴共同介面,不需要知道建構細節。

本範例:依通道名稱("email" / "sms")建立對應的通知器,未知通道回傳錯誤。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Product** | `trait Notifier` | 所有具體產品共同的行為介面 |
| **ConcreteProduct** | `EmailNotifier`、`SmsNotifier` | 實作 `Notifier` 的具體型別,不對外公開 |
| **Creator** | `new_notifier` | 依參數決定並建立對應的具體產品 |

## Rust 慣用寫法

Rust 沒有類別繼承,所以工廠方法通常就是一個以 `match` 分派的函式,回傳 `Box<dyn Notifier>`。輸入可能不合法,因此回傳 `Result<_, UnknownChannel>`,而不是 panic;`UnknownChannel` 實作 `Display` 與 `Error`,可以直接用 `?` 往上傳遞。

## 執行

```sh
cd rust && cargo test -p factory_method
```
