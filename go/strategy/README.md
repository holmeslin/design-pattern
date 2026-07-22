# Strategy Pattern(策略模式)

把「會變的行為」抽成一個介面,讓演算法可以在**執行期**互換,而呼叫端完全不知道背後是哪一種實作。

本範例以「把資料匯出成不同格式(CSV / JSON)」當題目,對應簡報第 6～9 頁。

## 三個角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **Strategy**(介面) | `Exporter` | 定義會變的行為長什麼形狀 |
| **ConcreteStrategy** | `CSVExporter`、`JSONExporter`、`CSVExportFunc` | 一種演算法一個實作,細節只住在自己這裡 |
| **Context** | `ExportService` | 持有一組策略,對外一個入口;自己不做事,只查表 + 轉交 |

呼叫端只丟一個 `format` 字串進來,`ExportService` 查 map、委派出去。要多支援一種格式,就多註冊一個策略——**不用改 Context、不用改呼叫端**(開放封閉原則)。

## 兩種寫法

### A. Interface(`strategy_inteface.go`)

一種演算法 = 一個型別 + 一個 `Export` 方法。適合策略本身**有狀態**、或需要多個方法時。

```go
svc := NewExportService(map[string]Exporter{
    "csv":  &CSVExporter{},
    "json": &JSONExporter{},
})
svc.Export(ctx, w, rows, "csv")
```

### B. Function type(`strategy_func_type.go`)

策略只是「一個函式」時,不必為它造一個 struct。用一個函式型別 `ExportFunc` 掛上 `Export` 方法當轉接器,裸函式包一下就變成 `Exporter`,能和 struct 策略**並排住在同一張 map**:

```go
svc := NewExportService(map[string]Exporter{
    "csv": ExportFunc(CSVExportFunc), // 函式策略,透過轉接器變成 Exporter
})
```

這就是 stdlib 的常見手法(`http.HandlerFunc`、`sort.Slice` 的 less 函式)——**函式直接就是策略**,由消費者反覆回呼它。

## 怎麼選

- 策略無狀態、單一方法、寫起來像個 lambda → **function type**,少一個型別。
- 策略有狀態、需要多個方法、或想要有名字方便 mock → **interface**。

兩者對 `ExportService` 完全無差別,因為它只認 `Exporter` 介面。

## 執行

```sh
go test ./...   # 驗證兩種寫法輸出一致、未註冊格式會短路報錯
```

程式內保留了 `[1]~[5]` 的呼叫順序軌跡(走 stderr),方便對照「組裝 → 查表 → 委派 → 真正執行」這條路徑。

## 可跑的教材:`example_test.go`

`example_test.go` 是 Go 內建的 [Example 測試](https://pkg.go.dev/testing#hdr-Examples):函式名以 `Example` 開頭,結尾的 `// Output:` 註解就是預期的 **stdout**,`go test` 會實跑並逐字比對。所以它同時是「可執行的教材」和「會驗證的測試」。

```sh
go test ./...              # 連同 example 一起跑(Output 不符會 FAIL)
go test -run Example -v    # 只跑範例,-v 印出每個 Example 名稱與 PASS/FAIL
```

資料輸出走 stdout(被 Example 比對),`[1]~[5]` 呼叫軌跡走 stderr(不進比對、只給人看)。想看軌跡就加 `-v`:

```sh
go test -run ExampleExportService -v   # stderr 的 [1]~[5] 會一起顯示
```

三個範例:

| 函式 | 展示 |
|------|------|
| `ExampleExportService` | interface 版,csv / json / 未註冊(短路報錯)三條路徑 |
| `ExampleExportFunc` | function type 版,輸出與 interface 版 CSV 相同 |
| `ExampleExportFunc_sortSlice` | 標準庫式:函式直接就是策略(`sort.Slice` 的 less) |

範例也會出現在 `go doc` / pkg.go.dev 的文件裡:

```sh
go doc -all .   # 對應型別下方會列出這些 Example
```
