package observer

// ExampleStation 展示 Station 如何依訂閱順序通知所有觀察者。
func ExampleStation() {
	station := &Station{}
	station.Subscribe(&display{"phone"})
	station.Subscribe(&display{"window"})

	station.SetTemp(25)
	// Output:
	// phone: 25
	// window: 25
}
