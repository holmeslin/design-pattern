package strategy

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

// ============================================================
// 定義「會變的行為」長什麼形狀。
// ============================================================
type Row map[string]string

type Exporter interface {
	Export(ctx context.Context, w io.Writer, rows []Row) error
}

// ============================================================
// ConcreteStrategy 一種演算法，一個型別；各自的細節只住在自己的檔案裡。
// ============================================================
type CSVExporter struct{}

func (e *CSVExporter) Export(ctx context.Context, w io.Writer, rows []Row) error {
	fmt.Fprintln(os.Stderr, "    [5] CSVExporter.Export 真正執行:CSV 的細節只發生在這裡")
	cw := csv.NewWriter(w)
	if len(rows) == 0 {
		return nil
	}
	header := make([]string, 0, len(rows[0]))
	for k := range rows[0] {
		header = append(header, k)
	}
	sort.Strings(header) // 固定欄位順序
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

type JSONExporter struct{}

func (e *JSONExporter) Export(ctx context.Context, w io.Writer, rows []Row) error {
	fmt.Fprintln(os.Stderr, "    [5] JSONExporter.Export 真正執行:JSON 的細節只發生在這裡")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(rows)
}

// ============================================================
// Context 持有一組策略，對外只有一個入口；自己不做事，只查表 + 轉交。
// ============================================================
type ExportService struct {
	exporters map[string]Exporter // format → 策略
}

// NewExportService 組裝期：建立 Context 並把每種 format 對應的策略登記進 map。
func NewExportService(exporters map[string]Exporter) *ExportService {
	return &ExportService{exporters: exporters}
}

func (s *ExportService) Export(ctx context.Context, w io.Writer, rows []Row, format string) error {
	fmt.Fprintf(os.Stderr, "  [3] ExportService.Export 收到 format=%q,拿去 map 查表\n", format)
	exp, ok := s.exporters[format]
	if !ok {
		return fmt.Errorf("unsupported: %q", format) // 查不到就短路,不呼叫任何策略
	}
	fmt.Fprintf(os.Stderr, "  [4] 查到策略 %T,委派出去(Context 不碰細節)\n", exp)
	return exp.Export(ctx, w, rows)
}
