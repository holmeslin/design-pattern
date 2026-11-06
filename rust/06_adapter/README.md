# Adapter(轉接器)

把一個既有類別的介面**轉成呼叫端期望的另一個介面**,讓原本不相容的東西能一起運作。

本範例:第三方支付 SDK 只收「元」(`f64`),我們的系統以「分」(`u64`)計價,用 `StripeAdapter` 轉接。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Target** | `trait PaymentProcessor` | 呼叫端期望的介面(`pay(cents)`) |
| **Adaptee** | `StripeSdk` | 既有、介面不相容的類別(`charge(dollars)`) |
| **Adapter** | `StripeAdapter` | 持有 Adaptee,實作 Target,把分換成元後委派 |

## Rust 慣用寫法

Rust 沒有繼承,Adapter 一律用**組合**:`StripeAdapter` 持有一個 `StripeSdk` 欄位,再實作 `PaymentProcessor` 並委派。若 Target trait 與 Adaptee 型別都來自外部 crate,孤兒規則(orphan rule)不允許直接 `impl`,就得靠這種自己的包裝型別(newtype)才能轉接。

## 執行

```sh
cd rust && cargo test -p adapter
```
