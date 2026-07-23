// Package command 示範命令模式(Command Pattern)。
//
// 命令模式將「請求」封裝成物件,讓呼叫者(Invoker)不需要知道請求的
// 實際執行細節,並可支援排入佇列、記錄歷史、以及復原(Undo)等操作。
package command

// Command 是命令介面,定義所有具體命令都必須實作的行為。
type Command interface {
	// Execute 執行命令。
	Execute()
	// Undo 復原命令所造成的效果。
	Undo()
}

// Account 是接收者(Receiver),實際持有狀態並知道如何完成請求。
type Account struct {
	balance int
}

// DepositCommand 是具體命令(ConcreteCommand),
// 將「存款」這個請求連同其接收者與參數封裝起來。
type DepositCommand struct {
	acc    *Account
	amount int
}

// NewDepositCommand 建立一個存款命令。
func NewDepositCommand(acc *Account, amount int) *DepositCommand {
	return &DepositCommand{acc: acc, amount: amount}
}

// Execute 對帳戶執行存款動作。
func (c *DepositCommand) Execute() {
	c.acc.balance += c.amount
}

// Undo 復原存款動作,將金額扣回。
func (c *DepositCommand) Undo() {
	c.acc.balance -= c.amount
}

// History 是呼叫者(Invoker),負責觸發命令並記錄歷史,
// 以便之後可以復原最近一次執行的命令。
type History struct {
	cmds []Command
}

// Run 執行一個命令,並將其記錄到歷史中。
func (h *History) Run(c Command) {
	c.Execute()
	h.cmds = append(h.cmds, c)
}

// UndoLast 復原最後一個執行過的命令,並將其從歷史中移除。
func (h *History) UndoLast() {
	if len(h.cmds) == 0 {
		return
	}
	last := h.cmds[len(h.cmds)-1]
	last.Undo()
	h.cmds = h.cmds[:len(h.cmds)-1]
}
