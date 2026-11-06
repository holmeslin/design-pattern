# Builder(建造者)

把複雜物件的建構拆成一連串步驟,逐步收集資訊,最後一次產出結果;同一套流程可以組出不同的產品。

本範例:SQL 查詢建構器,鏈式呼叫 `select` / `from` / `where_clause`,最後 `build` 出 SQL 字串。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Builder** | `QueryBuilder` | 提供 `select`、`from`、`where_clause` 逐步收集建構所需的資訊 |
| **Product** | `String`(SQL 字串) | 最終產出的物件,由 `build()` 回傳 |
| **Director** | 呼叫端(測試) | 決定呼叫順序與參數,組合出想要的結果 |

## Rust 慣用寫法

消耗型 builder:每個步驟都是 `self -> Self`,取得所有權再回傳,所以鏈式呼叫不需要 `&mut`,也不會留下半成品被誤用。`where` 是關鍵字,所以方法叫 `where_clause`。

## 執行

```sh
cd rust && cargo test -p builder
```
