package command

import "fmt"

// ExampleHistory 示範透過 History(Invoker)執行 DepositCommand(ConcreteCommand),
// 並在需要時復原最後一次操作。
func ExampleHistory() {
	acc := &Account{}
	history := &History{}

	history.Run(NewDepositCommand(acc, 100))
	fmt.Println("balance:", acc.balance)

	history.UndoLast()
	fmt.Println("balance:", acc.balance)

	// Output:
	// balance: 100
	// balance: 0
}
