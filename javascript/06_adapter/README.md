# Adapter(轉接器)

把一個類別的介面**轉成呼叫端期待的另一種介面**,讓原本不相容的東西能一起工作。

本範例:我們的程式用「分」付款,第三方 `StripeSDK` 只收「元」,`StripeAdapter` 負責換算。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Target** | `pay(cents)` 這個形狀 | 呼叫端期待的介面 |
| **Adaptee** | `StripeSDK` | 既有、不能改的第三方類別,`charge(dollars)` |
| **Adapter** | `StripeAdapter` | 實作 Target,把分轉成元後委派給 Adaptee |

## 執行

```sh
cd javascript && node --test 06_adapter/*.test.js
```
