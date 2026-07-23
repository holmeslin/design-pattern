// Package builder 示範建造者模式(Builder Pattern)。
//
// 建造者模式將一個複雜物件的建構過程與其表示分離,
// 讓同樣的建構步驟可以組合出不同的結果。
// 本範例以「流暢式(Fluent Interface)SQL SELECT 語句建構器」為例:
// 呼叫端可以透過鏈式呼叫,一步步指定要查詢的欄位、資料表與條件,
// 最後再呼叫 Build() 產出最終的 SQL 字串。
package builder

import "strings"

// QueryBuilder 扮演「Builder(建造者)」角色。
// 它負責一步步收集建構 SQL 語句所需的各項資訊(欄位、資料表、條件),
// 並在最後由 Build() 組裝成最終的「Product(產品)」—— 也就是 SQL 字串。
type QueryBuilder struct {
	columns []string
	table   string
	where   string
}

// NewQueryBuilder 建立一個空的 QueryBuilder,作為建構的起點。
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// Select 設定要查詢的欄位,回傳自身以支援鏈式呼叫。
func (b *QueryBuilder) Select(cols ...string) *QueryBuilder {
	b.columns = cols
	return b
}

// From 設定要查詢的資料表,回傳自身以支援鏈式呼叫。
func (b *QueryBuilder) From(t string) *QueryBuilder {
	b.table = t
	return b
}

// Where 設定查詢條件,回傳自身以支援鏈式呼叫。
func (b *QueryBuilder) Where(cond string) *QueryBuilder {
	b.where = cond
	return b
}

// Build 依序組裝先前設定的欄位、資料表與條件,產出最終的「Product」—— SQL 字串。
// 若未呼叫 Where 設定條件,則省略 WHERE 子句。
func (b *QueryBuilder) Build() string {
	var sb strings.Builder

	sb.WriteString("SELECT ")
	sb.WriteString(strings.Join(b.columns, ", "))
	sb.WriteString(" FROM ")
	sb.WriteString(b.table)

	if b.where != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(b.where)
	}

	return sb.String()
}
