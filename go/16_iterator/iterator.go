// Package iterator 示範 Iterator（迭代器）模式。
//
// Iterator 模式的目的是在不暴露聚合物件（Aggregate）內部表示方式的前提下，
// 提供一種循序走訪其元素的方法。
//
// 自 Go 1.23 起，語言原生支援 range-over-func：任何型別為
// func(yield func(T) bool) 的函式都可以直接放進 for-range 使用，
// 標準庫將這種函式型別命名為 iter.Seq[T]。
// 因此在慣用（idiomatic）Go 中，不再需要自行定義 HasNext()/Next() 這類
// 傳統的 Iterator 介面，而是讓聚合物件回傳一個 iter.Seq[T]，
// 直接交給呼叫端用 for-range 走訪即可。
package iterator

import "iter"

// Ring 是本範例的聚合物件（Aggregate），內部以 slice 儲存元素。
// 呼叫端不需要知道內部是用 slice、link list 還是其他結構實作。
type Ring[T any] struct {
	items []T
}

// NewRing 建立一個包含指定元素的 Ring。
func NewRing[T any](items ...T) *Ring[T] {
	return &Ring[T]{items: items}
}

// All 回傳一個 iter.Seq[T]，也就是迭代器（Iterator）角色。
// 呼叫端可以直接用 for v := range r.All() 依序取得每一個元素；
// 若 yield 回傳 false，代表呼叫端要提前中止走訪，此時應立即停止。
func (r *Ring[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range r.items {
			if !yield(item) {
				return
			}
		}
	}
}
