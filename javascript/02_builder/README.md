# Builder(建造者)

把複雜物件的**建構過程拆成一步一步**,呼叫端逐步設定,最後才組出結果。

本範例:SQL 查詢建造者,`select` / `from` / `where` 鏈式呼叫後 `build()` 出字串。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Builder** | `QueryBuilder` | 提供逐步設定的方法,並保管建構中的狀態 |
| **Product** | `build()` 回傳的 SQL 字串 | 最終產物 |

JS 慣用寫法是每個設定方法回傳 `this`(method chaining),不需要獨立的 Director 角色。

## 執行

```sh
cd javascript && node --test 02_builder/*.test.js
```
