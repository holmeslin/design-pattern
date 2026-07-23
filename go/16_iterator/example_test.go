package iterator

import "fmt"

// ExampleRing_All 示範如何用 for-range 走訪 Ring 內的元素。
func ExampleRing_All() {
	r := NewRing(10, 20, 30)
	for v := range r.All() {
		fmt.Println(v)
	}
	// Output:
	// 10
	// 20
	// 30
}
