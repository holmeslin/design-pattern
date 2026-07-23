# Mediator(中介者)

中介者模式將物件之間的直接溝通,改由一個中介者集中協調,降低同事物件彼此的耦合度。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Mediator(中介者介面) | `Mediator` | 定義 `Broadcast` 轉發訊息的行為 |
| ConcreteMediator(具體中介者) | `ChatRoom` | 持有所有成員,依加入順序轉發訊息 |
| Colleague(同事) | `User` | 只透過 Mediator 與其他成員互動,不直接依賴彼此 |

## 執行

```sh
go test ./...
```
