// Package factorymethod 示範工廠方法(Factory Method)模式。
//
// 工廠方法模式將「建立物件」的邏輯獨立成一個函式,呼叫端不需要知道
// 具體型別的建構細節,只需依賴共同的介面(Product)即可使用。
package factorymethod

import "fmt"

// Notifier 是本範例的 Product 角色,定義所有通知管道共同的行為。
type Notifier interface {
	// Send 傳送訊息,並回傳送出後的結果描述。
	Send(msg string) string
}

// emailNotifier 是 ConcreteProduct 之一,透過電子郵件傳送通知。
type emailNotifier struct{}

// Send 實作 Notifier 介面。
func (emailNotifier) Send(msg string) string {
	return "email: " + msg
}

// smsNotifier 是 ConcreteProduct 之一,透過簡訊傳送通知。
type smsNotifier struct{}

// Send 實作 Notifier 介面。
func (smsNotifier) Send(msg string) string {
	return "sms: " + msg
}

// NewNotifier 是 Creator 角色(工廠方法),依據 channel 建立對應的
// Notifier 實例。若 channel 不是已知的通知管道,則回傳 error。
func NewNotifier(channel string) (Notifier, error) {
	switch channel {
	case "email":
		return emailNotifier{}, nil
	case "sms":
		return smsNotifier{}, nil
	default:
		return nil, fmt.Errorf("unknown channel: %q", channel)
	}
}
