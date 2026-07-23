// Package bridge 示範橋接模式(Bridge Pattern)。
//
// 橋接模式把「抽象」(這裡是 RemoteControl,呼叫端操作的介面)與
// 「實作」(這裡是 Device,實際被操作的裝置)拆成兩個獨立的繼承體系,
// 兩者透過組合(而非繼承)連接起來,讓抽象與實作可以各自獨立變化、
// 各自擴充,不會互相牽制。
package bridge

import "fmt"

// Device 是實作維度(Implementor)的介面,定義裝置必須提供的能力。
// 未來要新增任何裝置類型(例如喇叭、投影機),只要實作這個介面即可,
// 完全不需要更動 RemoteControl。
type Device interface {
	Name() string
	SetVolume(v int)
	Volume() int
}

// TV 是 Device 的其中一種具體實作(ConcreteImplementor)。
type TV struct {
	name   string
	volume int
}

// Name 回傳裝置名稱。
func (t *TV) Name() string {
	return "TV"
}

// SetVolume 設定電視音量。
func (t *TV) SetVolume(v int) {
	t.volume = v
}

// Volume 回傳目前的電視音量。
func (t *TV) Volume() int {
	return t.volume
}

// Radio 是 Device 的另一種具體實作(ConcreteImplementor)。
type Radio struct {
	name   string
	volume int
}

// Name 回傳裝置名稱。
func (r *Radio) Name() string {
	return "Radio"
}

// SetVolume 設定收音機音量。
func (r *Radio) SetVolume(v int) {
	r.volume = v
}

// Volume 回傳目前的收音機音量。
func (r *Radio) Volume() int {
	return r.volume
}

// RemoteControl 是抽象維度(Abstraction),對外提供操作裝置的入口。
// 它只依賴 Device 介面,不知道背後究竟是哪一種具體裝置,
// 這就是橋接模式的核心:抽象與實作透過介面連接,而非繼承。
type RemoteControl struct {
	dev Device
}

// NewRemoteControl 建立一個操作指定裝置的遙控器。
func NewRemoteControl(dev Device) *RemoteControl {
	return &RemoteControl{dev: dev}
}

// VolumeUp 把裝置音量調高 10。
func (r *RemoteControl) VolumeUp() {
	r.dev.SetVolume(r.dev.Volume() + 10)
}

// Status 回傳裝置目前的名稱與音量狀態描述。
func (r *RemoteControl) Status() string {
	return fmt.Sprintf("%s volume: %d", r.dev.Name(), r.dev.Volume())
}
