# Interpreter(直譯器)

為一個簡單的語言定義**文法的物件表示**(語法樹),並寫一個直譯器走訪語法樹來求值。

本範例:整數的加法與乘法運算式,例如 `(1 + 2) * 3` 求值為 `9`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractExpression** | `enum Expr` | 所有運算式共同的型別,`interpret` 定義求值規則 |
| **TerminalExpression** | `Expr::Num` | 文法的終結符:直接是一個數字 |
| **NonterminalExpression** | `Expr::Add`、`Expr::Mul` | 由子運算式組成,遞迴求值再組合結果 |

## Rust 慣用寫法

AST 用 `enum` + `Box` 表示遞迴結構,`match` 求值,是 Rust 寫直譯器的標準做法。每條文法規則是一個 variant,`Box` 讓遞迴型別有固定大小;整個文法收進同一個 `enum`,不像 Go 版為每個節點各建一個型別、再靠介面分派。`match` 必須窮舉所有 variant,日後新增規則時,漏寫求值分支會直接編譯失敗。

運算用 `wrapping_add` / `wrapping_mul`,溢位時回繞而不是 panic(程式碼裡有 `shortcut:` 註解);輸入不可信時應改用 `checked_*` 並回傳 `Result`。

## 執行

```sh
cd rust && cargo test -p interpreter
```
