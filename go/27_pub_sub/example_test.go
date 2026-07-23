package pubsub_test

import (
	"fmt"
	"sync"

	pubsub "pubsub"
)

// ExampleBroker 展示一個 Publisher 透過 Broker 將訊息廣播給兩個 Subscriber。
// 兩個訂閱者各自在獨立的 goroutine 中從自己的 channel 收到訊息後印出結果,
// 兩行輸出內容相同,順序無關,因此 Output 結果穩定。
func ExampleBroker() {
	broker := &pubsub.Broker{}
	defer broker.Close()

	sub1 := broker.Subscribe()
	sub2 := broker.Subscribe()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("sub got:", <-sub1)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("sub got:", <-sub2)
	}()

	broker.Publish("hello")
	wg.Wait()

	// Output:
	// sub got: hello
	// sub got: hello
}
