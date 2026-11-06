# Factory Method(工廠方法)

把「建立哪一種物件」的決定**集中到一個工廠**,呼叫端只說要什麼,不用 `new` 具體型別。

本範例:通知器工廠,依管道名稱(`email` / `sms`)建立對應的通知器,未知管道直接報錯。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Creator** | `createNotifier` | 依參數決定建立哪一種產品 |
| **Product** | 有 `send(msg)` 的物件 | 產品共同的形狀 |
| **ConcreteProduct** | `creators` 表裡的 `email`、`sms` | 各管道的具體實作 |

JS 的工廠就是一個函式加一張物件查表,不需要 Creator 類別階層;新增管道只要在表裡多加一個 key。

## 執行

```sh
cd javascript && node --test 03_factory_method/*.test.js
```
