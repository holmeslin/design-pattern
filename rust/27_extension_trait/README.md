# Extension Trait(擴充 trait)

想替**不屬於自己的型別**加方法,就定義一個 trait,再為那個型別(或所有符合條件的型別)實作它。

本範例:替 `str` 加上 `to_title_case()`,替所有元素為 `i64` 的 `Iterator` 加上 `sum_squares()`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **擴充 trait** | `TitleCase`、`SumSquares` | 宣告要加上去的方法 |
| **被擴充的型別** | `str`、任何 `Iterator<Item = i64>` | 標準函式庫的型別,不是我們定義的 |
| **實作** | `impl TitleCase for str`、`impl<I: Iterator<Item = i64>> SumSquares for I {}` | 把方法掛上去;後者是 blanket impl |

## Rust 慣用寫法

**為什麼要用 trait。** `str` 與 `Iterator` 都定義在標準函式庫,我們不能直接在它們的定義裡加方法,但可以為它們實作自己定義的 trait(孤兒規則:trait 或型別至少有一個是自己的)。呼叫端寫起來就像內建方法:`"hello rust world".to_title_case()`、`(1..=5i64).sum_squares()`。

**blanket impl。** `impl<I: Iterator<Item = i64>> SumSquares for I {}` 為「所有」符合約束的型別一次實作;方法本體是 trait 的預設實作,所以 impl 區塊是空的。`Range`、`vec::IntoIter` 與自訂的 `Iterator` 只要 `Item = i64` 就都有 `sum_squares`。

**使用端要 `use` 這個 trait。** trait 的方法只有在 trait 位於作用域內時才看得到;沒有 `use` 就會得到「找不到方法」的編譯錯誤。`itertools::Itertools` 就是這個模式:`use itertools::Itertools;` 之後,所有 `Iterator` 都多出 `join`、`sorted` 等方法。

**對照。** Go 不能替內建型別或別人的型別加方法,只能寫以它為參數的函式;JS 可以改 `prototype`,但會影響全域,Rust 的擴充則只在 `use` 了該 trait 的地方生效。

## 執行

```sh
cd rust && cargo test -p extension_trait
```
