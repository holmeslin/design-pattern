package builder_test

import (
	"fmt"

	"builder"
)

// ExampleQueryBuilder 展示如何透過鏈式呼叫使用 QueryBuilder
// 一步步組裝出一段 SELECT 語句。
func ExampleQueryBuilder() {
	sql := builder.NewQueryBuilder().
		Select("id", "name").
		From("users").
		Where("age > 18").
		Build()

	fmt.Println(sql)
	// Output:
	// SELECT id, name FROM users WHERE age > 18
}
