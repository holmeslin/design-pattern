# Interpreter(直譯器)

把一個簡單語言的**文法規則做成物件樹**,每個節點知道怎麼解讀自己,整棵樹遞迴求值。

本範例:整數四則運算的子集(數字、加、乘),`mul(add(num(1), num(2)), num(3))` 即 `(1 + 2) * 3`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractExpression** | 「有 `interpret()` 的物件」這個形狀 | 所有節點共用的求值介面 |
| **TerminalExpression** | `num(n)` | 葉節點,直接回傳數值 |
| **NonterminalExpression** | `add(left, right)`、`mul(left, right)` | 先對左右子節點求值,再組合結果 |

方法叫 `interpret` 而不是 `eval`,避免和全域 `eval` 混淆;本範例直接手寫語法樹,沒有 parser。

## 執行

```sh
cd javascript && node --test 15_interpreter/*.test.js
```
