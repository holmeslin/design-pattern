# State(狀態)

讓物件在內部狀態改變時,自動改變行為,看起來就像換了一個類別。

本範例:訂單狀態流轉 created -> paid -> shipped,已出貨就不能再推進。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Context** | `Order` | 持有目前狀態,`advance()` 只委派給 state,自己不判斷狀態邏輯 |
| **State** | 有 `name` 與 `next()` 的物件這個形狀 | 宣告「我是誰」與「下一步轉去哪」 |
| **ConcreteState** | `created`、`paid`、`shipped` | 各自回傳下一個 state;`shipped.next()` 直接 throw |

JS 沒有 interface,State 只是一個約定的形狀(duck typing);三個 state 沒有各自的資料,所以用物件字面量而不是 class。

## 執行

```sh
cd javascript && node --test 20_state/*.test.js
```
