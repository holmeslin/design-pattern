# Chain of Responsibility(責任鏈)

責任鏈模式將多個處理者串成一條鏈,請求沿著鏈依序傳遞,直到被處理或拒絕為止,發送端無需知道最終是哪個處理者處理了請求。

## 角色

| 角色 | 對應型別 | 說明 |
| --- | --- | --- |
| Handler(處理者介面) | `Handler` | 定義 `Handle` 與 `SetNext`,是鏈中每個環節共同遵守的介面 |
| Base Handler(基礎處理者) | `base` | 內嵌於各具體環中,持有 `next` 並提供轉發請求的輔助方法 |
| Concrete Handler(具體處理者) | `authHandler`、`rateLimitHandler`、`bodyHandler` | 各自驗證請求的一個面向,通過後轉發給下一環 |

## 執行

```sh
go test ./...
```
