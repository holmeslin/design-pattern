package composite_test

import (
	"fmt"

	"composite"
)

// ExampleFolder_Size 展示組合模式如何以一致的方式處理 Leaf(File)
// 與 Composite(Folder),並示範 Folder 巢狀 Folder 時的遞迴加總行為。
func ExampleFolder_Size() {
	root := composite.NewFolder()
	root.Add(composite.NewFile(100))

	sub := composite.NewFolder()
	sub.Add(composite.NewFile(200))
	root.Add(sub)

	fmt.Println("total size:", root.Size())
	// Output:
	// total size: 300
}
