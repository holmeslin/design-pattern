# Iterator(迭代器)

迭代器模式讓呼叫者能在不了解聚合物件內部結構的情況下,依序走訪其中的元素。

自 Go 1.23 起,語言原生支援 range-over-func(標準庫型別為 `iter.Seq[T]`),
只要聚合物件回傳這種函式型別,呼叫端就能直接用 `for v := range seq` 走訪,
因此在慣用 Go 中,已不再需要自行定義 `HasNext()` / `Next()` 這類傳統的
Iterator 介面,`iter.Seq[T]` 本身就是語言原生的 Iterator。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Aggregate(聚合) | `Ring` | 持有元素集合,提供建立迭代器的方法 |
| Iterator(迭代器) | `iter.Seq[T]`(標準庫,`All` 回傳) | 語言原生的迭代器型別,搭配 for-range 走訪元素 |

## 執行

```sh
go test ./...
```
