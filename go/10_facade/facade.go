// Package facade 示範門面模式(Facade Pattern)。
//
// 門面模式為一組複雜的子系統提供一個簡化的統一介面,
// 讓使用者不需要理解各子系統的細節,即可完成常見的操作。
package facade

// lights 代表燈光子系統(Subsystem)。
type lights struct{}

// dim 將燈光調暗。
func (lights) dim() string {
	return "lights dim"
}

// projector 代表投影機子系統(Subsystem)。
type projector struct{}

// on 開啟投影機。
func (projector) on() string {
	return "projector on"
}

// amplifier 代表擴大機子系統(Subsystem)。
type amplifier struct{}

// on 開啟擴大機。
func (amplifier) on() string {
	return "amplifier on"
}

// HomeTheater 是門面(Facade),整合 lights、projector、amplifier
// 三個子系統,提供一個簡單的方法讓使用者一次完成看電影前的準備工作。
type HomeTheater struct {
	lights    lights
	projector projector
	amplifier amplifier
}

// WatchMovie 依序啟動看電影所需的子系統,並回傳每個步驟的結果。
// 呼叫順序:燈光調暗 -> 開啟投影機 -> 開啟擴大機。
func (h *HomeTheater) WatchMovie() []string {
	return []string{
		h.lights.dim(),
		h.projector.on(),
		h.amplifier.on(),
	}
}
