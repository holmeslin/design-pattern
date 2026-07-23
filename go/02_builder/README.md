# Builder(建造者)

建造者模式將複雜物件的建構過程拆解成一連串步驟,讓相同的建構流程可以組裝出不同的結果。

## 三個角色

| 角色 | 對應本專案 | 職責 |
| --- | --- | --- |
| Builder(建造者) | `QueryBuilder` | 提供 `Select`、`From`、`Where` 等鏈式方法,逐步收集建構所需的資訊 |
| Product(產品) | `string`(SQL 字串) | 最終被建構出來的物件,由 `Build()` 產出 |
| Director(指揮者) | 呼叫端程式碼(`ExampleQueryBuilder`) | 決定呼叫 Builder 方法的順序與參數,組合出所需的結果 |

## 執行

```sh
go test ./...
```
