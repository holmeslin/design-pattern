# Go 設計模式範例

GoF 23 種設計模式 + 4 種 Go 慣用並發模式,各自一個獨立範例。目錄依 **GoF 書本順序**編號(`01`–`23`),並發模式接在後面(`24`–`27`)。

## 慣例

每個模式一個獨立資料夾,彼此不互相 import:

```
NN_pattern/
├── go.mod           # module <pattern>;各自獨立,可單獨測試
├── README.md        # 一句話定義 + 角色表 + 怎麼跑
├── <pattern>.go     # 實作
└── example_test.go  # 可跑的 Example(結尾 // Output: 會被 go test 逐字比對)
```

驗證單一模式:

```sh
cd NN_pattern && go test ./...     # 並發模式(24–27)用 go test -race ./...
```

`21_strategy` 是最完整的示範(interface + function type 兩種寫法、Clean Architecture 分層對照、呼叫軌跡);其餘為聚焦單一重點的精簡版。

## 建立型 Creational

| # | 模式 | 範例題目 |
|---|------|---------|
| [01](01_abstract_factory) | Abstract Factory | UI 主題工廠,成套產出 Button + Checkbox(Dark / Light) |
| [02](02_builder) | Builder | 流暢式 SQL SELECT builder |
| [03](03_factory_method) | Factory Method | 依 channel 名產出 Notifier(email / sms) |
| [04](04_prototype) | Prototype | 深拷貝一份設定範本 |
| [05](05_singleton) | Singleton | `sync.Once` 的設定單例 |

## 結構型 Structural

| # | 模式 | 範例題目 |
|---|------|---------|
| [06](06_adapter) | Adapter | 把第三方 payment SDK 轉成自家介面 |
| [07](07_bridge) | Bridge | 遙控器(抽象)× 裝置(TV / Radio 實作) |
| [08](08_composite) | Composite | 檔案系統 File / Folder 遞迴算大小 |
| [09](09_decorator) | Decorator | `http.Handler` 式 middleware 疊 logging + auth |
| [10](10_facade) | Facade | 家庭劇院一鍵開關多子系統 |
| [11](11_flyweight) | Flyweight | 共享字元格式物件 |
| [12](12_proxy) | Proxy | 帶快取的圖片載入 proxy |

## 行為型 Behavioral

| # | 模式 | 範例題目 |
|---|------|---------|
| [13](13_chain_of_responsibility) | Chain of Responsibility | 請求驗證鏈 auth → ratelimit → body |
| [14](14_command) | Command | 可 undo 的帳戶操作 |
| [15](15_interpreter) | Interpreter | 簡單四則運算式直譯器 |
| [16](16_iterator) | Iterator | 自訂集合的 `iter.Seq`(Go 1.23 range-over-func) |
| [17](17_mediator) | Mediator | 聊天室,users 經 mediator 廣播 |
| [18](18_memento) | Memento | 編輯器狀態存檔 / 還原 |
| [19](19_observer) | Observer | 事件訂閱,溫度變化通知 observers |
| [20](20_state) | State | 訂單狀態機 |
| [21](21_strategy) | Strategy | 把資料匯出成不同格式(CSV / JSON) |
| [22](22_template_method) | Template Method | 資料處理骨架 讀 → 處理 → 寫,子步驟可換 |
| [23](23_visitor) | Visitor | 圖形面積 visitor(double dispatch) |

## 並發 Concurrency（Go 慣用）

| # | 模式 | 範例題目 |
|---|------|---------|
| [24](24_worker_pool) | Worker Pool | 固定數量 goroutine 消費 job channel |
| [25](25_pipeline) | Pipeline | 多階段 channel 串接 |
| [26](26_fan_in_out) | Fan-in / Fan-out | 一分多、多合一 |
| [27](27_pub_sub) | Pub/Sub | channel 廣播給多個訂閱者 |

## 全部跑一遍

各模式是獨立 module,一行掃過全部:

```sh
for d in */; do (cd "$d" && go test ./...); done
```
