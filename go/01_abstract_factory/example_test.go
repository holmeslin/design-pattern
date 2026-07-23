package abstractfactory

import "fmt"

func ExampleGUIFactory() {
	fmt.Println(renderUI(DarkFactory{}))
	fmt.Println(renderUI(LightFactory{}))
	// Output:
	// [dark button] [dark checkbox]
	// [light button] [light checkbox]
}
