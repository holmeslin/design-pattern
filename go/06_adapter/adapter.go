// Package adapter 示範 Adapter(轉接器)模式。
//
// 情境:我方系統只認得自家定義的 PaymentProcessor 介面,
// 但要串接的第三方 StripeSDK 介面不相容(單位不同、方法名稱不同)。
// 透過 StripeAdapter 把 StripeSDK 包裝成 PaymentProcessor,
// 讓不相容的介面得以協同運作,而不需修改任一方原本的程式碼。
package adapter

import "fmt"

// PaymentProcessor 是我方系統期望呼叫的目標介面(Target)。
// 呼叫端只依賴這個介面,不需要知道背後實際的付款服務為何。
type PaymentProcessor interface {
	// Pay 以「分」為單位收款,回傳處理結果描述。
	Pay(cents int) string
}

// StripeSDK 是第三方付款服務(Adaptee),介面與我方系統不相容:
// 它以「元」(float64)為單位,方法名稱也是 Charge 而非 Pay。
type StripeSDK struct{}

// Charge 以「元」為單位進行請款。
func (StripeSDK) Charge(dollars float64) string {
	return fmt.Sprintf("charged $%.2f", dollars)
}

// StripeAdapter 是轉接器(Adapter),把 StripeSDK 轉接成 PaymentProcessor。
// 它持有一個 StripeSDK 實例,負責在呼叫前做單位轉換(分轉元)。
type StripeAdapter struct {
	sdk StripeSDK
}

// Pay 實作 PaymentProcessor 介面,將 cents 換算成 dollars 後
// 轉呼叫 StripeSDK.Charge,完成介面轉接。
func (a StripeAdapter) Pay(cents int) string {
	dollars := float64(cents) / 100
	return a.sdk.Charge(dollars)
}
