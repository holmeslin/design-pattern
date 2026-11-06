# Iterator(迭代器)

提供一種**依序走訪集合元素**的方法,而不暴露集合的內部結構。

本範例:`Ring<T>` 包著一個 `Vec<T>`,對外只提供走訪,內部儲存方式呼叫端看不到。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Iterator** | 標準庫的 `Iterator` trait | 定義 `next`,依序產出元素 |
| **Aggregate** | `Ring<T>` | 持有元素的集合;`iter()` 與 `IntoIterator` 是它建立迭代器的入口 |
| **ConcreteIterator** | `std::slice::Iter`(由 `iter()` 以 `impl Iterator` 隱藏) | 記錄走訪位置並產出 `&T` |

## Rust 慣用寫法

語言內建 `Iterator` trait,不必自己定義迭代器介面。為 `&Ring<T>` 實作 `IntoIterator` 之後,就能直接寫 `for x in &ring`,而且不會消耗 `ring`;`map`、`filter`、`sum` 等 adapter 也都自動可用。`iter()` 回傳 `impl Iterator<Item = &T>`,隱藏具體型別,日後換掉內部儲存也不影響呼叫端。這些 adapter 在編譯期單態化、可以內聯,是零成本抽象。

## 執行

```sh
cd rust && cargo test -p iterator
```
