package decorator

import "fmt"

// ExampleLogging 展示如何疊加多個裝飾者:Logging 包在最外層、Auth 次之、
// base 為最內層的實際處理邏輯。呼叫時由外而內(Logging -> Auth -> base),
// 結果字串則是由內而外組裝(base 的 "ok" 先產生,再依序被 Auth、Logging
// 加上各自的前綴)。
func ExampleLogging() {
	h := Logging(Auth(base{}))
	fmt.Println(h.Handle("req"))
	// Output:
	// [log] [auth] ok
}
