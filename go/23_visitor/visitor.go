// Package visitor 示範 Visitor(訪問者)模式。
//
// 核心想法:把「對一組物件結構做什麼操作」從物件本身抽離出去,
// 讓新增操作時不必修改既有的 Element 型別,只要新增一個 Visitor 實作即可。
// 代價是反過來——要新增一種 Element,就得修改 Visitor 介面與所有既有的實作。
package visitor

import "math"

// Visitor 定義「能夠拜訪哪些具體型別」的介面。
// 每個 VisitXxx 方法對應一種 ConcreteElement,回傳值以本範例的需求(面積)為例。
type Visitor interface {
	VisitCircle(Circle) float64
	VisitRectangle(Rectangle) float64
}

// Shape 是 Element 介面:只負責把自己「交給」傳進來的 Visitor 處理。
type Shape interface {
	Accept(Visitor) float64
}

// Circle 是 ConcreteElement:圓形,只保存自身資料(半徑)。
type Circle struct {
	r float64
}

// Accept 實現 double dispatch:先依賴實際型別(Circle)呼叫
// Visitor 對應的 VisitCircle,再由 Visitor 內部依實際邏輯處理。
func (c Circle) Accept(v Visitor) float64 {
	return v.VisitCircle(c)
}

// Rectangle 是 ConcreteElement:矩形,保存寬與高。
type Rectangle struct {
	w, h float64
}

// Accept 同樣把自己交給 Visitor,呼叫對應的 VisitRectangle。
func (r Rectangle) Accept(v Visitor) float64 {
	return v.VisitRectangle(r)
}

// AreaVisitor 是 ConcreteVisitor:針對每一種 Shape 計算面積。
type AreaVisitor struct{}

// VisitCircle 回傳圓形面積 πr²。
func (AreaVisitor) VisitCircle(c Circle) float64 {
	return math.Pi * c.r * c.r
}

// VisitRectangle 回傳矩形面積 w*h。
func (AreaVisitor) VisitRectangle(r Rectangle) float64 {
	return r.w * r.h
}
