package singleton_test

import (
	"fmt"

	"singleton"
)

// ExampleInstance 展示多次呼叫 Instance() 都會取得同一個實例。
func ExampleInstance() {
	fmt.Println("same instance:", singleton.Instance() == singleton.Instance())
	// Output:
	// same instance: true
}
