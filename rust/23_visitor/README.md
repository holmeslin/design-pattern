# Visitor(訪問者)

把「對一群物件做的操作」從物件本身**抽出來**放進獨立的 visitor,新增操作不必修改物件的型別。

本範例:計算形狀面積。`Circle`、`Rectangle` 各自的 `accept` 把自己交給 `AreaVisitor`,由它算出面積。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Visitor** | `trait Visitor` | 每種形狀一個 `visit_*` 方法 |
| **ConcreteVisitor** | `AreaVisitor` | 實作一種操作(算面積) |
| **Element** | `trait Shape` | 宣告 `accept(&self, v: &dyn Visitor)` |
| **ConcreteElement** | `Circle`、`Rectangle` | `accept` 呼叫對應的 `visit_*`,把自己傳進去 |

## Rust 慣用寫法

`accept` 內呼叫哪個 `visit_*` 由元素的具體型別決定,而實際執行哪段程式由傳入的 visitor 決定(`&dyn Visitor`,執行期分派);合起來就是雙重分派。

不過 Rust 更常用 `enum Shape { Circle(..), Rectangle(..) }` 加 `match` 取代 Visitor:

- 新增「操作」只要多寫一個函式,不需要 `Visitor` trait 與 `accept`。
- 新增 variant 時,凡是沒有 `_` 萬用分支的 `match` 都會編譯失敗,編譯器會指出所有沒處理到的地方。
- 代價是型別集合是封閉的:要新增形狀就得修改 enum 本身。

## 執行

```sh
cd rust && cargo test -p visitor
```
