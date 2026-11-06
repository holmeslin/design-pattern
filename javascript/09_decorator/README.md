# Decorator(裝飾者)

**動態地為物件加上額外行為**,不改原本的類別,而且可以一層一層疊。

本範例:HTTP handler 的 middleware,`logging`、`auth` 各自包一層,`logging(auth(base))` 輸出 `[log] [auth] ok`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Component** | 「`(req) => string` 的函式」這個形狀 | 被裝飾者與裝飾者共用的介面 |
| **ConcreteComponent** | `base` | 最內層的原始行為 |
| **Decorator** | `logging`、`auth` | 接收一個 handler,回傳同形狀、多了一層行為的新 handler |

JS 原生已內建:函式是一級公民,裝飾者就是高階函式,不需要「持有 component 的裝飾者類別」。

注意:TC39 的 `@decorator` 語法尚未在 Node 原生啟用,本範例不用。

## 執行

```sh
cd javascript && node --test 09_decorator/*.test.js
```
