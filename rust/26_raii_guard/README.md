# RAII Guard(資源守衛)

把資源的取得與釋放綁在一個值的生命週期上:建立值就取得資源,值被 drop 就自動釋放,呼叫端不需要記得手動關閉。

本範例:`open` 取得一個連線守衛並記錄 `open <name>`,守衛被 drop 時記錄 `close <name>`,用 log 驗證釋放的時機與順序。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Guard** | `ConnectionGuard` | 持有資源(連線名稱與共用的 log);使用資源的方法掛在它身上 |
| **取得** | `open(name, &log)` | 取得資源並回傳 guard |
| **釋放** | `impl Drop for ConnectionGuard` | guard 被 drop 時釋放資源 |

## Rust 慣用寫法

**對照 Go 的 `defer`、JS 的 `using`。** Go 要在取得資源後自己寫 `defer conn.Close()`,JS 要寫 `using`;Rust 的資源釋放由所有權決定,不需要額外語法:值在擁有它的 scope 結束時被 drop,`Drop::drop` 就會執行。忘記釋放不會發生,因為沒有「釋放」這個步驟可以忘。標準函式庫的 `MutexGuard`(drop 時解鎖)與 `File`(drop 時關閉檔案)都是這個模式。

**釋放的時機。** 同一個 scope 內的區域變數依宣告的相反順序 drop(測試裡 `b` 先於 `a` 關閉);也可以用 `drop(x)` 提早釋放;panic 展開(unwind)時一樣會 drop,測試 `drops_on_panic` 驗證了這點。

**限制。** `Drop::drop` 不能回傳錯誤,所以像「關閉時可能失敗」的資源,通常另外提供一個會回傳 `Result` 的 `close(self)`,`Drop` 只當作最後的保險。另外 `drop` 在 panic 展開時才會執行;若 profile 設成 `panic = "abort"`,程式直接終止,不會執行 `Drop`。

## 執行

```sh
cd rust && cargo test -p raii_guard
```
