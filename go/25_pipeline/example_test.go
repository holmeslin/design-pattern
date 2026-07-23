package pipeline

import "fmt"

// Example 展示三階段管線:gen 產生來源資料、square 計算平方、
// sum 加總結果。1^2 + 2^2 + 3^2 = 1 + 4 + 9 = 14。
//
// 註:Go 的 Example 函式命名須對應套件內可匯出的識別字
// (Example<Type>、Example<Func> 等),本套件全為未匯出識別字,
// 故採用不綁定特定識別字的套件層級寫法 Example。
func Example() {
	fmt.Println(sum(square(gen(1, 2, 3))))
	// Output:
	// 14
}
