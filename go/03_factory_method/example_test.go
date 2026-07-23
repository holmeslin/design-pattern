package factorymethod_test

import (
	"fmt"

	"factorymethod"
)

// ExampleNewNotifier 展示如何透過工廠方法 NewNotifier 建立不同的
// Notifier,以及傳入未知 channel 時的錯誤處理。
func ExampleNewNotifier() {
	email, err := factorymethod.NewNotifier("email")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(email.Send("hello"))

	sms, err := factorymethod.NewNotifier("sms")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sms.Send("hello"))

	if _, err := factorymethod.NewNotifier("fax"); err != nil {
		fmt.Println(err)
	}

	// Output:
	// email: hello
	// sms: hello
	// unknown channel: "fax"
}
