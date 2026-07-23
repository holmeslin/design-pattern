package templatemethod

import "fmt"

// ExampleRun 展示樣板方法的固定流程:Read -> Process -> Write。
// 這裡的 Process 步驟過濾出偶數,示範可變步驟如何在不改變流程順序的前提下客製化。
func ExampleRun() {
	s := Steps{
		Read: func() []int {
			return []int{1, 2, 3, 4}
		},
		Process: func(data []int) []int {
			var result []int
			for _, v := range data {
				if v%2 == 0 {
					result = append(result, v)
				}
			}
			return result
		},
		Write: func(data []int) string {
			return fmt.Sprint(data)
		},
	}

	fmt.Println(Run(s))
	// Output:
	// [2 4]
}
