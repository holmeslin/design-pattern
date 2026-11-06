# Flyweight(享元)

把大量物件中**相同的內部狀態抽出來共享**,用工廠快取,同樣的資料只存一份,以節省記憶體。

本範例:文字格式(字型 + 大小),同樣的 `("Arial", 12)` 不論被多少個字元使用,都共用同一個 `CharFormat`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Flyweight** | `CharFormat` | 可共享的內部狀態(字型、大小),建立後不再變動 |
| **FlyweightFactory** | `FormatFactory` | 以 `(font, size)` 為 key 快取,命中就回傳同一個 `Rc` |

## Rust 慣用寫法

`Rc<CharFormat>` 負責共享:`get` 回傳的是同一份資料的另一個引用計數指標,可用 `Rc::ptr_eq` 驗證是同一個物件,`Rc::strong_count` 看有幾個持有者(工廠的 cache 本身也算一份)。`Rc<T>` 只給出 `&T`,所以共享的物件天生不可變,正是 Flyweight 需要的性質(對應 JS 版要用 `Object.freeze` 才做得到)。`Rc` 不是 `Send`,跨執行緒共享要改用 `Arc`。

## 執行

```sh
cd rust && cargo test -p flyweight
```
