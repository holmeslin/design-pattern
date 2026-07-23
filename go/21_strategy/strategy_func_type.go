// Row / Exporter / ExportService 沿用 strategy_inteface.go,不重複定義——
// 這正是重點:Context 只認 Exporter 介面,不管背後是 struct 還是 function。
package strategy

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
)

// ============================================================
// ExportFunc 是一個「函式型別」;它掛了一個 Export 方法,
// 方法內容就是「回頭呼叫自己這個函式值」——這是 function type 的關鍵一跳。
// ============================================================

type ExportFunc func(ctx context.Context, w io.Writer, rows []Row) error

func (f ExportFunc) Export(ctx context.Context, w io.Writer, rows []Row) error {
	fmt.Fprintln(os.Stderr, "    [4] ExportFunc.Export(轉接器方法)被呼叫 → 轉去呼叫底層函式 f(...)")
	return f(ctx, w, rows)
}

// CSVExportFunc 一個「裸函式」策略:只要簽名符合 ExportFunc 即可,
// 不必是任何型別的方法。用 ExportFunc(CSVExportFunc) 包起來就變成 Exporter,
// 能和 CSVExporter 這種 struct 策略並排住在同一張 map 裡。
func CSVExportFunc(ctx context.Context, w io.Writer, rows []Row) error {
	fmt.Fprintln(os.Stderr, "    [5] CSVExportFunc 裸函式真正執行:寫 CSV 的細節只發生在這裡")
	cw := csv.NewWriter(w)
	if len(rows) == 0 {
		return nil
	}
	header := make([]string, 0, len(rows[0]))
	for k := range rows[0] {
		header = append(header, k)
	}
	sort.Strings(header)
	if err := cw.Write(header); err != nil {
		return err
	}
	for _, r := range rows {
		rec := make([]string, 0, len(header))
		for _, k := range header {
			rec = append(rec, r[k])
		}
		if err := cw.Write(rec); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}
