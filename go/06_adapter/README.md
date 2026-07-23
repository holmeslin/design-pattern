# Adapter(轉接器)

Adapter 模式把一個介面不相容的既有類別,包裝成呼叫端期望的介面,讓原本無法協同運作的雙方得以合作,而不需修改任何一方原本的程式碼。

## 三個角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Target(目標介面) | `PaymentProcessor` | 呼叫端依賴的介面,以「分」為單位收款 |
| Adaptee(被轉接者) | `StripeSDK` | 既有、介面不相容的第三方服務,以「元」為單位請款 |
| Adapter(轉接器) | `StripeAdapter` | 包裝 `StripeSDK`,實作 `PaymentProcessor`,負責單位轉換與轉呼叫 |

## 執行

```sh
go test ./...
```
