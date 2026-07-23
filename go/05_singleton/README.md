# Singleton(單例)

Singleton 模式保證一個類別在整個程式生命週期中只有一個實例,並提供一個全域存取點來取得它。

本範例使用 Go 標準庫的 `sync.Once` 來實作,而不是像 C++/Java 常見的手寫「雙重檢查鎖」
(double-checked locking)。`sync.Once` 是 Go 慣用(idiomatic)的做法:`Do` 方法保證
傳入的初始化函式只會被執行一次,即使多個 goroutine 同時呼叫 `Instance()` 也是
goroutine 安全的,不需要自己處理鎖與記憶體可見性的細節。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Singleton | `registry` | 只會被實例化一次的型別 |
| 存取點(Access Point) | `Instance()` | 對外唯一取得實例的函式,內部以 `sync.Once` 保護初始化 |

## 執行

```sh
go test ./...
```
