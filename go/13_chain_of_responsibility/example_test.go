package chain_test

import (
	"fmt"

	"chain"
)

// ExampleChain 展示如何串接 auth -> rateLimit -> body 三個處理環節,
// 並分別以合法與不合法的請求驗證責任鏈的行為。
func ExampleChain() {
	auth := chain.NewAuthHandler()
	rateLimit := chain.NewRateLimitHandler()
	body := chain.NewBodyHandler()

	auth.SetNext(rateLimit).SetNext(body)

	if err := auth.Handle(chain.Request{User: "alice", Body: "hello"}); err == nil {
		fmt.Println("ok")
	}

	if err := auth.Handle(chain.Request{User: "", Body: "hello"}); err != nil {
		fmt.Println(err)
	}

	// Output:
	// ok
	// auth failed
}
