# Fearless Concurrency(無畏並行)

在 safe Rust 裡,資料競爭(data race)是編譯錯誤而不是執行期的 bug:能不能跨 thread 傳遞、共用,由型別系統(`Send` / `Sync`)在編譯期檢查。

本範例:把平方和拆給多個 thread 平行計算(`thread::scope` + `mpsc::channel`),以及用 `Arc<Mutex<u64>>` 讓多個 thread 累加同一個計數器。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **scoped thread** | `sum_squares_parallel` 裡的 `s.spawn` | 直接借用呼叫端的切片,scope 結束前一定被 join |
| **channel** | `mpsc::channel` | worker 把部分和送給主執行緒,主執行緒收集後加總 |
| **共享狀態** | `Arc<Mutex<u64>>`(`count_with_threads`) | 多個 thread 共同擁有並互斥地修改同一個計數器 |
| **編譯期檢查** | `Send` / `Sync` | 不安全的跨 thread 用法(例如 `Rc`)編譯不過 |

## Rust 慣用寫法

**對照 Go 的 worker pool 與 channel。** Go 的 goroutine 與 channel 概念相近,但共用資料有沒有加鎖,編譯器不會檢查,要靠 `go test -race` 在執行期才抓得到;Rust 則是 `Send` / `Sync` 讓資料競爭在編譯期就被擋下。`lib.rs` 的 doctest 放了一對範例,只差 `Arc` 換成 `Rc`:`Arc` 能搬進 `thread::spawn`,`Rc` 的參考計數不是原子操作、不是 `Send`,所以標成 `compile_fail`。stable 版 rustdoc 不檢查 `compile_fail` 的錯誤碼,任何編譯錯誤都會讓它「通過」,所以才需要能編譯的對照組,證明失敗確實來自那一處。

**`thread::scope`:借用堆疊上的資料。** `thread::spawn` 要求 `'static`,thread 不能借用區域變數,只好用 `Arc` 共用(`count_with_threads`)。`thread::scope` 保證所有 scoped thread 在 scope 結束前被 join,所以 `sum_squares_parallel` 的 worker 可以直接借用 `nums` 的切片,不需要 `Arc`、也不需要複製資料。

**channel 與 `drop(tx)`。** 每個 worker 拿的是 `tx.clone()`;`rx.iter()` 要等到「所有」發送端都被 drop 才會結束。如果沒有 drop 原本那個 `tx`,channel 永遠不會關閉,收集會一直卡住。

**錯誤處理。** `lock()` 取得鎖失敗只代表別的 thread 在持鎖時 panic(鎖中毒),這裡的操作不會讓資料不一致,所以用 `unwrap_or_else(|e| e.into_inner())` 取回資料繼續用。`join()` 失敗代表 worker panic,那是程式有 bug、沒有可恢復的狀態,所以用 `expect`;`send` 只會在 receiver 已被 drop 時失敗,而 receiver 活到 scope 結束之後,所以同樣是 `expect`。`workers == 0` 是呼叫端的錯誤用法,用 `assert!` 直接 panic。

**限制。** `Send` / `Sync` 只防資料競爭,不防死鎖或邏輯上的競爭條件(例如多個 thread 先讀再寫、順序不對)。

## 執行

```sh
cd rust && cargo test -p fearless_concurrency
```
