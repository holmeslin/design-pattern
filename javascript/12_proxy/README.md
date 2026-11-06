# Proxy(代理)

提供一個**替身**物件來控制對真物件的存取,呼叫端用起來和真物件一樣。

本範例:圖片的快取代理,第一次 `load()` 真的載入,之後直接回傳快取結果。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Subject** | 「有 `load()` 的物件」這個形狀 | 真物件與代理共用的介面 |
| **RealSubject** | `RealImage` | 真正做事的物件 |
| **Proxy** | `withCache(image)` 回傳的 `Proxy` | 攔截 `load`,第一次轉給真物件,之後回傳快取 |

JS 原生已內建:`Proxy` 就是語言層級的代理,不必手寫一個實作相同介面的代理類別。

注意:`Proxy` 和 `#private` 不相容,方法的 `this` 若是 proxy,讀 `#field` 會 `TypeError`,所以 `load` 對 `target` 呼叫(不經 proxy);`filename` 用公開欄位,是因為 trap 寫在類別外面,讀不到 `#filename`。

## 執行

```sh
cd javascript && node --test 12_proxy/*.test.js
```
