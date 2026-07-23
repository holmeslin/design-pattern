// Package prototype 示範原型模式(Prototype Pattern)。
//
// 原型模式的核心是透過「複製既有物件」來建立新物件,
// 而不是每次都重新初始化。當物件的建構成本較高,
// 或是想保留某個已設定好的物件狀態作為範本時特別實用。
package prototype

// Config 是本範例中的原型(Prototype)角色。
// 它包含基本型別欄位(Name)以及參考型別欄位(Tags、Meta),
// 用來凸顯深拷貝(deep copy)與淺拷貝(shallow copy)的差異。
type Config struct {
	Name string
	Tags []string
	Meta map[string]string
}

// Clone 複製出一份與原物件完全獨立的 *Config。
//
// 深拷貝的重點:
//   - 基本型別(如 Name)在 Go 中值傳遞即可自動複製,不需特別處理。
//   - slice、map 屬於參考型別,若只是直接指派(如 clone.Tags = c.Tags),
//     兩者會共用同一塊底層陣列或雜湊表,修改其中一方會影響另一方。
//   - 因此必須重新配置(make)一塊新的記憶體空間,
//     再逐一複製每個元素,確保複製後的物件與原物件不共享任何底層資料結構。
func (c *Config) Clone() *Config {
	clone := &Config{
		Name: c.Name,
	}

	if c.Tags != nil {
		clone.Tags = make([]string, len(c.Tags))
		copy(clone.Tags, c.Tags)
	}

	if c.Meta != nil {
		clone.Meta = make(map[string]string, len(c.Meta))
		for k, v := range c.Meta {
			clone.Meta[k] = v
		}
	}

	return clone
}
