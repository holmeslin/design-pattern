package facade_test

import (
	"fmt"

	"facade"
)

// ExampleHomeTheater_WatchMovie 示範透過門面 HomeTheater 一次啟動
// 所有子系統,而不需要個別操作 lights、projector、amplifier。
func ExampleHomeTheater_WatchMovie() {
	ht := &facade.HomeTheater{}
	for _, s := range ht.WatchMovie() {
		fmt.Println(s)
	}
	// Output:
	// lights dim
	// projector on
	// amplifier on
}
