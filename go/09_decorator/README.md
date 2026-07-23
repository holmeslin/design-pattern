# Decorator(裝飾者)

在不修改原始物件的前提下,動態為其疊加新的行為或職責。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Component | `Handler` | 定義被裝飾對象的統一介面 |
| ConcreteComponent | `base` | 最基本、未被裝飾過的實作 |
| Decorator | `Logging`、`Auth` | 包裝一個 `Handler`,在其結果前疊加前綴 |

`h := Logging(Auth(base{}))` 呼叫時由外而內(Logging -> Auth -> base),
但結果字串是由內而外組裝:`base` 先產生 `"ok"`,再依序被 `Auth`、`Logging`
加上前綴,最終得到 `"[log] [auth] ok"`。

這其實就是 `net/http` middleware 的形狀:`func(http.Handler) http.Handler`,
一層層把 `http.Handler` 包起來,在呼叫內層 handler 前後加上額外行為。

## 執行

```sh
go test ./...
```
