# Pub/Sub(發布訂閱)

Pub/Sub 是一種並發模式:發布者(Publisher)只將訊息交給中介者(Broker),不需要知道誰會收到;訂閱者(Subscriber)各自透過獨立的 channel 接收訊息,彼此互不干擾。

## 角色

| 角色 | 對應程式碼 |
| --- | --- |
| Broker(訊息中樞) | `Broker` struct |
| Publisher(發布者) | `Broker.Publish` |
| Subscriber(訂閱者) | `Broker.Subscribe` 回傳的 channel |

## 執行

```sh
go test -race ./...
```
