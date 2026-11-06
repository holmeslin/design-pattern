# Visitor(訪問者)

把「要對一組物件做什麼操作」從物件本身抽離成獨立的 Visitor,新增操作時不必修改既有的 Element 型別。

本範例:圖形(Circle、Rectangle)的面積計算。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Element** | 有 `accept(visitor)` 的物件這個形狀 | 把自己交給 Visitor 處理 |
| **ConcreteElement** | `Circle`、`Rectangle` | `accept` 內回呼 visitor 對應自己型別的方法(double dispatch) |
| **Visitor** | 有 `visitCircle` / `visitRectangle` 的物件這個形狀 | 針對每一種 ConcreteElement 宣告一個操作 |
| **ConcreteVisitor** | `areaVisitor` | 計算面積 |

JS 沒有 interface,Element 與 Visitor 都只是約定的形狀(duck typing);Visitor 沒有狀態,所以用物件字面量而不是 class。

## 執行

```sh
cd javascript && node --test 23_visitor/*.test.js
```
