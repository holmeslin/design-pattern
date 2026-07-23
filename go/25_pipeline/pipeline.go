// Package pipeline 示範以多階段 channel 串接資料的並發模式：
// 每個階段(stage)都是一個接收輸入 channel、回傳輸出 channel 的函式,
// 內部啟動一個 goroutine 處理資料,處理完畢後由「該階段自己」關閉輸出
// channel。下游只需持續從 channel 讀取,直到 channel 被關閉、range 迴圈
// 自然結束,不需要額外的同步機制。
package pipeline

// gen 是管線的第一階段(來源),將傳入的整數逐一送進回傳的 channel,
// 送完所有值後關閉 channel,通知下游沒有更多資料。
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square 是管線的中間階段,從 in 讀入每個整數、計算平方後送到輸出
// channel。當 in 被關閉且資料讀取完畢,range 迴圈結束,本階段隨即
// 關閉自己的輸出 channel,讓關閉狀態沿著管線往下游傳遞。
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

// sum 是管線的最後階段(匯集),讀取 in 的所有數值並加總,
// 直到 in 關閉才回傳最終結果。
func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}
