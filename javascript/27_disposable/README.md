# Disposable(可釋放資源)

資源實作 `Symbol.dispose` / `Symbol.asyncDispose`,用 `using` / `await using` 宣告後,離開 block 時語言保證自動釋放,不必手寫 `try / finally`。

本範例:`Connection`(同步關閉)與 `FileLock`(非同步解鎖)。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Disposable** | `Connection`,有 `[Symbol.dispose]()` | 同步釋放:推入 `close <name>` |
| **AsyncDisposable** | `FileLock`,有 `[Symbol.asyncDispose]()` | 非同步釋放:推入 `unlock <name>`;`acquire` 先推入 `lock <name>` |
| **Scope** | `using` / `await using` 所在的 block | 離開 block(正常或丟例外)時,依宣告的**相反順序**呼叫釋放方法 |

ES2026 Explicit Resource Management:同一個概念在別的語言各有對應,對照 Go 的 `defer`、C# 的 `using`、Python 的 `with`。

需要 Node 24 以上(`package.json` 的 `engines` 已標明)。

## 執行

```sh
cd javascript && node --test 27_disposable/*.test.js
```
