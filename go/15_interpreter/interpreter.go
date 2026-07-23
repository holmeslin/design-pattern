// Package interpreter 示範直譯器模式(Interpreter Pattern)。
//
// 直譯器模式為某個語言定義一套文法表示,並用一個直譯器來解釋這些句子。
// 本範例以簡單的四則運算(加法、乘法)運算式樹作為示範,
// 每個節點都實作 Expr 介面,呼叫 Eval 即可求值。
package interpreter

// Expr 是抽象表達式(AbstractExpression),定義所有運算式節點共同的求值行為。
type Expr interface {
	Eval() int
}

// Num 是終結表達式(Terminal Expression),代表運算式樹中的葉節點,即一個常數數值。
type Num struct {
	n int
}

// Eval 回傳常數本身的數值。
func (num Num) Eval() int {
	return num.n
}

// Add 是非終結表達式(NonTerminal Expression),代表加法運算,
// 由左右兩個子運算式組成。
type Add struct {
	left  Expr
	right Expr
}

// Eval 遞迴求出左右子運算式的值後相加。
func (add Add) Eval() int {
	return add.left.Eval() + add.right.Eval()
}

// Mul 是非終結表達式(NonTerminal Expression),代表乘法運算,
// 由左右兩個子運算式組成。
type Mul struct {
	left  Expr
	right Expr
}

// Eval 遞迴求出左右子運算式的值後相乘。
func (mul Mul) Eval() int {
	return mul.left.Eval() * mul.right.Eval()
}
