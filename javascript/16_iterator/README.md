# Iterator(迭代器)

提供一種**依序走訪集合元素**的方式,呼叫端不必知道集合內部怎麼存。

本範例:`Ring` 集合,對外只暴露走訪行為,內部陣列是私有的。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Iterable(Aggregate)** | `Ring` | 實作 `[Symbol.iterator]()`,交出迭代器 |
| **Iterator** | `[Symbol.iterator]()` 回傳的 generator 物件 | 每次 `next()` 給出下一個元素 |
| **Client** | `[...ring]`、`for...of`、`.map(...).toArray()` | 只靠 iterator protocol 走訪 |

JS 原生已內建:iterator protocol(`Symbol.iterator` + generator)是語言本身的一部分,`for...of`、展開運算子、`Array.from`、ES2025 的 Iterator helpers(`map`、`filter`、`take`、`toArray` 等)全部自動支援,不必自己寫 Iterator 類別。對照 Go 的 `iter.Seq`。

## 執行

```sh
cd javascript && node --test 16_iterator/*.test.js
```
