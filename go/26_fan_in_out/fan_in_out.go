// Package faninout 示範並發模式中的「Fan-out / Fan-in」。
//
// Fan-out(一分多):同一條輸入 channel,由多個 goroutine 同時讀取並各自處理,
// 藉此把工作分散到多個 worker 上平行執行。
//
// Fan-in(多合一):多條輸出 channel,由一個 goroutine 統一收斂成單一 channel,
// 讓下游只需面對一個來源即可取得所有 worker 的結果。
package faninout

import "sync"

// gen 產生一條 channel,依序把 nums 送出後即關閉。
// 這是整條管線(pipeline)的來源。
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

// square 從 in 讀取數值,計算平方後送到回傳的 channel。
// 當 in 關閉且所有值都處理完畢後,回傳的 channel 也會關閉。
//
// Fan-out:呼叫端可以對同一個 in 啟動多個 square goroutine,
// 讓多個 worker 同時競爭讀取同一條輸入 channel,藉此平行處理工作。
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

// merge 是 Fan-in 的實作:把多條輸入 channel(cs)收斂成單一 channel 輸出。
// 內部為每條輸入 channel 各啟動一個 goroutine 負責轉送,
// 並用 sync.WaitGroup 追蹤所有轉送 goroutine 是否都已完成;
// 待全部完成後才關閉合併後的 channel,避免下游過早讀到「已關閉」的訊號。
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(c)
	}

	// 另開一個 goroutine 等待所有轉送工作結束後再關閉 out,
	// 避免阻塞呼叫者去做等待與關閉。
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// sum 讀取 in 內所有數值並回傳總和,直到 in 關閉為止。
// 由於加法滿足交換律,無論 merge 內部的併發交錯順序為何,
// 最終加總結果都是穩定且可重現的。
func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}
