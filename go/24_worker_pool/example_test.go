package workerpool_test

import (
	"fmt"

	"workerpool"
)

// ExampleRun 展示以 3 個 worker 處理 5 筆工作,
// 平方後加總: 1+4+9+16+25 = 55。
func ExampleRun() {
	fmt.Println(workerpool.Run([]int{1, 2, 3, 4, 5}, 3))
	// Output:
	// 55
}
