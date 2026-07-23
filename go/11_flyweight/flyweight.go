// Package flyweight 示範享元模式(Flyweight Pattern)。
//
// 享元模式將物件的內部狀態(不隨情境變動、可共享)獨立出來,
// 交由工廠統一快取與配發,避免建立大量相同內容的物件,藉此節省記憶體。
package flyweight

import "fmt"

// Format 是享元(Flyweight)角色,代表可共享的文字格式(字型與大小)。
// 一旦建立即視為不可變(immutable),因此可安全地被多處共用同一個實體。
type Format struct {
	Font string
	Size int
}

// Factory 是享元工廠(FlyweightFactory)角色,
// 負責依 (Font, Size) 快取並配發共享的 *Format。
type Factory struct {
	cache map[string]*Format
}

// NewFactory 建立一個空的享元工廠。
func NewFactory() *Factory {
	return &Factory{cache: make(map[string]*Format)}
}

// Get 回傳對應 font 與 size 的共享 *Format。
// 若快取中已存在相同組合,直接回傳同一個指標;否則建立新實體並存入快取。
func (f *Factory) Get(font string, size int) *Format {
	key := fmt.Sprintf("%s-%d", font, size)

	if format, ok := f.cache[key]; ok {
		return format
	}

	format := &Format{Font: font, Size: size}
	f.cache[key] = format

	return format
}
