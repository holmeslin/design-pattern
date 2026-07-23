// Package observer 示範觀察者模式(Observer Pattern)。
//
// 觀察者模式定義了一種一對多的依賴關係:當 Subject(主題)狀態改變時,
// 會自動通知所有訂閱的 Observer(觀察者),讓它們能同步更新自身狀態。
package observer

import "fmt"

// Observer 是觀察者介面,任何想要接收溫度變化通知的型別都必須實作此介面。
type Observer interface {
	Notify(temp int)
}

// Station 是 Subject(主題),負責維護觀察者清單並在狀態改變時發出通知。
type Station struct {
	obs []Observer
}

// Subscribe 將一個 Observer 加入通知清單。
func (s *Station) Subscribe(o Observer) {
	s.obs = append(s.obs, o)
}

// SetTemp 設定新的溫度,並依訂閱順序通知所有已訂閱的 Observer。
func (s *Station) SetTemp(t int) {
	for _, o := range s.obs {
		o.Notify(t)
	}
}

// display 是具體的 Observer,收到通知時印出自己的名稱與溫度。
type display struct {
	name string
}

// Notify 實作 Observer 介面,印出 "名稱: 溫度"。
func (d *display) Notify(temp int) {
	fmt.Printf("%s: %d\n", d.name, temp)
}
