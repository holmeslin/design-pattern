// Package chain 示範責任鏈模式(Chain of Responsibility)。
//
// 責任鏈模式將多個處理者串成一條鏈,請求沿著鏈依序傳遞,
// 直到某個處理者處理完畢或選擇拒絕為止,呼叫端不需知道
// 究竟是哪一個處理者最終處理了請求。
package chain

import "errors"

// Request 代表一次待驗證的請求。
type Request struct {
	User string
	Body string
}

// Handler 是責任鏈中每個環節都必須實作的介面。
type Handler interface {
	// Handle 處理請求,驗證失敗時回傳 error。
	Handle(req Request) error
	// SetNext 設定鏈中的下一個處理者,並回傳該處理者以便串接。
	SetNext(next Handler) Handler
}

// Chain 是 Handler 的別名,代表一條已串接完成的責任鏈,
// 通常用來指稱鏈中最前端(head)的處理者。
type Chain = Handler

// base 提供 Handler 的共同行為:持有下一個處理者,
// 並提供轉發請求給下一環的輔助方法,讓具體處理者可以內嵌重用。
type base struct {
	next Handler
}

// SetNext 設定下一個處理者,並回傳傳入的處理者以便鏈式呼叫。
func (b *base) SetNext(next Handler) Handler {
	b.next = next
	return next
}

// callNext 將請求轉發給下一個處理者;若鏈已到尾端則視為通過。
func (b *base) callNext(req Request) error {
	if b.next == nil {
		return nil
	}
	return b.next.Handle(req)
}

// authHandler 驗證請求是否帶有使用者身分。
type authHandler struct {
	base
}

// NewAuthHandler 建立身分驗證環。
func NewAuthHandler() Handler {
	return &authHandler{}
}

func (h *authHandler) Handle(req Request) error {
	if req.User == "" {
		return errors.New("auth failed")
	}
	return h.callNext(req)
}

// rateLimitHandler 驗證請求是否超過流量限制。
type rateLimitHandler struct {
	base
}

// NewRateLimitHandler 建立流量限制環。
func NewRateLimitHandler() Handler {
	return &rateLimitHandler{}
}

func (h *rateLimitHandler) Handle(req Request) error {
	// 範例中未實作真正的計數邏輯,僅示意此環節的存在。
	return h.callNext(req)
}

// bodyHandler 驗證請求是否帶有內容。
type bodyHandler struct {
	base
}

// NewBodyHandler 建立內容驗證環。
func NewBodyHandler() Handler {
	return &bodyHandler{}
}

func (h *bodyHandler) Handle(req Request) error {
	if req.Body == "" {
		return errors.New("body is empty")
	}
	return h.callNext(req)
}
