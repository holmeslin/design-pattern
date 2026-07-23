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

## 對照 Clean Architecture 分層

本範例為求聚焦,把所有角色平放在同一個 package;但每個角色在 Clean Architecture 裡各有歸屬。若要落成分層目錄,對應關係如下:

| 元件 | 分層 | 理由 |
|------|------|------|
| `Row` | **Entity / Domain** | 純資料結構,不依賴任何外層 |
| `Exporter`(介面) | **Use Case** | 抽象策略是「呼叫者」定義的 port,屬於內層 |
| `ExportService`(Context) | **Use Case** | 查表、委派、流程編排;不碰任何格式細節 |
| `CSVExporter`、`JSONExporter`、`ExportFunc`、`CSVExportFunc` | **Interface Adapters** | 具體策略,把 domain 資料轉成 CSV/JSON,依賴 `encoding/csv`、`encoding/json` |
| 策略註冊 `NewExportService(map{...})` | **Frameworks & Drivers** | 組裝點;決定有哪些具體策略,只發生在最外層 |

依賴方向永遠單向,由外指向內:

```
組裝(main)  ──►  具體策略(adapter)  ──►  Exporter + ExportService(usecase)  ──►  Row(domain)
```

關鍵:內層的 `ExportService` 只認 `Exporter` 介面,**不知道**外層有哪些具體策略存在;外層策略反過來實作內層宣告的介面。這正是 **依賴反轉原則(DIP)**——Strategy pattern 天然就是 Clean Architecture 邊界跨越的標準手法。要換格式、加格式,內層一行都不用動。

> 註:實務上 2~3 種 exporter 平放單一 package 反而更好維護;分層目錄的價值要在多 use case、需要強制依賴邊界時才浮現。

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
