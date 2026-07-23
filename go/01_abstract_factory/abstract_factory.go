// Package abstractfactory 示範 Abstract Factory:一個工廠介面成套產出「相關的一族產品」,
// 呼叫端只認工廠與產品介面,不知道具體是 Dark 還是 Light。
package abstractfactory

// ---- 產品介面:一族相關產品 ----
type Button interface{ Render() string }
type Checkbox interface{ Render() string }

// ---- 抽象工廠:成套產出一族產品 ----
type GUIFactory interface {
	Button() Button
	Checkbox() Checkbox
}

// ---- Dark 這一族 ----
type darkButton struct{}

func (darkButton) Render() string { return "[dark button]" }

type darkCheckbox struct{}

func (darkCheckbox) Render() string { return "[dark checkbox]" }

type DarkFactory struct{}

func (DarkFactory) Button() Button     { return darkButton{} }
func (DarkFactory) Checkbox() Checkbox { return darkCheckbox{} }

// ---- Light 這一族 ----
type lightButton struct{}

func (lightButton) Render() string { return "[light button]" }

type lightCheckbox struct{}

func (lightCheckbox) Render() string { return "[light checkbox]" }

type LightFactory struct{}

func (LightFactory) Button() Button     { return lightButton{} }
func (LightFactory) Checkbox() Checkbox { return lightCheckbox{} }

// renderUI 是呼叫端:只認 GUIFactory,拿到哪一族由外部決定。
func renderUI(f GUIFactory) string {
	return f.Button().Render() + " " + f.Checkbox().Render()
}
