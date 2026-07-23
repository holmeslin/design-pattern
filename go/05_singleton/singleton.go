// Package singleton 示範 Singleton(單例)模式。
//
// Singleton 模式保證一個類別在整個程式生命週期中只有一個實例,
// 並提供一個全域存取點來取得該實例。
//
// 在 Go 中,慣用的實作方式是使用 sync.Once,而不是像其他語言
// 那樣手寫「雙重檢查鎖」(double-checked locking)。sync.Once
// 保證傳入 Do 的函式只會被執行一次,且對多個 goroutine 而言
// 是安全的:即使有多個 goroutine 同時呼叫 Instance,底層的
// 同步機制也能確保初始化邏輯只跑一次、且所有呼叫者都能看到
// 初始化完成後的結果。
package singleton

import "sync"

// registry 是本範例中的 Singleton 角色,代表整個程式共用的
// 一份資料存放區。
type registry struct {
	data map[string]string
}

var (
	once     sync.Once
	instance *registry
)

// Instance 是 Singleton 的存取點(access point)。
// 第一次呼叫時會透過 once.Do 建立唯一的 registry 實例,
// 之後每次呼叫都會回傳同一個實例,goroutine 安全。
func Instance() *registry {
	once.Do(func() {
		instance = &registry{
			data: make(map[string]string),
		}
	})
	return instance
}

// Set 將一組鍵值寫入單例的資料存放區。
func (r *registry) Set(key, value string) {
	r.data[key] = value
}

// Get 從單例的資料存放區讀取指定鍵的值。
func (r *registry) Get(key string) (string, bool) {
	value, ok := r.data[key]
	return value, ok
}
