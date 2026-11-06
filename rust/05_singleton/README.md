# Singleton(單例)

保證整個程式只有一個實例,並提供一個全域存取點來取得它。

本範例:全域設定 `Config`(key/value),透過 `instance()` 取得唯一實例,可跨執行緒讀寫。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Singleton** | `Config` | 只會被實例化一次的型別;欄位私有,外部無法自行建構 |
| **存取點** | `instance()` | 唯一取得實例的函式,背後是 `static INSTANCE: LazyLock<Config>` |

## Rust 慣用寫法

`LazyLock` 在第一次存取時才初始化,而且保證只初始化一次,即使多個執行緒同時呼叫 `instance()` 也一樣(對應 Go 的 `sync.Once`),不必手寫雙重檢查鎖。

Rust 不允許未同步的可變全域變數,所以全域狀態要可寫入,就必須包 `Mutex`,這是編譯器強制的。`Mutex` 中毒(持鎖的執行緒 panic)時,用 `lock().unwrap_or_else(|e| e.into_inner())` 取回資料,而不是 `unwrap` 讓每個呼叫端跟著 panic。

## 執行

```sh
cd rust && cargo test -p singleton
```
