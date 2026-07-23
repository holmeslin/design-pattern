package prototype_test

import (
	"fmt"

	"prototype"
)

// ExampleConfig_Clone 展示 Clone() 產生的是深拷貝:
// 對 clone 的 Tags 進行 append 之後,原本的 orig.Tags 不會受到影響。
func ExampleConfig_Clone() {
	orig := &prototype.Config{
		Name: "base",
		Tags: []string{"a"},
	}

	clone := orig.Clone()
	clone.Tags = append(clone.Tags, "b")

	fmt.Println("original:", orig.Tags)
	fmt.Println("clone:   ", clone.Tags)
	// Output:
	// original: [a]
	// clone:    [a b]
}
