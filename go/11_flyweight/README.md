# Flyweight(享元)

享元模式透過共享不可變的內部狀態物件,減少大量相似物件所造成的記憶體浪費。

## 角色

| 角色 | 對應型別 | 說明 |
| --- | --- | --- |
| Flyweight(享元) | `Format` | 可共享的不可變物件,持有字型與大小 |
| FlyweightFactory(享元工廠) | `Factory` | 依鍵值快取並配發共享的 `*Format` |

## 執行

```sh
go test ./...
```
