# Prototype(原型)

透過複製既有物件(原型)來建立新物件,避免重複執行昂貴或繁瑣的初始化流程。

本範例:設定物件 `Config`(含 `String` 與 `Vec<String>` 欄位),複製後修改副本,原型不受影響。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Prototype** | `Config` | 定義可被複製的物件,內含需要深拷貝的欄位 |
| **clone 操作** | `Config::clone()`(`#[derive(Clone)]`) | 產生與原型完全獨立的副本 |
| **Client** | 呼叫端(測試) | 複製原型後再調整副本,而不是重新建構 |

## Rust 慣用寫法

`Clone` trait 就是 Prototype,不需要自己定義 `Clone()` 方法。`#[derive(Clone)]` 會對擁有所有權的欄位(`String`、`Vec<T>`)做深拷貝。要注意:若欄位是 `Rc<T>`(或 `Arc<T>`),clone 只會複製指標,副本與原型會共享同一份資料,並不是深拷貝。

## 執行

```sh
cd rust && cargo test -p prototype
```
