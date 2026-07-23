# State(狀態)

讓物件在內部狀態改變時,自動變更其行為,看起來就像換了一個類別。

## 角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **Context** | `Order` | 持有目前狀態,對外提供 `Advance()`,自己不判斷狀態邏輯,只委派給 `State` |
| **State**(介面) | `State` | 定義 `Name()`、`Next(*Order) State` 兩個行為 |
| **ConcreteState** | `createdState`、`paidState`、`shippedState` | 各自決定「我是誰」以及「下一步轉去哪個狀態」 |

## 執行

```sh
go test ./...
```
