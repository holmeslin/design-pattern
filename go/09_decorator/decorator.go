// Package decorator 示範裝飾者模式(Decorator Pattern)。
//
// 裝飾者模式讓我們可以在不修改原始物件的前提下,動態為它疊加新的行為。
// 每個裝飾者都實作與被裝飾對象相同的介面(Component),並在內部持有一個
// 該介面的實例(可能是原始物件,也可能是另一個裝飾者),藉由「包裝」的
// 方式一層層疊加職責。這正是 Go 標準庫 net/http 中 middleware 的形狀:
// http.Handler 包 http.Handler。
package decorator

// Handler 是本範例的 Component 角色,定義了被裝飾對象的統一介面。
type Handler interface {
	Handle(req string) string
}

// base 是 ConcreteComponent 角色,提供最基本、未被裝飾過的行為。
type base struct{}

// Handle 回傳固定的結果 "ok",作為最內層的實際處理邏輯。
func (base) Handle(req string) string {
	return "ok"
}

// handlerFunc 讓一般函式可以當作 Handler 使用,方便裝飾器內部組裝。
type handlerFunc func(req string) string

func (f handlerFunc) Handle(req string) string {
	return f(req)
}

// Logging 是 Decorator 角色之一,為傳入的 Handler 疊加日誌前綴 "[log] "。
// 它會先呼叫內層 h.Handle 取得結果,再於結果前方加上前綴。
func Logging(h Handler) Handler {
	return handlerFunc(func(req string) string {
		return "[log] " + h.Handle(req)
	})
}

// Auth 是 Decorator 角色之一,為傳入的 Handler 疊加驗證前綴 "[auth] "。
// 它會先呼叫內層 h.Handle 取得結果,再於結果前方加上前綴。
func Auth(h Handler) Handler {
	return handlerFunc(func(req string) string {
		return "[auth] " + h.Handle(req)
	})
}
