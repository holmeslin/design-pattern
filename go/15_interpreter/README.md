# Interpreter(直譯器)

為某個語言定義文法的物件表示,並提供一個直譯器來解釋該語言中的句子。

## 角色

| 角色 | 對應型別 | 說明 |
| --- | --- | --- |
| AbstractExpression | `Expr` | 定義所有運算式節點共同的求值介面(`Eval`)。 |
| TerminalExpression | `Num` | 文法中的終結節點,代表一個常數數值。 |
| NonterminalExpression | `Add`、`Mul` | 文法中的非終結節點,由子運算式組合而成,分別代表加法與乘法。 |

範例中沒有實作 parser,運算式樹(如 `(1 + 2) * 3`)是以程式手動組出,
重點在展示各節點如何遞迴呼叫 `Eval` 完成求值。

## 執行

```sh
go test ./...
```
