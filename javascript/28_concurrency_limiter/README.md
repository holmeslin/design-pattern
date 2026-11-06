# Concurrency Limiter(並行數限制)

任務可以一次全部丟進來,但同時執行的數量有上限,超過的先排隊,有名額空出來再依序放行。

本範例:`createLimiter(max)` 回傳一個 `limit(task)` 函式,用它包住每個非同步任務。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Limiter** | `createLimiter(max)` 回傳的 `limit` | 接收任務、決定立刻執行或排隊,結果原樣回傳給呼叫端 |
| **Slot** | 閉包內的 `active` 計數 | 目前佔用的名額,不超過 `max` |
| **Waiting queue** | 閉包內的 `waiting` 陣列 | 滿額時以 `Promise.withResolvers()` 建立等待者,名額空出時依序喚醒 |
| **Task** | `() => Promise<T>` | 實際的工作;成功或失敗都在 `finally` 釋放名額 |

對照 Go 的 worker pool:Go 用固定數量的 goroutine 從 channel 取工作;JS 沒有執行緒池,改用「計數 + 等待佇列」限制同時進行的 Promise 數量,效果相同。社群套件 `p-limit` 就是這個模式。

## 執行

```sh
cd javascript && node --test 28_concurrency_limiter/*.test.js
```
