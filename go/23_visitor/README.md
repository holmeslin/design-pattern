# Visitor(訪問者)

把「要對一組物件做什麼操作」從物件本身抽離成獨立的 Visitor,新增操作時不必修改既有的 Element 型別。

## 角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **Element**(介面) | `Shape` | 定義 `Accept(Visitor)`,把自己交給 Visitor 處理 |
| **ConcreteElement** | `Circle`、`Rectangle` | 具體的資料型別,`Accept` 內回呼 Visitor 對應的方法(double dispatch) |
| **Visitor**(介面) | `Visitor` | 定義針對每一種 ConcreteElement 的操作:`VisitCircle`、`VisitRectangle` |
| **ConcreteVisitor** | `AreaVisitor` | 具體操作實作,本範例用來計算面積 |

## 執行

```sh
go test ./...
```
