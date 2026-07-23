// Package memento 示範 Memento(備忘錄)模式。
//
// Memento 模式的目的是在不破壞封裝的前提下,擷取並保存物件的內部狀態,
// 以便日後能將該物件還原到先前的狀態,常用來實作復原(undo)功能。
//
// 本範例的三個角色分別是:
//   - Originator(發起人):Editor,擁有需要被保存與還原的狀態。
//   - Memento(備忘錄):memento,只有 Editor 能讀寫其內部欄位,
//     其他呼叫端只能把它當成不透明(opaque)的物件傳遞。
//   - Caretaker(管理者):由呼叫端(例如 example_test.go)自行維護一個
//     []memento 堆疊,負責保存與取出 memento,但不窺探其內容。
package memento

// memento 儲存 Editor 在某個時間點的狀態快照。
// 欄位刻意不外露,確保只有 Editor 能夠建立與還原狀態,
// 其餘外部程式碼(Caretaker)只能持有並傳遞這個型別,無法窺探或竄改其內容。
type memento struct {
	state string
}

// Editor 是本範例的 Originator(發起人),代表一個簡易的文字編輯器。
type Editor struct {
	content string
}

// Type 模擬使用者輸入文字,直接覆蓋目前的內容。
func (e *Editor) Type(s string) {
	e.content = s
}

// Content 回傳目前的編輯內容。
func (e *Editor) Content() string {
	return e.content
}

// Save 建立並回傳目前狀態的備忘錄快照,供 Caretaker 保存。
func (e *Editor) Save() memento {
	return memento{state: e.content}
}

// Restore 將 Editor 的狀態還原成指定備忘錄所記錄的內容。
func (e *Editor) Restore(m memento) {
	e.content = m.state
}
