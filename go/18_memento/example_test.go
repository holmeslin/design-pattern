package memento_test

import (
	"fmt"

	"memento"
)

// ExampleEditor 示範 Editor(Originator)如何透過 Save/Restore
// 保存與還原自身狀態,達成類似復原(undo)的效果。
func ExampleEditor() {
	ed := &memento.Editor{}

	ed.Type("v1")
	saved := ed.Save() // Caretaker 保存 v1 當下的快照

	ed.Type("v2")
	fmt.Println(ed.Content())

	ed.Restore(saved) // 還原成 v1
	fmt.Println(ed.Content())

	// Output:
	// v2
	// v1
}
