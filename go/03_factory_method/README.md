# Factory Method(工廠方法)

定義一個建立物件的介面,將實際實例化哪個具體型別的決定權交給子類別或工廠函式,讓呼叫端只依賴共同介面而不需知道建構細節。

## 三個角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Product | `Notifier` | 定義所有具體產品共同的行為介面 |
| ConcreteProduct | `emailNotifier`、`smsNotifier` | 實作 `Notifier` 介面的具體型別 |
| Creator | `NewNotifier` | 依據參數建立對應具體產品的工廠函式 |

## 執行

```sh
go test ./...
```
