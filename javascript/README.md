# JavaScript 設計模式範例

GoF 23 種設計模式 + 5 種 JS 現代特有模式,各自一個獨立範例。`01`–`23` 的題目與編號對齊 [`go/`](../go)(依 GoF 書本順序);JS 現代特有模式接在後面(`24`–`28`),`go/` 的 `24`–`27` 是不同的模式,不對應。

## 慣例

每個模式一個獨立資料夾,彼此不互相 import:

```
NN_pattern/
├── README.md          # 一句話定義 + 角色表 + 怎麼跑
├── <pattern>.js       # 實作(只用標準 JS 與 Web 平台 API,不 import node:*、不碰 process)
└── <pattern>.test.js  # 測試(node:test + node:assert/strict)
```

`<pattern>` 是目錄名去掉編號前綴,例如 `01_abstract_factory/abstract_factory.js`。`21_strategy` 例外,有 `strategy_class.js` 與 `strategy_function.js` 兩種寫法。

`package.json` 零依賴,標明 `"type": "module"` 與 `"engines": { "node": ">=24" }`(`27_disposable` 的 `using` 需要 Node 24 以上)。

驗證單一模式(要給檔案 glob,不能只給目錄,Node 會把目錄當檔案而失敗):

```sh
cd javascript && node --test NN_pattern/*.test.js
```

`21_strategy` 是最完整的示範(class + 函式兩種寫法、Clean Architecture 分層對照);其餘為聚焦單一重點的精簡版。

## 建立型 Creational

| # | 模式 | 範例題目 | JS 慣用寫法 |
|---|------|---------|------------|
| [01](01_abstract_factory) | Abstract Factory | UI 主題工廠,成套產出 Button + Checkbox(Dark / Light) | 沒有 interface,工廠只是約定形狀的物件(duck typing) |
| [02](02_builder) | Builder | 流暢式 SQL SELECT builder | 每個設定方法回傳 `this`(method chaining),不需要 Director |
| [03](03_factory_method) | Factory Method | 依 channel 名產出 Notifier(email / sms) | 一個函式加一張物件查表,不需要 Creator 類別階層 |
| [04](04_prototype) | Prototype | 深拷貝一份設定範本 | 原生 `structuredClone` |
| [05](05_singleton) | Singleton | 全域設定物件,所有 import 看到同一份 | ES module 天生單例 |

## 結構型 Structural

| # | 模式 | 範例題目 | JS 慣用寫法 |
|---|------|---------|------------|
| [06](06_adapter) | Adapter | 把第三方 payment SDK(元)轉成自家介面(分) | 一般 class 包住既有類別,委派並換算 |
| [07](07_bridge) | Bridge | 遙控器(抽象)× 裝置(TV / Radio 實作) | 組合:抽象持有 device 欄位,duck typing 代替 Implementor 介面 |
| [08](08_composite) | Composite | 檔案系統 file / folder 遞迴算大小 | 小寫工廠函式 `file` / `folder`,避開全域 `File` 撞名 |
| [09](09_decorator) | Decorator | middleware 疊 logging + auth | 高階函式(函式是一級公民),不用 `@decorator` 語法 |
| [10](10_facade) | Facade | 家庭劇院一鍵開關多子系統 | module 不匯出子系統,呼叫端只能走 Facade |
| [11](11_flyweight) | Flyweight | 共享字元格式物件 | `Map` 快取 + `Object.freeze` 凍結共享物件 |
| [12](12_proxy) | Proxy | 帶快取的圖片載入 proxy | 原生 `Proxy` |

## 行為型 Behavioral

| # | 模式 | 範例題目 | JS 慣用寫法 |
|---|------|---------|------------|
| [13](13_chain_of_responsibility) | Chain of Responsibility | 請求驗證鏈 auth → ratelimit → body | `setNext` 回傳傳入的 handler,可連續串接 |
| [14](14_command) | Command | 可 undo 的帳戶存款操作 | `{ execute, undo }` 物件字面值(閉包),不需要 Command 類別 |
| [15](15_interpreter) | Interpreter | 整數加、乘運算式的語法樹直譯器 | 節點是工廠函式,方法叫 `interpret` 避免和全域 `eval` 混淆 |
| [16](16_iterator) | Iterator | `Ring` 集合的自訂走訪 | 原生 iterator protocol(`Symbol.iterator` + generator)與 Iterator helpers |
| [17](17_mediator) | Mediator | 聊天室,users 經 mediator 廣播 | 一般 class,Colleague 只持有 Mediator |
| [18](18_memento) | Memento | 編輯器內容存檔 / 還原 | `#private` 狀態 + `Object.freeze` 的凍結快照 |
| [19](19_observer) | Observer | 氣象站溫度變化通知多個顯示器 | 原生 `EventTarget` |
| [20](20_state) | State | 訂單狀態機 created → paid → shipped | 無資料的 state 用物件字面量,不用 class |
| [21](21_strategy) | Strategy | 把資料匯出成不同格式(CSV / JSON) | class 與函式兩種寫法並排,物件字面值 `{ export }` 當轉接器 |
| [22](22_template_method) | Template Method | 資料處理骨架 讀 → 處理 → 寫,子步驟可換 | `class extends` 覆寫 hook(GoF 原版做法) |
| [23](23_visitor) | Visitor | 圖形面積 visitor(double dispatch) | 無狀態 visitor 用物件字面量,duck typing 代替介面 |

## JS 現代特有 Modern JS

| # | 模式 | 範例題目 | JS 慣用寫法 |
|---|------|---------|------------|
| [24](24_signals) | Signals | 購物車 `total = price * qty` 的 computed 與 effect | 相依自動追蹤的最小實作(JS 目前無內建,TC39 提案中) |
| [25](25_async_iterator_pipeline) | Async Iterator Pipeline | 日誌管線:篩 `ERROR`、去前綴、取前 2 筆 | `async function*` 串接,pull-based 天生有 backpressure |
| [26](26_cancellation) | Cancellation | 慢速 API 搭配逾時與使用者主動取消 | 原生 `AbortController`、`AbortSignal.timeout()`、`AbortSignal.any()` |
| [27](27_disposable) | Disposable | `Connection`(同步)與 `FileLock`(非同步)自動釋放 | `Symbol.dispose` / `Symbol.asyncDispose` 搭配 `using` / `await using` |
| [28](28_concurrency_limiter) | Concurrency Limiter | `createLimiter(max)` 限制同時執行的非同步任務數 | 閉包內計數 + 等待佇列,`Promise.withResolvers()` |

## 全部跑一遍

在 `javascript/` 下,一行掃過全部 28 個測試檔:

```sh
cd javascript && node --test
```
