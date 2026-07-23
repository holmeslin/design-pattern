package mediator

// ExampleChatRoom 展示中介者模式:Alice、Bob、Carol 依序加入同一個聊天室,
// 當 Alice 發送訊息時,ChatRoom 會依加入順序將訊息轉發給除了 Alice 之外的成員。
func ExampleChatRoom() {
	room := NewChatRoom()
	alice := NewUser("Alice", room)
	NewUser("Bob", room)
	NewUser("Carol", room)

	alice.Send("hi")

	// Output:
	// Bob received from Alice: hi
	// Carol received from Alice: hi
}
