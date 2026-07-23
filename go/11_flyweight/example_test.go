package flyweight_test

import (
	"fmt"

	"flyweight"
)

// ExampleFactory_Get 展示同樣的 (font, size) 組合會從工廠取得同一個共享實體。
func ExampleFactory_Get() {
	factory := flyweight.NewFactory()

	a := factory.Get("Arial", 12)
	b := factory.Get("Arial", 12)

	fmt.Println("shared:", a == b)
	// Output:
	// shared: true
}
