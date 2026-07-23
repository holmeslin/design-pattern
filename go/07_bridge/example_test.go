package bridge_test

import (
	"fmt"

	"bridge"
)

// ExampleRemoteControl 示範同一支 RemoteControl(Abstraction)
// 如何透過 Device 介面(Implementor)操作不同的具體裝置。
func ExampleRemoteControl() {
	tv := bridge.NewRemoteControl(&bridge.TV{})
	tv.VolumeUp()
	fmt.Println(tv.Status())

	radio := bridge.NewRemoteControl(&bridge.Radio{})
	radio.VolumeUp()
	fmt.Println(radio.Status())

	// Output:
	// TV volume: 10
	// Radio volume: 10
}
