# Signals(訊號)

Observer 的現代化版本:把「值」包成 signal,讀取它的 effect / computed 會在執行時被**自動記為相依**,值改變時只重跑真正用到它的那些,不用手動 subscribe。

本範例:購物車,`total = price * qty` 是 computed,一個 effect 記錄每次 `total` 的值。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Signal** | `price`、`qty` | 持有值;`get()` 時登記目前的 effect,`set()` 時通知它們(值相同就不通知) |
| **Computed** | `total` | 由其他 signal 推導出來的值,相依變了就重算 |
| **Effect** | `effect(() => log.push(total.get()))` | 立即執行一次,之後相依變化就重跑 |

和 Observer 的差別:Observer 的訂閱關係由呼叫端手動 `addEventListener`;Signals 的訂閱關係由「誰在 `get()`」自動決定,相依變了也會自己更新。

這個模式正在被主流框架採用:Solid、Preact Signals、Angular Signals、Vue(`ref` / `computed`),另有 TC39 Signals 提案想把它標準化進語言。JS 目前沒有內建,這裡的實作是最小版本(eager、無排程),僅用來說明原理。

## 執行

```sh
cd javascript && node --test 24_signals/*.test.js
```
