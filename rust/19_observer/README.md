# Observer(觀察者)

物件狀態改變時,**自動通知所有訂閱者**,發佈者不需要知道訂閱者是誰。

本範例:氣象站。`Station` 收到新氣溫後,依訂閱順序通知每個 observer,並收集它們的回傳值。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Subject** | `Station` | 保存訂閱清單,狀態改變時逐一通知 |
| **Observer** | 閉包 `Fn(i32) -> String` | 收到氣溫後做自己的事,回傳一段字串 |
| **Subscription** | `SubscriptionId` | `subscribe` 回傳的憑證,交給 `unsubscribe` 退訂 |

## Rust 慣用寫法

observer 就是閉包(`Box<dyn Fn(i32) -> String>`),不需要另外定義 Observer trait。`subscribe` 接受任何 `impl Fn(i32) -> String + 'static`,所以閉包不能借用呼叫端的區域變數(`'static` 的限制)。

因為型別是 `Fn` 而不是 `FnMut`,observer 無法直接修改捕獲的變數;要累積狀態得自己用 `Cell` 或 `RefCell`。

退訂靠 `SubscriptionId`,不是比較閉包本身(閉包無法比較相等)。`SubscriptionId` 的欄位是私有的,只有 `subscribe` 能產生;`unsubscribe` 遇到不認得的 id 會直接忽略。

## 執行

```sh
cd rust && cargo test -p observer
```
