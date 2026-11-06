# Cancellation(取消)

用一個可傳遞的 signal 通知長時間操作「別做了」,操作自己負責收尾並以 `signal.reason` 拒絕,呼叫端靠 reason 分辨取消原因。

本範例:模擬慢速 API(`slowApi`),搭配逾時(`withTimeout`)與使用者主動取消。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Controller** | `AbortController`(使用者呼叫 `abort()`)、`AbortSignal.timeout(ms)` | 發出取消訊號 |
| **Signal** | `AbortSignal` | 唯一被傳遞的東西;`AbortSignal.any([...])` 把多個來源合併成一個 |
| **Cancellable task** | `slowApi(ms, { signal })` | 檢查 `signal.aborted`、監聽 `abort`,取消時 `clearTimeout` 並 reject `signal.reason` |
| **Timeout wrapper** | `withTimeout(task, ms, { signal })` | 合併逾時與使用者 signal,交給 task |

`AbortSignal` 是 JS 版的 Go `context.Context`:同樣一路往下傳、同樣用來取消與逾時;差別是 JS 的 signal 只管取消,不攜帶值。

`TimeoutError` / `AbortError` 讓呼叫端分辨「逾時」與「主動取消」:`AbortSignal.timeout()` 的 reason 是 `TimeoutError`,`controller.abort()` 預設的 reason 是 `AbortError`。

JS 原生已內建:`AbortController`、`AbortSignal.timeout()`、`AbortSignal.any()` 都是平台 API,本範例只負責把它們接起來。

## 執行

```sh
cd javascript && node --test 26_cancellation/*.test.js
```
