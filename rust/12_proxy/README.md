# Proxy(代理)

用一個與真實物件**介面相同的替身**控制對它的存取,可以在存取前後加上延遲載入、快取、權限檢查等行為。

本範例:圖片的延遲載入(virtual proxy),`ProxyImage` 直到第一次 `load()` 才建立 `RealImage`,之後回傳快取結果。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Subject** | `trait Image` | 真實物件與代理共同的介面 |
| **RealSubject** | `RealImage` | 真正載入圖片的昂貴物件 |
| **Proxy** | `ProxyImage` | 持有 `OnceCell<RealImage>`,第一次使用才建立並委派 |

## Rust 慣用寫法

`std::cell::OnceCell` 提供延遲初始化:`real` 一開始是空的,第一次 `get_or_init` 時才建立 `RealImage`,之後只會讀到同一個。它靠 interior mutability 讓 `load` 只需要 `&self`,不必改成 `&mut self`,所以代理與真實物件能共用同一個 `Image` 介面。這是 virtual proxy:真正的物件直到第一次使用才建立。`OnceCell` 不是 `Sync`,多執行緒要改用 `std::sync::OnceLock`。

## 執行

```sh
cd rust && cargo test -p proxy
```
