// Package state 示範 State(狀態)模式:
// 把「物件在不同狀態下有不同行為」這件事,拆成一組實作同一介面的具體狀態,
// 讓 Context 只需要委派給目前的狀態物件,不需要自己寫一堆 if/switch 判斷狀態。
package state

import "fmt"

// State 是狀態介面,每個具體狀態都要能回答「我是誰」以及「下一步要轉去哪個狀態」。
type State interface {
	// Next 根據目前的 Order,回傳轉移後的下一個狀態。
	Next(o *Order) State
	// Name 回傳這個狀態的名稱。
	Name() string
}

// createdState 是訂單剛建立、尚未付款的狀態。
type createdState struct{}

func (createdState) Name() string { return "created" }

func (createdState) Next(o *Order) State { return paidState{} }

// paidState 是訂單已付款、尚未出貨的狀態。
type paidState struct{}

func (paidState) Name() string { return "paid" }

func (paidState) Next(o *Order) State { return shippedState{} }

// shippedState 是訂單已出貨的狀態,屬於終態,Next 回傳自己表示不再轉移。
type shippedState struct{}

func (shippedState) Name() string { return "shipped" }

func (s shippedState) Next(o *Order) State { return s }

// Order 是 State 模式中的 Context,持有目前狀態並把行為委派給它。
type Order struct {
	state State
}

// NewOrder 建立一筆新訂單,初始狀態為 created。
func NewOrder() *Order {
	return &Order{state: createdState{}}
}

// Advance 把訂單推進到下一個狀態,回傳「舊狀態 -> 新狀態」的說明字串。
func (o *Order) Advance() string {
	old := o.state.Name()
	o.state = o.state.Next(o)
	return fmt.Sprintf("%s -> %s", old, o.state.Name())
}
