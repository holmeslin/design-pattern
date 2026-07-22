package strategy

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

var rows = []Row{
	{"id": "1", "name": "Alice"},
	{"id": "2", "name": "Bob"},
}

// 介面版:struct 策略。欄位固定字典序,所以 header 是 id,name。
func TestInterfaceStrategy(t *testing.T) {
	svc := NewExportService(map[string]Exporter{
		"csv":  &CSVExporter{},
		"json": &JSONExporter{},
	})

	var buf bytes.Buffer
	if err := svc.Export(context.Background(), &buf, rows, "csv"); err != nil {
		t.Fatalf("csv export: %v", err)
	}
	got := buf.String()
	want := "id,name\n1,Alice\n2,Bob\n"
	if got != want {
		t.Errorf("csv\n got: %q\nwant: %q", got, want)
	}

	buf.Reset()
	if err := svc.Export(context.Background(), &buf, rows, "json"); err != nil {
		t.Fatalf("json export: %v", err)
	}
	if !strings.Contains(buf.String(), `"name": "Alice"`) {
		t.Errorf("json missing Alice: %q", buf.String())
	}
}

// 沒註冊的 format 要短路回錯誤,而且不寫任何東西。
func TestUnsupportedFormat(t *testing.T) {
	svc := NewExportService(map[string]Exporter{"csv": &CSVExporter{}})
	var buf bytes.Buffer
	err := svc.Export(context.Background(), &buf, rows, "xml")
	if err == nil {
		t.Fatal("expected error for unsupported format")
	}
	if buf.Len() != 0 {
		t.Errorf("unsupported format should write nothing, got %q", buf.String())
	}
}

// function type 版:裸函式包成 ExportFunc,和 struct 策略同進一張 map,
// 對 Context 而言兩者無差別,輸出應與介面版 CSV 完全相同。
func TestFuncTypeStrategy(t *testing.T) {
	svc := NewExportService(map[string]Exporter{
		"csv": ExportFunc(CSVExportFunc),
	})
	var buf bytes.Buffer
	if err := svc.Export(context.Background(), &buf, rows, "csv"); err != nil {
		t.Fatalf("func csv export: %v", err)
	}
	if got, want := buf.String(), "id,name\n1,Alice\n2,Bob\n"; got != want {
		t.Errorf("func csv\n got: %q\nwant: %q", got, want)
	}
}
