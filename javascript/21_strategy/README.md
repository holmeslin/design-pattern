# Strategy(策略)

把「會變的行為」抽成一個形狀,讓演算法可以在**執行期**互換,而呼叫端完全不知道背後是哪一種實作。

本範例:把資料匯出成不同格式(CSV / JSON)。

## 三個角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Strategy** | 「有 `export(rows)` 的物件」這個形狀 | 定義會變的行為長什麼形狀 |
| **ConcreteStrategy** | `CSVExporter`、`JSONExporter`、`{ export: csvExport }` | 一種演算法一個實作,細節只住在自己這裡 |
| **Context** | `ExportService` | 持有一組策略,對外一個入口;自己不做事,只查表 + 轉交 |

呼叫端只丟一個 `format` 字串進來,`ExportService` 查 `Map`、委派出去;未註冊的格式直接 throw,不呼叫任何策略。要多支援一種格式,就多註冊一個策略,**不用改 Context、不用改呼叫端**(開放封閉原則)。

JS 沒有 interface,Strategy 只是一個約定的形狀(duck typing);Go 版用 `Exporter` 介面明確宣告。

## 對照 Clean Architecture 分層

本範例為求聚焦,把所有角色平放在同一個資料夾;但每個角色在 Clean Architecture 裡各有歸屬。若要落成分層目錄,對應關係如下:

| 元件 | 分層 | 理由 |
|------|------|------|
| `rows`(`[{ id, name }, ...]`) | **Entity / Domain** | 純資料,不依賴任何外層 |
| `{ export(rows) }` 這個形狀 | **Use Case** | 抽象策略是「呼叫者」定義的 port,屬於內層 |
| `ExportService`(Context) | **Use Case** | 查表、委派;不碰任何格式細節 |
| `CSVExporter`、`JSONExporter`、`csvExport` | **Interface Adapters** | 具體策略,把 domain 資料轉成 CSV / JSON |
| `new ExportService(new Map([...]))` | **Frameworks & Drivers** | 組裝點;決定有哪些具體策略,只發生在最外層 |

依賴方向永遠單向,由外指向內:

```
組裝(new ExportService(new Map([...])))  ──►  具體策略(adapter)  ──►  { export } 形狀 + ExportService(use case)  ──►  rows(domain)
```

關鍵:內層的 `ExportService` 只認 `export(rows)` 這個形狀,**不知道**外層有哪些具體策略存在;外層策略反過來符合內層宣告的形狀。這正是**依賴反轉原則(DIP)**,Strategy 是跨越 Clean Architecture 邊界的標準手法。要換格式、加格式,內層一行都不用動。

> 註:實務上 2~3 種 exporter 平放同一個資料夾反而更好維護;分層目錄的價值要在多 use case、需要強制依賴邊界時才浮現。

## 兩種寫法

### A. Class(`strategy_class.js`)

一種演算法 = 一個 class + 一個 `export` 方法。適合策略本身**有狀態**、或需要多個方法時。

```js
const svc = new ExportService(new Map([
  ['csv', new CSVExporter()],
  ['json', new JSONExporter()],
]));
svc.export(rows, 'csv');
```

### B. 函式(`strategy_function.js`)

策略只是「一個函式」時,不必為它造一個 class。把裸函式包成 `{ export: csvExport }`,**物件字面值就是轉接器**(對應 Go 的 `ExportFunc`),能和 class 策略**並排住在同一張 `Map`**:

```js
const svc = new ExportService(new Map([
  ['csv', { export: csvExport }], // 函式策略,用物件字面值變成 { export } 形狀
]));
```

這也是標準庫的常見手法:`[3, 1, 2].toSorted((a, b) => a - b)` 的比較函式**直接就是策略**,由消費者反覆回呼它。

## 怎麼選

- 策略有狀態、或需要多個方法 → **class**。
- 其他(無狀態、單一方法、寫起來像個 lambda) → **函式**,少一個 class。

兩者對 `ExportService` 完全無差別,因為它只認 `export(rows)` 這個形狀。

## 執行

```sh
cd javascript && node --test 21_strategy/*.test.js
```
