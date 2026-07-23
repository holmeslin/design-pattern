// Package templatemethod 示範樣板方法模式(Template Method)。
//
// 傳統物件導向語言會用抽象類別定義固定流程,再由子類別覆寫其中的可變步驟。
// Go 沒有繼承,因此慣用做法是把可變步驟宣告成「函式欄位」(function field),
// 由呼叫端把不同的行為以函式值的形式注入,固定流程本身則由一個獨立函式負責執行。
package templatemethod

// Steps 定義樣板方法中可替換的各個步驟(hook)。
// 呼叫端透過填入不同的函式值來改變演算法的細節,而不需要改動固定流程本身。
type Steps struct {
	// Read 負責取得原始資料。
	Read func() []int
	// Process 負責處理/轉換資料。
	Process func([]int) []int
	// Write 負責將處理後的資料輸出成字串。
	Write func([]int) string
}

// Run 是樣板方法(Template Method):它定義了演算法的固定骨架,
// 依序呼叫 Read -> Process -> Write,並回傳 Write 的結果。
// 呼叫順序本身不可變,可變的只有 Steps 裡各個 hook 的實作內容。
func Run(s Steps) string {
	data := s.Read()
	data = s.Process(data)
	return s.Write(data)
}
