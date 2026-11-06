# Async Iterator Pipeline(非同步迭代器管線)

把資料處理拆成一串小階段,每個階段接收上游的 async iterable、產出新的 async iterable,串起來一次處理一筆,不必先把全部資料載入記憶體。

本範例:日誌管線,從一串日誌中只留 `ERROR`、去掉前綴、取前 2 筆。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Source** | 測試內的 `source()` | 依序產出日誌;被提早關閉時在 `finally` 釋放資源 |
| **Stage** | `filter`、`map`、`take` | 各自是 `async function*`,只負責一件事 |
| **Pipeline** | `take(map(filter(source(), …), …), 2)` | 階段串接;消費端用 `Array.fromAsync` 或 `for await` 拉資料 |

對照 Go 的 channel pipeline(`go/25_pipeline`):Go 用 goroutine + channel 串接各階段;JS 用 async generator,是 **pull-based**:下游沒來拉,上游就不會往前跑,天生有 backpressure,不需要自己管 buffer 大小。`take` 取滿後 `return`,`for await` 的收尾會一路呼叫上游的 `return()`,讓每一層的 `finally` 都有機會執行。

JS 原生已內建:同步版的 Iterator helpers(`iter.filter().map().take()`)已經內建;async 版的 helpers 還在提案階段,所以這裡手寫三個階段。

## 執行

```sh
cd javascript && node --test 25_async_iterator_pipeline/*.test.js
```
