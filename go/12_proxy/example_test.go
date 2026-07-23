package proxy_test

import (
	"fmt"

	"proxy"
)

// ExampleProxyImage 示範同一個 proxyImage 連續呼叫兩次 Load():
// 第一次會實際載入(委派給 realImage),第二次則直接回傳快取結果。
func ExampleProxyImage() {
	image := proxy.NewProxyImage("cat.png")

	fmt.Println(image.Load())
	fmt.Println(image.Load())

	// Output:
	// loading cat.png
	// cached cat.png
}
