package strategy

import (
	"context"
	"fmt"
	"os"
	"sort"
)

var demoRows = []Row{
	{"id": "1", "name": "Alice"},
	{"id": "2", "name": "Bob"},
}

// 寫法 A:interface 版。呼叫端只給 format 字串,不知道 CSVExporter 這型別存在。
// 資料輸出走 stdout;呼叫順序軌跡([1]~[5])走 stderr,不影響 Output 比對。
func ExampleExportService() {
	svc := NewExportService(map[string]Exporter{
		"csv":  &CSVExporter{},
		"json": &JSONExporter{},
	})
	ctx := context.Background()

	_ = svc.Export(ctx, os.Stdout, demoRows, "csv")
	_ = svc.Export(ctx, os.Stdout, demoRows, "json")

	// 沒註冊的格式:第 3 步查表就短路,不呼叫任何策略。
	if err := svc.Export(ctx, os.Stdout, demoRows, "xml"); err != nil {
		fmt.Println(err)
	}
	// Output:
	// id,name
	// 1,Alice
	// 2,Bob
	// [
	//   {
	//     "id": "1",
	//     "name": "Alice"
	//   },
	//   {
	//     "id": "2",
	//     "name": "Bob"
	//   }
	// ]
	// unsupported: "xml"
}

// 寫法 B:function type。裸函式包成 ExportFunc,和 struct 策略同進一張 map,
// 對 Context 完全無差別,輸出與 interface 版 CSV 相同。
func ExampleExportFunc() {
	svc := NewExportService(map[string]Exporter{
		"csv": ExportFunc(CSVExportFunc),
	})
	_ = svc.Export(context.Background(), os.Stdout, demoRows, "csv")
	// Output:
	// id,name
	// 1,Alice
	// 2,Bob
}

// 標準庫式:沒有 Context、沒有 map、沒有 interface——函式「直接就是」策略,
// 由 sort.Slice 這個消費者反覆回呼它。
func ExampleExportFunc_sortSlice() {
	nums := []int{3, 1, 2}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	fmt.Println(nums)
	// Output:
	// [1 2 3]
}
