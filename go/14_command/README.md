# Command(命令)

命令模式將一個請求封裝成獨立物件,使呼叫者與實際執行請求的物件解耦,並可支援排隊、記錄與復原(Undo)等操作。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Command(命令介面) | `Command` | 定義 `Execute` 與 `Undo` 兩個行為 |
| ConcreteCommand(具體命令) | `DepositCommand` | 封裝存款請求與其接收者、參數 |
| Receiver(接收者) | `Account` | 實際持有狀態,知道如何完成請求 |
| Invoker(呼叫者) | `History` | 觸發命令執行,並記錄歷史以支援復原 |

## 執行

```sh
go test ./...
```
