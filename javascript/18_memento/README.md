# Memento(備忘錄)

在不破壞封裝的前提下,擷取並保存物件的內部狀態,日後可以還原到先前的狀態。

本範例:文字編輯器,`save` 存下當下內容,編輯後可 `restore` 回去。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Originator** | `Editor` | 擁有狀態,負責產生快照(`save`)與從快照還原(`restore`) |
| **Memento** | `save()` 回傳的凍結物件 `{ content }` | 狀態快照,建立後不可修改 |
| **Caretaker** | 測試中保管 `saved` 的呼叫端 | 只保存 memento,不碰內容 |

`#content` 是 `#private`,外部只能透過 `save` / `restore` 碰狀態;memento 用 `Object.freeze` 防竄改。

## 執行

```sh
cd javascript && node --test 18_memento/*.test.js
```
