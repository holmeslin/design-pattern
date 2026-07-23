package visitor

import "fmt"

// ExampleAreaVisitor 示範用 AreaVisitor 拜訪 Rectangle 計算面積。
// Rectangle{2, 3} 面積為整數 6,避免浮點數格式化帶來的比對問題。
func ExampleAreaVisitor() {
	r := Rectangle{2, 3}
	area := r.Accept(AreaVisitor{})
	fmt.Printf("rectangle area: %g\n", area)
	// Output:
	// rectangle area: 6
}
