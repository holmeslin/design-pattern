package state_test

import (
	"fmt"

	"state"
)

// ExampleOrder 展示訂單狀態機:created -> paid -> shipped。
func ExampleOrder() {
	o := state.NewOrder()
	fmt.Println(o.Advance())
	fmt.Println(o.Advance())
	// Output:
	// created -> paid
	// paid -> shipped
}
