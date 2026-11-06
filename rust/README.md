# Rust 設計模式範例

GoF 23 種設計模式 + 5 種 Rust 特有模式,各自一個獨立 crate。`01`–`23` 的題目與編號對齊 [`go/`](../go)(依 GoF 書本順序);Rust 特有模式接在後面(`24`–`28`),`go/` 的 `24`–`27`(worker pool、pipeline、fan-in/out、pub/sub)是不同的模式,不對應。

## 慣例

整個 `rust/` 是一個 Cargo workspace(`members = ["[0-9][0-9]_*"]`),每個模式一個獨立 crate,彼此不互相依賴:

```
NN_pattern/
├── Cargo.toml     # 沿用 workspace 的版本、edition、lints
├── README.md      # 一句話定義 + 角色表 + Rust 慣用寫法 + 怎麼跑
└── src/lib.rs     # 實作 + #[cfg(test)] mod tests
```

package 名稱是目錄名去掉編號前綴,例如 `05_singleton` 的 package 叫 `singleton`。`21_strategy` 例外,`lib.rs` 之外還有 `strategy_trait.rs` 與 `strategy_closure.rs` 兩種寫法。

需要 Rust 1.96 以上(edition 2024、`resolver = "3"`)。所有 crate 都沒有 `[dependencies]`,只用標準函式庫,並以 workspace lint 禁止 `unsafe`。library 程式碼以 `Result` 回報錯誤,不靠 `unwrap` / `panic`(例外見 `28_fearless_concurrency` 的 README)。`24`、`25`、`28` 另有 `compile_fail` doctest,證明非法用法在編譯期就被擋下,所以 `cargo test` 也會跑 doctest。

驗證單一模式:

```sh
cd rust && cargo test -p <name>
```

`21_strategy` 是最完整的示範(trait + 函式 / 閉包兩種寫法、Clean Architecture 分層對照);其餘為聚焦單一重點的精簡版。

## 建立型 Creational

| # | 模式 | 範例題目 | Rust 慣用寫法 |
|---|------|---------|--------------|
| [01](01_abstract_factory) | Abstract Factory | UI 主題工廠,成套產出 Button + Checkbox(Dark / Light) | `trait GuiFactory` 回傳 `Box<dyn Button>`(執行期分派);編譯期就確定用哪一族時可改泛型 |
| [02](02_builder) | Builder | 鏈式 SQL SELECT builder(`select` / `from` / `where_clause` 再 `build`) | 消耗型 builder:每步 `self -> Self`,不需要 `&mut`,不會留下半成品;`where` 是關鍵字,所以叫 `where_clause` |
| [03](03_factory_method) | Factory Method | 依通道名("email" / "sms")產出 Notifier,未知通道回傳錯誤 | 一個 `match` 函式回傳 `Box<dyn Notifier>`,輸入不合法回 `Result<_, UnknownChannel>`,不 panic |
| [04](04_prototype) | Prototype | 複製設定物件 `Config`,改副本不影響原型 | `Clone` trait 就是 Prototype,`#[derive(Clone)]` 對 `String`、`Vec` 深拷貝;`Rc` / `Arc` 欄位只複製指標 |
| [05](05_singleton) | Singleton | 全域設定 `Config`(key/value),`instance()` 取唯一實例,可跨執行緒讀寫 | `static LazyLock`(對應 Go 的 `sync.Once`);可寫的全域狀態必須包 `Mutex`,中毒時用 `into_inner` 取回資料 |

## 結構型 Structural

| # | 模式 | 範例題目 | Rust 慣用寫法 |
|---|------|---------|--------------|
| [06](06_adapter) | Adapter | 把第三方支付 SDK(元,`f64`)轉成自家介面(分,`u64`) | 組合:Adapter 持有 Adaptee 欄位,實作 Target 並換算後委派;沒有繼承可用 |
| [07](07_bridge) | Bridge | 遙控器(抽象)× 裝置(TV / Radio 實作) | 泛型 `RemoteControl<D: Device>` 是靜態分派;要執行期才決定裝置,改 `Box<dyn Device>` |
| [08](08_composite) | Composite | 檔案系統 file / folder 遞迴算大小 | `enum Node` + `match`,編譯器檢查窮舉、不需要 `Box<dyn ...>`;種類要對外開放擴充才改 `trait` |
| [09](09_decorator) | Decorator | 請求處理器疊 logging + auth 前綴 | 泛型包裝 `Logging<H: Handler>`,`Logging<Auth<Base>>` 在編譯期是單一具體型別,零成本 |
| [10](10_facade) | Facade | 家庭劇院 `watch_movie()` 一次啟動燈光、投影機、擴大機 | 子系統型別不加 `pub`,由可見性規則把關,外部只能走 `HomeTheater` |
| [11](11_flyweight) | Flyweight | 共享文字格式(字型 + 大小) | `Rc<CharFormat>` + 工廠快取,可用 `Rc::ptr_eq` 驗證共享;`Rc` 只給 `&T`,共享物件天生不可變 |
| [12](12_proxy) | Proxy | 圖片延遲載入 proxy,第一次 `load()` 才建立真實物件 | `OnceCell` 延遲初始化,interior mutability 讓 `load` 只需 `&self`;多執行緒改用 `OnceLock` |

## 行為型 Behavioral

| # | 模式 | 範例題目 | Rust 慣用寫法 |
|---|------|---------|--------------|
| [13](13_chain_of_responsibility) | Chain of Responsibility | 請求驗證鏈 Auth → RateLimit → Body | 每環以 `Option<Box<dyn Handler>>` 擁有下一環,所有權沿鏈流動;拒絕以 `Result<(), ChainError>` 回傳 |
| [14](14_command) | Command | 可 undo 的帳戶存款(`Deposit` + `History`) | command 不持有 `&mut Account`,執行時才傳入,借用只活在單次呼叫裡;`History` 用 `Vec<Box<dyn Command>>` |
| [15](15_interpreter) | Interpreter | 整數加、乘運算式的語法樹,`(1 + 2) * 3` 求值為 `9` | `enum Expr` + `Box` 表示遞迴語法樹,`match` 求值;新增規則漏寫分支會編譯失敗 |
| [16](16_iterator) | Iterator | `Ring<T>` 集合的自訂走訪 | 內建 `Iterator` trait;為 `&Ring<T>` 實作 `IntoIterator`,`for x in &ring` 不消耗集合,`iter()` 回傳 `impl Iterator` |
| [17](17_mediator) | Mediator | 聊天室,成員經 `ChatRoom` 轉發訊息(跳過發送者) | mediator 擁有成員(`Vec<User>`),外部只拿 `UserId` handle,不需要 `Rc<RefCell>` + `Weak` |
| [18](18_memento) | Memento | 文字編輯器 `save` / `restore` | `Memento` 欄位私有,不透明由編譯器保證;`restore` 取走所有權且不實作 `Clone`,同一份快照不能還原兩次 |
| [19](19_observer) | Observer | 氣象站氣溫變化通知多個 observer,可退訂 | observer 就是閉包 `Box<dyn Fn(i32) -> String>`;退訂靠 `SubscriptionId`,不比較閉包 |
| [20](20_state) | State | 訂單狀態機 created → paid → shipped,已出貨再推進回傳錯誤 | 執行期狀態機:`Box<dyn State>`,非法轉換到執行期才收到 `Err`(編譯期版本見 `25_typestate`) |
| [21](21_strategy) | Strategy | 把資料匯出成不同格式(CSV / JSON) | trait 與函式 / 閉包兩種寫法並排,blanket impl 讓裸函式與閉包直接當策略;附 Clean Architecture 分層對照 |
| [22](22_template_method) | Template Method | 資料處理骨架 讀取 → 處理 → 輸出,各步驟可換 | trait 的 default method 就是樣板方法;漏實作任何步驟在編譯期就報錯 |
| [23](23_visitor) | Visitor | 圖形面積 visitor(double dispatch) | `accept` 加 `&dyn Visitor` 雙重分派;型別集合固定時更常用 `enum` + `match` 取代 Visitor |

## Rust 特有 Rust Idioms

| # | 模式 | 範例題目 | Rust 慣用寫法 |
|---|------|---------|--------------|
| [24](24_newtype) | Newtype | `Meters` / `Kilometers` 長度單位,以及替 `Vec<String>` 實作 `Display` 的 `Tags` | 單欄位 tuple struct,`#[repr(transparent)]` 零成本;單位混用編譯失敗,包成本地型別可繞過 orphan rule |
| [25](25_typestate) | Typestate | 訂單 created → paid → shipped,只能依序 `pay()` 再 `ship()` | 狀態編進型別 `Order<S>`(`PhantomData<S>`),轉換取走 `self`,非法順序編譯不過 |
| [26](26_raii_guard) | RAII Guard | `ConnectionGuard`:`open` 取得、被 drop 時自動釋放,用 log 驗證時機與順序 | `impl Drop`,釋放由所有權決定,不需要 Go 的 `defer` 或 JS 的 `using`;`Drop` 不能回傳錯誤 |
| [27](27_extension_trait) | Extension Trait | 替 `str` 加 `to_title_case()`,替所有 `Iterator<Item = i64>` 加 `sum_squares()` | 自訂 trait 加 blanket impl,呼叫端 `use` 該 trait 才看得到方法 |
| [28](28_fearless_concurrency) | Fearless Concurrency | 平方和拆給多個 thread 平行計算,以及 `Arc<Mutex<u64>>` 共用計數器 | `Send` / `Sync` 讓資料競爭在編譯期被擋;`thread::scope` 可直接借用呼叫端的切片,搭配 `mpsc::channel` 收集 |

## 全部跑一遍

在 `rust/` 下,一行跑完全部 28 個 crate 的測試(含 doctest):

```sh
cd rust && cargo test
```
