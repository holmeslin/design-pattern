# Chain of Responsibility(責任鏈)

把請求沿著一條**處理者鏈**往下傳,每個處理者決定自己處理、擋下,或交給下一棒。

本範例:請求驗證鏈,依序檢查登入(Auth)、頻率限制(RateLimit)、內容(Body)。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Handler** | `Handler` | 持有下一棒,預設把請求轉給下一棒;沒有下一棒就回傳 `'ok'` |
| **ConcreteHandler** | `AuthHandler`、`RateLimitHandler`、`BodyHandler` | 檢查自己負責的條件,不通過就 throw,通過就 `super.handle(req)` 交棒 |
| **Client** | 測試中的 `auth.setNext(rateLimit).setNext(body)` | 組裝鏈並把請求丟給第一棒 |

`setNext` 回傳傳入的 handler,所以可以連續串接;`RateLimitHandler` 的計數屬於該實例,所以同一條鏈才會累計。

## 執行

```sh
cd javascript && node --test 13_chain_of_responsibility/*.test.js
```
