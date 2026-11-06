# Command(命令)

把「一個操作」包成**物件**,就能記錄、排隊、延後執行,也能反向撤銷。

本範例:帳戶存款(`Deposit`)可執行也可撤銷,`History` 記錄已執行的命令並支援 `undo_last`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Command** | `trait Command` | 宣告 `execute` 與 `undo` |
| **ConcreteCommand** | `Deposit` | 封裝存款金額,知道如何執行與還原 |
| **Receiver** | `Account` | 真正被操作的對象 |
| **Invoker** | `History` | 執行命令並以堆疊保存,`undo_last` 撤銷最近一筆;空的時候不動作 |

## Rust 慣用寫法

command **不持有** `&mut Account`,改成呼叫 `execute(&self, account: &mut Account)` 時才把帳戶傳進來。借用規則不允許 `History` 裡存著多個對同一帳戶的可變參考,這是 Rust 版和 Go、JS 版最大的差異:那兩個語言的 command 可以直接記住 receiver,Rust 則讓 receiver 的借用只活在單次呼叫裡。`History` 以 `Vec<Box<dyn Command>>` 保存不同種類的命令。`Deposit` 用 `wrapping_add` / `wrapping_sub`,讓 `undo` 一定是 `execute` 的精確反運算且溢位時不 panic,但這只是示範用的捷徑(溢位時餘額會悄悄回繞),真實金額應改用 `checked_*` 並回傳錯誤。

## 執行

```sh
cd rust && cargo test -p command
```
