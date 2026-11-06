# Singleton(單例)

確保**一個類別只有一個實例**,並提供全域存取點。

本範例:全域設定物件,`config.set` 寫入、`config.get` 讀取,所有 import 的地方看到同一份。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Singleton** | `config` | 全程唯一的實例,內部用 `Map` 保管狀態 |
| **全域存取點** | `import { config }` | 任何檔案 import 都拿到同一個物件 |

JS 原生已內建:ES module 只會被評估一次,module 匯出的物件天生就是單例,不需要 `getInstance()`;Go 版需要 `sync.Once` 來確保只初始化一次。

## 執行

```sh
cd javascript && node --test 05_singleton/*.test.js
```
