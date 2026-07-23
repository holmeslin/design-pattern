# Go 設計模式範例集 — Design Spec

**日期**:2026-07-23
**目標**:在 `go/` 下建立 GoF 23 種設計模式 + 4 種 Go 慣用並發模式,共 **27 個**獨立、可跑的教學範例,照 **GoF 書本順序**編號。

## 背景

已有 `go/21_strategy`(原 `go/14_strategy`,已重新命名)作為最豐富的示範樣本:含 interface + function type 兩種寫法、Clean Architecture 分層對照表、`[1]~[N]` 呼叫軌跡、`strategy_test.go` + `example_test.go`。本專案要把其餘 26 個模式補齊,風格一致但採**精簡版**。

## 決策(已與使用者確認)

| 項目 | 決定 |
|------|------|
| 範圍 | GoF 23 + 並發 4(Worker Pool、Pipeline、Fan-in/out、Pub/Sub) |
| 編號 | GoF 書本順序 01–23;並發 24–27 |
| strategy 位置 | 從 14 改為 **21**(目錄已改名 `go/21_strategy`) |
| 模組結構 | **每個 pattern 自己一個 `go.mod`**,可單獨 `go test` |
| 豐富度 | 精簡版:README + 實作 + `example_test.go`;strategy(21)維持完整版當旗艦 |
| Go 版本 | `go 1.25.7`(對齊現有) |

## 精簡版模板

每個 `go/NN_pattern/` 目錄含:

```
go/NN_pattern/
  go.mod            # module <pattern>; go 1.25.7;各自獨立
  README.md         # 繁中:一句話定義 + 角色表 + 怎麼跑;篇幅隨複雜度縮放
  <pattern>.go      # 實作(角色註解標清 Client / 抽象 / 具體)
  example_test.go   # 可跑的 Example(結尾 // Output: 逐字比對)——主要驗證
```

- `example_test.go` 是主要「可跑檢查」。若有錯誤/短路等 Example 蓋不到的分支,才另加一個小的 `NN_pattern_test.go`。
- README 結構沿用 strategy:一句話定義 → 角色對應表 → 執行方式;**不含**分層對照表與 `[1]~[N]` trace(那是 strategy 旗艦版專屬)。
- 命名採 idiomatic Go(介面用 `-er`、避免無謂 struct、能用函式就用函式)。

## 完整清單與範例題目

### 建立型 Creational(01–05)

| # | Pattern | 範例題目 | 關鍵 Go 手法 |
|---|---------|---------|-------------|
| 01 | Abstract Factory | UI 主題工廠,成套產出 Button + Checkbox(Dark/Light) | 一個工廠介面回傳一族相關產品 |
| 02 | Builder | 流暢式 SQL SELECT builder,分步組裝 | method chaining 回傳 `*Builder` |
| 03 | Factory Method | 依 channel 名產出 Notifier(email/sms) | 建構函式 + 介面回傳 |
| 04 | Prototype | 深拷貝一份設定範本 | `Clone()` 方法做 deep copy |
| 05 | Singleton | `sync.Once` 的設定單例 | `sync.Once` |

### 結構型 Structural(06–12)

| # | Pattern | 範例題目 | 關鍵 Go 手法 |
|---|---------|---------|-------------|
| 06 | Adapter | 把第三方 payment SDK 轉成自家介面 | wrapper struct 實作目標介面 |
| 07 | Bridge | 遙控器(抽象)× 裝置(TV/Radio 實作) | 抽象持有實作介面 |
| 08 | Composite | 檔案系統 File/Folder 遞迴算大小 | 共同介面 + 遞迴 |
| 09 | Decorator | `http.Handler` middleware 疊 logging + auth | 包裝同型別介面 |
| 10 | Facade | 家庭劇院一鍵開關多子系統 | 一個 struct 聚合多子系統 |
| 11 | Flyweight | 共享字元格式物件 | map 快取共享不可變物件 |
| 12 | Proxy | 帶快取的圖片載入 proxy | 同介面代理 + lazy/cache |

### 行為型 Behavioral(13–23)

| # | Pattern | 範例題目 | 關鍵 Go 手法 |
|---|---------|---------|-------------|
| 13 | Chain of Responsibility | 請求驗證鏈 auth→ratelimit→body | 每環持有 next |
| 14 | Command | 可 undo/redo 的帳戶操作 | Command 介面 + `Execute`/`Undo` |
| 15 | Interpreter | 簡單四則運算式直譯器 | 運算式樹 + `Eval` |
| 16 | Iterator | 自訂集合的 `iter.Seq`(Go 1.23 range-over-func) | `iter.Seq` |
| 17 | Mediator | 聊天室,users 經 mediator 廣播 | 中介物件轉發 |
| 18 | Memento | 編輯器狀態存檔/還原 | 快照 struct + 堆疊 |
| 19 | Observer | 事件訂閱,溫度變化通知 observers | 訂閱者清單 + 通知 |
| 20 | State | 訂單狀態機 | State 介面 + 轉移 |
| 21 | Strategy | (現有)CSV/JSON export — **完整版旗艦** | interface + function type |
| 22 | Template Method | 資料處理骨架 讀→處理→寫,子步驟可換 | embedding + hook 函式 |
| 23 | Visitor | 圖形 area/perimeter visitor | double dispatch via 介面 |

### 並發 Concurrency(24–27)

| # | Pattern | 範例題目 |
|---|---------|---------|
| 24 | Worker Pool | 固定數量 goroutine 消費 job channel |
| 25 | Pipeline | 多階段 channel 串接,每階段一組 goroutine |
| 26 | Fan-in / Fan-out | 一分多(fan-out)、多合一(fan-in) |
| 27 | Pub/Sub | channel 廣播給多個訂閱者 |

## 交付方式(分批 + 檢查點)

27 個一次做不完也不好審,分 4 批:

1. **建立型(01–05)** — 做完先讓使用者審一個,**確認精簡版格式合意再往下**(避免格式錯 27 次)。
2. **結構型(06–12)**
3. **行為型(13–23)** — 含確認 strategy 改名後仍正常。
4. **並發(24–27)**

每批完成的驗收:`cd go/NN_pattern && go test ./...` 全綠(含 Example 的 `// Output:` 比對)。

## 明確不做(YAGNI)

- 不做跨模組的 workspace(`go.work`)——各 pattern 獨立即可。
- 精簡版不寫分層對照表與 `[1]~[N]` trace(strategy 旗艦版專屬)。
- 不為每個函式寫完整測試套件;每個 pattern 一個 `example_test.go` 為主,分支不足才補小測試。
- strategy(21)與其餘深淺不一致是刻意的;使用者已確認接受。

## 驗收標準

- `go/` 下有 27 個編號目錄,命名符合上表。
- 每個目錄 `go test ./...` 通過。
- 每個 README 能讓讀者看懂「這是什麼模式、三個角色是誰、怎麼跑」。
- 原 strategy 內容不被破壞,僅目錄改名為 `21_strategy`。
