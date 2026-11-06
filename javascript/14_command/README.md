# Command(命令)

把「一個操作」**包成物件**,就能儲存、排隊、記錄,也能反向撤銷。

本範例:帳戶存款指令,`History` 執行指令並記錄,`undoLast()` 撤銷最近一次。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Command** | `{ execute, undo }` 這個形狀 | 把操作與撤銷包在一起 |
| **ConcreteCommand** | `deposit(account, amount)` 回傳的物件 | 對 `Account` 加錢 / 扣回 |
| **Receiver** | `Account` | 真正被操作的對象 |
| **Invoker** | `History` | 執行指令並留下紀錄,負責 `undoLast()` |

JS 慣用:command 就是 `{ execute, undo }` 物件字面值(閉包捕捉 `account` 與 `amount`),不需要 Command 類別;空歷史時 `undoLast()` 不動作。

## 執行

```sh
cd javascript && node --test 14_command/*.test.js
```
