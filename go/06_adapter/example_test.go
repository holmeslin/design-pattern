package adapter

import "fmt"

// ExampleStripeAdapter 展示如何透過 StripeAdapter
// 讓不相容的 StripeSDK 滿足 PaymentProcessor 介面。
func ExampleStripeAdapter() {
	var p PaymentProcessor = StripeAdapter{}
	fmt.Println(p.Pay(1234))
	// Output:
	// charged $12.34
}
