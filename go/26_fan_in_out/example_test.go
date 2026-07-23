package faninout

import "fmt"

// Example_fanInOut 展示完整的 Fan-out / Fan-in 流程:
//
//  1. gen 產生來源 channel(1, 2, 3)。
//  2. Fan-out:啟動兩個 square goroutine,同時從同一條來源 channel 讀取並平方。
//  3. Fan-in:merge 把兩個 square 的輸出收斂成單一 channel。
//  4. sum 把合併後的所有數值加總,得到 1*1 + 2*2 + 3*3 = 14。
//
// 由於加總滿足交換律,即使 fan-out 的處理順序在每次執行時不同,
// 最終輸出的總和仍然穩定不變。
func Example_fanInOut() {
	in := gen(1, 2, 3)

	// fan-out:兩個 worker 同時消費同一條 in
	c1 := square(in)
	c2 := square(in)

	// fan-in:把兩個 worker 的輸出合併成一條 channel
	merged := merge(c1, c2)

	total := sum(merged)
	fmt.Println(total)
	// Output:
	// 14
}
