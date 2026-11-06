# Strategy(策略)

把「會變的行為」抽成一個形狀,讓演算法可以在**執行期**互換,而呼叫端完全不知道背後是哪一種實作。

本範例:把資料匯出成不同格式(CSV / JSON)。

## 三個角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Strategy** | `trait Exporter` | 定義會變的行為長什麼形狀 |
| **ConcreteStrategy** | `CsvExporter`、`JsonExporter`、`csv_export`(函式)、閉包 | 一種演算法一個實作,細節只住在自己這裡 |
| **Context** | `ExportService` | 持有一組策略,對外一個入口;自己不做事,只查表 + 轉交 |

呼叫端只丟一個 `format` 字串進來,`ExportService` 查 `HashMap`、委派出去;未註冊的格式直接回 `Err(ExportError::Unsupported(..))`,不呼叫任何策略。要多支援一種格式,就多 `register` 一個策略,**不用改 Context、不用改呼叫端**(開放封閉原則)。

## 對照 Clean Architecture 分層

本範例為求聚焦,把所有角色放在同一個 crate;但每個角色在 Clean Architecture 裡各有歸屬。若要落成分層,對應關係如下:

| 元件 | 分層 | 理由 |
|------|------|------|
| `Row`(`BTreeMap<String, String>`) | **Entity / Domain** | 純資料結構,不依賴任何外層 |
| `Exporter`(trait) | **Use Case** | 抽象策略是「呼叫者」定義的 port,屬於內層 |
| `ExportService`(Context) | **Use Case** | 查表、委派;不碰任何格式細節 |
| `CsvExporter`、`JsonExporter`、`csv_export` | **Interface Adapters** | 具體策略,把 domain 資料轉成 CSV / JSON(手寫,只用 std) |
| `register` 的呼叫處(`svc.register("csv", ..)`) | **Frameworks & Drivers** | 組裝點;決定有哪些具體策略,只發生在最外層 |

依賴方向永遠單向,由外指向內:

```
組裝(svc.register(..))  ──►  具體策略(adapter)  ──►  Exporter + ExportService(use case)  ──►  Row(domain)
```

關鍵:內層的 `ExportService` 只認 `Exporter` trait,**不知道**外層有哪些具體策略存在;外層策略反過來實作(或符合)內層宣告的 trait。這正是**依賴反轉原則(DIP)**,Strategy 是跨越 Clean Architecture 邊界的標準手法。要換格式、加格式,內層一行都不用動。`strategy_trait.rs`、`strategy_closure.rs` 都只從 `crate::` 引入 `Row`(以及 trait 版的 `Exporter`),也就是往內依賴;`lib.rs` 的 `mod` / `pub use` 只負責把兩個檔案接進 crate 並對外公開,`ExportService` 的程式碼沒有引用任何具體策略。

> 註:實務上 2~3 種 exporter 平放同一個 crate 反而更好維護;分層的價值要在多 use case、需要強制依賴邊界時才浮現。

## 兩種寫法

### A. Trait(`strategy_trait.rs`)

一種演算法 = 一個型別 + 一個 `export` 方法。適合策略本身**有狀態**、或需要多個方法時。

```rust
let mut svc = ExportService::default();
svc.register("csv", CsvExporter);
svc.register("json", JsonExporter);
svc.export(&rows, "csv"); // Ok("id,name\n1,Alice\n2,Bob")
```

### B. 函式與閉包(`strategy_closure.rs`)

策略只是「一個函式」時,不必為它造一個型別。`lib.rs` 裡有一個 blanket impl:

```rust
impl<F: Fn(&[Row]) -> String> Exporter for F { .. }
```

任何簽名符合 `Fn(&[Row]) -> String` 的函式或閉包都**自動**是 `Exporter`,裸函式直接註冊,能和 struct 策略**並排住在同一張 `HashMap`**:

```rust
svc.register("csv", csv_export); // 函式直接當策略,不用包轉接器
svc.register("count", |rows: &[Row]| rows.len().to_string()); // 閉包也行
```

這對應 Go 版的 `ExportFunc`:Go 要自己宣告函式型別、掛上方法,再用 `ExportFunc(CSVExportFunc)` 手動包一層;Rust 由 blanket impl 一次做完,呼叫端不用包。有兩個細節:

- 閉包要寫出參數型別(`|rows: &[Row]|`);編譯器無法透過 blanket impl 推斷閉包簽名,省略時(例如 `|rows| rows.len().to_string()`)會回報 `type annotations needed`。
- 要求的是 `Fn`,所以會修改捕獲變數的閉包(只實作 `FnMut`)不能直接註冊;要在策略裡記狀態,得用 `Cell` / `RefCell` 之類的內部可變性(測試裡計數呼叫次數的閉包就是這樣做)。

`register` 收的是 `impl Exporter + 'static`:註冊時以泛型接收(各自單態化),存進 `Box<dyn Exporter>` 之後,`export` 時走動態分派。`'static` 是因為 `Box<dyn Exporter>` 預設就是 `dyn Exporter + 'static`,策略不能借用呼叫端的區域變數。

這也是標準庫的常見手法:`v.sort_by_key(|&x| std::cmp::Reverse(x))` 的 key 閉包**直接就是策略**,由消費者(`sort_by_key`)反覆回呼它。

## 怎麼選

- 策略有狀態、或需要多個方法 → **trait**(struct)。
- 其他(無狀態、單一方法、寫起來像個 lambda) → **函式 / 閉包**,少一個型別。

兩者對 `ExportService` 完全無差別,因為它只認 `Exporter`。

## 執行

```sh
cd rust && cargo test -p strategy
```
