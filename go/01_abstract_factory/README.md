# Abstract Factory(抽象工廠)

一個工廠介面**成套產出「相關的一族產品」**,呼叫端只認工廠與產品介面,換一整族只換一個工廠。

本範例:UI 主題工廠,產出成套的 Button + Checkbox(Dark / Light)。

## 三個角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **AbstractFactory** | `GUIFactory` | 宣告「成套產出哪些產品」 |
| **ConcreteFactory** | `DarkFactory`、`LightFactory` | 產出同一族的具體產品 |
| **Product** | `Button`、`Checkbox` | 產品介面;具體產品細節住自己這裡 |

呼叫端 `renderUI` 只認 `GUIFactory`,要換整套主題就換一個工廠,不改呼叫端。

## 執行

```sh
go test ./...
```
