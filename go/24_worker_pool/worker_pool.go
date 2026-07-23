// Package workerpool 示範 Worker Pool(工作池)並發模式:
// 由固定數量的 worker goroutine 共同消費同一個 job channel,
// 將工作平均分攤出去,避免無限制地為每筆工作各開一個 goroutine。
package workerpool

import "sync"

// Run 啟動 workers 個 goroutine 處理 jobs,每個 worker 從同一個
// jobs channel 取值、計算平方後送進 results channel,最後回傳所有結果的總和。
//
// 重點:
//  1. 所有 worker 共享同一個唯讀的 jobs channel,Go runtime 保證每筆工作
//     只會被其中一個 worker 取走,因此可用固定數量的 goroutine 消化任意數量的工作。
//  2. 使用 sync.WaitGroup 追蹤所有 worker 是否收工;當最後一個 worker
//     呼叫 Done 後,由另一個 goroutine 負責 close(results),
//     讓主 goroutine 的讀取迴圈能正常結束,不會永久阻塞。
//  3. 加總(整數相加)具有交換律,不論 worker 執行順序或完成順序為何,
//     最終加總結果都相同,因此輸出是穩定、可重現的。
func Run(jobs []int, workers int) int {
	jobCh := make(chan int)
	resultCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(workers)

	// 啟動固定數量的 worker,各自從同一個 jobCh 取值運算。
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for j := range jobCh {
				resultCh <- j * j
			}
		}()
	}

	// 另開一個 goroutine 負責派送工作,並在派送完畢後關閉 jobCh,
	// 讓每個 worker 的 range 迴圈得以自然結束。
	go func() {
		defer close(jobCh)
		for _, j := range jobs {
			jobCh <- j
		}
	}()

	// 等所有 worker 收工後才 close(results),
	// 避免主 goroutine 讀取 resultCh 時提早看到 channel 關閉。
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	sum := 0
	for r := range resultCh {
		sum += r
	}
	return sum
}
