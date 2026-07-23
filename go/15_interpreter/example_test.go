package interpreter

import "fmt"

// ExampleExpr 手動組出運算式樹 (1 + 2) * 3,並求值。
func ExampleExpr() {
	var expr Expr = Mul{
		left:  Add{left: Num{n: 1}, right: Num{n: 2}},
		right: Num{n: 3},
	}

	fmt.Println(expr.Eval())
	// Output:
	// 9
}
