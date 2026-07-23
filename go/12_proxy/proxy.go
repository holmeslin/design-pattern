// Package proxy 示範代理模式(Proxy Pattern)。
//
// 代理模式為另一個物件提供一個替身或占位符,以控制對它的存取。
// 常見用途包含:延遲初始化(virtual proxy)、存取控制(protection proxy)、
// 快取(caching proxy)等。本範例實作快取代理:第一次載入圖片時委派給
// 真正的物件執行(模擬昂貴的載入動作),之後的請求則直接回傳快取結果。
package proxy

import "fmt"

// Image 是 Subject 角色,定義 RealSubject 與 Proxy 共同遵守的介面。
type Image interface {
	Load() string
}

// realImage 是 RealSubject 角色,代表實際執行昂貴載入動作的物件。
type realImage struct {
	name string
}

// Load 模擬一次昂貴的圖片載入動作。
func (r *realImage) Load() string {
	return fmt.Sprintf("loading %s", r.name)
}

// ProxyImage 是 Proxy 角色,持有 RealSubject 的參照,並在其前面加上快取邏輯。
// 第一次呼叫 Load 時才建立 realImage 並委派執行,結果會被快取;
// 之後的呼叫則直接回傳快取內容,不再重複進行昂貴的載入動作。
type ProxyImage struct {
	name  string
	real  *realImage
	cache string
}

// NewProxyImage 建立一個新的 ProxyImage。
func NewProxyImage(name string) *ProxyImage {
	return &ProxyImage{name: name}
}

// Load 實作 Image 介面,依情況委派給 RealSubject 或回傳快取結果。
func (p *ProxyImage) Load() string {
	if p.cache != "" {
		return fmt.Sprintf("cached %s", p.name)
	}

	p.real = &realImage{name: p.name}
	p.cache = p.real.Load()
	return p.cache
}
