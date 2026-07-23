// Package pubsub 示範發布訂閱(Publish/Subscribe)並發模式:
// Publisher 只需將訊息交給 Broker,不必知道有哪些訂閱者;
// Subscriber 各自透過獨立的 channel 接收訊息,彼此互不影響。
package pubsub

import "sync"

// Broker 是發布訂閱模式的訊息中樞。
// 它持有所有訂閱者的 channel,負責將訊息廣播給每一位訂閱者。
type Broker struct {
	mu   sync.Mutex
	subs []chan string
}

// Subscribe 註冊一個新的訂閱者,回傳唯讀 channel 供該訂閱者接收訊息。
// channel 帶有緩衝區,讓 Publish 在訂閱者尚未即時接收時也不會被阻塞。
func (b *Broker) Subscribe() <-chan string {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan string, 1)
	b.subs = append(b.subs, ch)
	return ch
}

// Publish 將訊息廣播給目前註冊的所有訂閱者。
func (b *Broker) Publish(msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subs {
		ch <- msg
	}
}

// Close 結束 Broker 的生命週期,關閉所有訂閱者的 channel。
// 呼叫 Close 之後不應再呼叫 Publish 或 Subscribe。
func (b *Broker) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subs {
		close(ch)
	}
	b.subs = nil
}
