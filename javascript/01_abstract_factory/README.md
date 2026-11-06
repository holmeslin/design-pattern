# Abstract Factory(抽象工廠)

一個工廠**成套產出「相關的一族產品」**,呼叫端只認工廠的形狀,換一整族只換一個工廠。

本範例:UI 主題工廠,產出成套的 Button + Checkbox(Dark / Light)。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractFactory** | 「有 `createButton` / `createCheckbox` 的物件」這個形狀 | 宣告成套產出哪些產品 |
| **ConcreteFactory** | `darkFactory`、`lightFactory` | 產出同一族的具體產品 |
| **Product** | 有 `render()` 的物件 | 產品細節住在自己這裡 |

JS 沒有 interface,AbstractFactory 只是一個約定的形狀(duck typing);Go 版用 `GUIFactory` 介面明確宣告。

## 執行

```sh
cd javascript && node --test 01_abstract_factory/*.test.js
```
