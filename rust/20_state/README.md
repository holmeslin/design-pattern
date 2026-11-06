# State(狀態)

物件的**行為隨內部狀態改變**,每個狀態封裝成獨立的型別,由狀態自己決定下一個狀態。

本範例:訂單狀態機,`created -> paid -> shipped`;已出貨的訂單再推進會回傳錯誤。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Context** | `Order` | 持有目前狀態,`advance` 委派給它 |
| **State** | `trait State` | 宣告 `name` 與 `next` |
| **ConcreteState** | `Created`、`Paid`、`Shipped`(私有) | 各自決定下一個狀態;`Shipped` 回傳 `OrderError::AlreadyShipped` |

## Rust 慣用寫法

這是**執行期**狀態機:`Order` 持有 `Box<dyn State>`(trait object),非法轉換(例如已出貨再推進)要到執行期才以 `Err` 回報。由編譯器擋住非法轉換的版本見 `25_typestate`。

`Box<dyn State>` 無法 derive `Default`,所以 `impl Default for Order` 是手寫的,內容只是呼叫 `new()`(clippy 的 `new_without_default` 也會要求這樣)。

`advance` 先呼叫 `next()`、成功後才賦值,所以失敗時 `Order` 的狀態維持原樣。

## 執行

```sh
cd rust && cargo test -p state
```
