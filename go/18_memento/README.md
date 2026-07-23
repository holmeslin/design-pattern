# Memento(備忘錄)

備忘錄模式在不破壞封裝的前提下,擷取並保存物件的內部狀態,以便日後將該物件還原到先前的狀態。

## 角色

| 角色 | 說明 | 對應程式碼 |
| --- | --- | --- |
| Originator | 擁有需要被保存與還原之狀態的物件 | `Editor` |
| Memento | 儲存狀態快照的不透明物件,只有 Originator 能讀寫其內容 | `memento` |
| Caretaker | 負責保存 Memento(例如以堆疊管理),但不窺探其內容 | `example_test.go` 內的呼叫端 |

## 執行

```sh
go test ./...
```
