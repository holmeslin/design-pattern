# Go 設計模式範例集 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在 `go/` 下建立 GoF 23 種 + 4 種 Go 並發模式,共 27 個獨立可跑的教學範例,照 GoF 書本順序編號。

**Architecture:** 每個 pattern 是一個獨立目錄 `go/NN_pattern/`,含自己的 `go.mod`、README、實作、`example_test.go`。彼此不互相 import。`example_test.go` 的 `// Output:` 逐字比對是主要驗證。

**Tech Stack:** Go 1.25.7,只用標準庫。

## Global Constraints

- Go 版本:每個 `go.mod` 用 `go 1.25.7`。
- 每個 pattern 一個獨立 module:`module <pattern>`(小寫、無數字前綴,例如 `module abstractfactory`)。
- 只用標準庫,不加任何第三方相依。
- 每個 pattern 目錄:`go.mod` + `README.md`(繁中:一句話定義 + 角色表 + 怎麼跑)+ `<pattern>.go` + `example_test.go`。
- `example_test.go` 為主要驗證;錯誤/短路等 Example 蓋不到的分支才另加 `<pattern>_test.go`。
- 驗收:`cd go/NN_pattern && go test ./...` 全綠。
- 註解與 README 一律繁體中文;程式識別字保持英文。
- 精簡版**不寫** Clean Architecture 分層表與 `[1]~[N]` trace(那是 `21_strategy` 旗艦版專屬,不要動它)。
- 命名 idiomatic:介面用 `-er`;能用函式就別造 struct;固定 map 迭代順序時要排序,避免 Example Output 不穩定。

---

## Task 1(範本任務):建立型 01 Abstract Factory — 鎖死精簡版格式

這個任務寫出**完整程式碼**,作為其餘 26 個的格式範本。執行完請人工審一次格式再往下。

**Files:**
- Create: `go/01_abstract_factory/go.mod`
- Create: `go/01_abstract_factory/README.md`
- Create: `go/01_abstract_factory/abstract_factory.go`
- Create: `go/01_abstract_factory/example_test.go`

**Interfaces:**
- Produces: `GUIFactory` 介面(`Button() Button`、`Checkbox() Checkbox`);`Button`/`Checkbox` 介面(各一個 `Render() string`);`DarkFactory`、`LightFactory` 實作。

- [ ] **Step 1: 寫 go.mod**

```
module abstractfactory

go 1.25.7
```

- [ ] **Step 2: 寫 abstract_factory.go**

```go
// Package abstractfactory 示範 Abstract Factory:一個工廠介面成套產出「相關的一族產品」,
// 呼叫端只認工廠與產品介面,不知道具體是 Dark 還是 Light。
package abstractfactory

// ---- 產品介面:一族相關產品 ----
type Button interface{ Render() string }
type Checkbox interface{ Render() string }

// ---- 抽象工廠:成套產出一族產品 ----
type GUIFactory interface {
	Button() Button
	Checkbox() Checkbox
}

// ---- Dark 這一族 ----
type darkButton struct{}

func (darkButton) Render() string { return "[dark button]" }

type darkCheckbox struct{}

func (darkCheckbox) Render() string { return "[dark checkbox]" }

type DarkFactory struct{}

func (DarkFactory) Button() Button     { return darkButton{} }
func (DarkFactory) Checkbox() Checkbox { return darkCheckbox{} }

// ---- Light 這一族 ----
type lightButton struct{}

func (lightButton) Render() string { return "[light button]" }

type lightCheckbox struct{}

func (lightCheckbox) Render() string { return "[light checkbox]" }

type LightFactory struct{}

func (LightFactory) Button() Button     { return lightButton{} }
func (LightFactory) Checkbox() Checkbox { return lightCheckbox{} }

// renderUI 是呼叫端:只認 GUIFactory,拿到哪一族由外部決定。
func renderUI(f GUIFactory) string {
	return f.Button().Render() + " " + f.Checkbox().Render()
}
```

- [ ] **Step 3: 寫 example_test.go**

```go
package abstractfactory

import "fmt"

func ExampleGUIFactory() {
	fmt.Println(renderUI(DarkFactory{}))
	fmt.Println(renderUI(LightFactory{}))
	// Output:
	// [dark button] [dark checkbox]
	// [light button] [light checkbox]
}
```

- [ ] **Step 4: 寫 README.md**

````markdown
# Abstract Factory(抽象工廠)

一個工廠介面**成套產出「相關的一族產品」**,呼叫端只認工廠與產品介面,換一整族只換一個工廠。

本範例:UI 主題工廠,產出成套的 Button + Checkbox(Dark / Light)。

## 三個角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **AbstractFactory** | `GUIFactory` | 宣告「成套產出哪些產品」 |
| **ConcreteFactory** | `DarkFactory`、`LightFactory` | 產出同一族的具體產品 |
| **Product** | `Button`、`Checkbox` | 產品介面;具體產品細節住自己這裡 |

呼叫端 `renderUI` 只認 `GUIFactory`,要換整套主題就換一個工廠,不改呼叫端。

## 執行

```sh
go test ./...
```
````

- [ ] **Step 5: 驗證並 commit**

Run: `cd go/01_abstract_factory && go test ./...`
Expected: `ok  	abstractfactory`

```bash
git add go/01_abstract_factory
git commit -m "feat(go): add 01 abstract factory example"
```

**→ 檢查點:人工審此格式,確認後其餘任務照此範本產出。**

---

## Batch 1 剩餘:建立型 02–05

以下每個任務照 Task 1 範本(go.mod + README 角色表 + 實作 + example_test 的 // Output),module 名見括號。

### Task 2: 02 Builder(module `builder`)

**Files:** `go/02_builder/{go.mod,README.md,builder.go,example_test.go}`

**Design:** 流暢式 SQL SELECT builder。`type QueryBuilder struct` 有 `Select(cols ...string) *QueryBuilder`、`From(t string) *QueryBuilder`、`Where(cond string) *QueryBuilder`、`Build() string`,每個 setter 回傳 `*QueryBuilder` 以鏈接。`Build` 依序組出 SQL 字串。角色:Builder=`QueryBuilder`,Product=`string`(SQL)。

**Example Output(example_test.go `ExampleQueryBuilder`):**
```
SELECT id, name FROM users WHERE age > 18
```
(以 `NewQueryBuilder().Select("id","name").From("users").Where("age > 18").Build()` 產生)

**Verify:** `cd go/02_builder && go test ./...` → `ok`。Commit: `feat(go): add 02 builder example`。

### Task 3: 03 Factory Method(module `factorymethod`)

**Files:** `go/03_factory_method/{go.mod,README.md,factory_method.go,example_test.go}`

**Design:** `type Notifier interface { Send(msg string) string }`;`emailNotifier`、`smsNotifier` 實作。工廠函式 `NewNotifier(channel string) (Notifier, error)`,`channel` 為 `"email"`/`"sms"`,其餘回 error。角色:Product=`Notifier`,ConcreteProduct=email/sms,Creator=`NewNotifier`。

**Example Output(`ExampleNewNotifier`):**
```
email: hello
sms: hello
unknown channel: "fax"
```
(email/sms 各 `Send("hello")`,再對 `"fax"` 取得 error 並印出)

**Verify:** `cd go/03_factory_method && go test ./...`。Commit: `feat(go): add 03 factory method example`。

### Task 4: 04 Prototype(module `prototype`)

**Files:** `go/04_prototype/{go.mod,README.md,prototype.go,example_test.go}`

**Design:** `type Config struct { Name string; Tags []string; Meta map[string]string }`,方法 `Clone() *Config` 做 **deep copy**(slice 與 map 都要新配置,不共享底層)。角色:Prototype=`Config`,`Clone` 產生獨立副本。

**Example Output(`ExampleConfig_Clone`):** 示範改動 clone 的 slice/map 不影響原件,例如:
```
original: [a]
clone:    [a b]
```
(clone 後 append 一個 tag,印出兩者的 Tags 證明不共享)

**Verify:** `cd go/04_prototype && go test ./...`。Commit: `feat(go): add 04 prototype example`。

### Task 5: 05 Singleton(module `singleton`)

**Files:** `go/05_singleton/{go.mod,README.md,singleton.go,example_test.go}`

**Design:** 用 `sync.Once` 保證單例。`type registry struct { data map[string]string }`,套件層 `var once sync.Once; var instance *registry`,`Instance() *registry` 首次呼叫時初始化。角色:Singleton=`registry`,存取點=`Instance()`。

**Example Output(`ExampleInstance`):** 印出兩次 `Instance()` 是同一個指標:
```
same instance: true
```
(`Instance() == Instance()`)

**Verify:** `cd go/05_singleton && go test ./...`。Commit: `feat(go): add 05 singleton example`。

---

## Batch 2:結構型 06–12

### Task 6: 06 Adapter(module `adapter`)

**Files:** `go/06_adapter/{go.mod,README.md,adapter.go,example_test.go}`

**Design:** 自家目標介面 `type PaymentProcessor interface { Pay(cents int) string }`;第三方 `type StripeSDK struct{}` 有不相容方法 `Charge(dollars float64) string`。`type StripeAdapter struct { sdk StripeSDK }` 實作 `PaymentProcessor`,`Pay(cents)` 內把 cents 換算 dollars 再呼叫 `Charge`。角色:Target=`PaymentProcessor`,Adaptee=`StripeSDK`,Adapter=`StripeAdapter`。

**Example Output(`ExampleStripeAdapter`):**
```
charged $12.34
```
(`StripeAdapter{}.Pay(1234)`)

**Verify:** `cd go/06_adapter && go test ./...`。Commit: `feat(go): add 06 adapter example`。

### Task 7: 07 Bridge(module `bridge`)

**Files:** `go/07_bridge/{go.mod,README.md,bridge.go,example_test.go}`

**Design:** 實作維度 `type Device interface { Name() string; SetVolume(int); Volume() int }`,具體 `TV`、`Radio`。抽象維度 `type RemoteControl struct { dev Device }`,方法 `VolumeUp()`(把 device 音量 +10)、`Status() string`。抽象持有實作介面 → 兩維度各自擴充。角色:Abstraction=`RemoteControl`,Implementor=`Device`。

**Example Output(`ExampleRemoteControl`):**
```
TV volume: 10
Radio volume: 10
```
(對 TV 與 Radio 各 `NewRemoteControl(dev)` 後 `VolumeUp()` 再印 `Status()`)

**Verify:** `cd go/07_bridge && go test ./...`。Commit: `feat(go): add 07 bridge example`。

### Task 8: 08 Composite(module `composite`)

**Files:** `go/08_composite/{go.mod,README.md,composite.go,example_test.go}`

**Design:** `type Node interface { Size() int }`;`File{ size int }` 回傳自身大小;`Folder{ children []Node }` 的 `Size()` 遞迴加總。角色:Component=`Node`,Leaf=`File`,Composite=`Folder`。

**Example Output(`ExampleFolder_Size`):**
```
total size: 300
```
(組一個含兩個 File(100+200)的 Folder,可再巢狀一層)

**Verify:** `cd go/08_composite && go test ./...`。Commit: `feat(go): add 08 composite example`。

### Task 9: 09 Decorator(module `decorator`)

**Files:** `go/09_decorator/{go.mod,README.md,decorator.go,example_test.go}`

**Design:** 用 `http.Handler` 概念但不起 server:`type Handler interface { Handle(req string) string }`;基底 `base` 回 `"ok"`;裝飾器 `Logging(h Handler) Handler`、`Auth(h Handler) Handler`,各包一層在前後加字串。角色:Component=`Handler`,ConcreteComponent=`base`,Decorator=Logging/Auth(同型別包裝)。

**Example Output(`ExampleLogging`):**
```
[log] [auth] ok
```
(`Logging(Auth(base{}))` 對某 req 呼叫;log 在最外、auth 次之)

**Verify:** `cd go/09_decorator && go test ./...`。Commit: `feat(go): add 09 decorator example`。

### Task 10: 10 Facade(module `facade`)

**Files:** `go/10_facade/{go.mod,README.md,facade.go,example_test.go}`

**Design:** 子系統 `amplifier`、`projector`、`lights`,各有開關方法回字串。`type HomeTheater struct` 聚合三者,`WatchMovie() string` 依序呼叫並串接輸出。角色:Facade=`HomeTheater`,Subsystems=三個元件。

**Example Output(`ExampleHomeTheater_WatchMovie`):**
```
lights dim
projector on
amplifier on
```

**Verify:** `cd go/10_facade && go test ./...`。Commit: `feat(go): add 10 facade example`。

### Task 11: 11 Flyweight(module `flyweight`)

**Files:** `go/11_flyweight/{go.mod,README.md,flyweight.go,example_test.go}`

**Design:** 共享不可變的字元格式。`type Format struct { Font string; Size int }`(flyweight,不可變);`type factory struct { cache map[string]*Format }`,`Get(font string, size int) *Format` 對同 key 回同一指標。角色:Flyweight=`Format`,FlyweightFactory=`factory`。

**Example Output(`ExampleFactory_Get`):**
```
shared: true
```
(同樣 `Get("Arial",12)` 兩次得同一指標)

**Verify:** `cd go/11_flyweight && go test ./...`。Commit: `feat(go): add 11 flyweight example`。

### Task 12: 12 Proxy(module `proxy`)

**Files:** `go/12_proxy/{go.mod,README.md,proxy.go,example_test.go}`

**Design:** `type Image interface { Load() string }`;`realImage{ name string, loaded bool }` 的 `Load()` 首次回 `"loading X"`。`proxyImage{ real *realImage; cache string }` 實作快取:第一次委派 real,之後回 `"cached X"`。角色:Subject=`Image`,RealSubject=`realImage`,Proxy=`proxyImage`。

**Example Output(`ExampleProxyImage`):**
```
loading cat.png
cached cat.png
```
(同一 proxy 連呼 `Load()` 兩次)

**Verify:** `cd go/12_proxy && go test ./...`。Commit: `feat(go): add 12 proxy example`。

---

## Batch 3:行為型 13–23(21 為現有,只驗證)

### Task 13: 13 Chain of Responsibility(module `chain`)

**Files:** `go/13_chain_of_responsibility/{go.mod,README.md,chain.go,example_test.go}`

**Design:** `type Handler interface { Handle(req Request) error; SetNext(Handler) Handler }`,`Request{ User string; Body string }`。具體環 `authHandler`、`rateLimitHandler`、`bodyHandler`,各檢查一項,通過就傳給 next,失敗回 error 短路。用 base struct 嵌入處理 next 轉發。角色:Handler 介面 + 各具體環。

**Example Output(`ExampleChain`):** 一個合法 request 通過全部;一個缺 User 的在第一環短路:
```
ok
auth failed
```

**Verify:** `cd go/13_chain_of_responsibility && go test ./...`。Commit: `feat(go): add 13 chain of responsibility example`。

### Task 14: 14 Command(module `command`)

**Files:** `go/14_command/{go.mod,README.md,command.go,example_test.go}`

**Design:** `type Command interface { Execute(); Undo() }`;`Account{ balance int }`;`DepositCommand{ acc *Account; amount int }` 的 Execute 加、Undo 減。Invoker `type History struct { cmds []Command }`,`Run(c Command)` 執行並記錄,`UndoLast()` 還原。角色:Command 介面、ConcreteCommand=Deposit、Receiver=Account、Invoker=History。

**Example Output(`ExampleHistory`):**
```
balance: 100
balance: 0
```
(Deposit 100 後印,再 UndoLast 後印)

**Verify:** `cd go/14_command && go test ./...`。Commit: `feat(go): add 14 command example`。

### Task 15: 15 Interpreter(module `interpreter`)

**Files:** `go/15_interpreter/{go.mod,README.md,interpreter.go,example_test.go}`

**Design:** 運算式樹。`type Expr interface { Eval() int }`;`Num{ n int }`;`Add{ left, right Expr }`;`Mul{ left, right Expr }`。角色:AbstractExpression=`Expr`,Terminal=`Num`,NonTerminal=`Add`/`Mul`。(可手組樹,不必寫 parser。)

**Example Output(`ExampleExpr`):** 對 `(1 + 2) * 3`:
```
9
```

**Verify:** `cd go/15_interpreter && go test ./...`。Commit: `feat(go): add 15 interpreter example`。

### Task 16: 16 Iterator(module `iterator`)

**Files:** `go/16_iterator/{go.mod,README.md,iterator.go,example_test.go}`

**Design:** 用 Go 1.23 `iter.Seq`(range-over-func)。`type Ring[T any] struct { items []T }`,方法 `All() iter.Seq[T]` 回傳可被 `for v := range r.All()` 迭代的函式。角色:Aggregate=`Ring`,Iterator=`iter.Seq`(語言原生)。README 點出這是 Go 慣用寫法,取代傳統 Iterator 介面。

**Example Output(`ExampleRing_All`):**
```
10
20
30
```
(Ring 裝 10,20,30,range 印出)

**Verify:** `cd go/16_iterator && go test ./...`。Commit: `feat(go): add 16 iterator example`。

### Task 17: 17 Mediator(module `mediator`)

**Files:** `go/17_mediator/{go.mod,README.md,mediator.go,example_test.go}`

**Design:** 聊天室。`type Mediator interface { Broadcast(from string, msg string) }`;`ChatRoom` 持有 `[]*User`,`Broadcast` 轉發給除 from 外的所有 user(收集成字串以利 Example 穩定,依加入順序)。`User{ name string; room Mediator }` 的 `Send(msg)` 呼叫 `room.Broadcast`。角色:Mediator=`ChatRoom`,Colleague=`User`。

**Example Output(`ExampleChatRoom`):**
```
Bob received from Alice: hi
Carol received from Alice: hi
```
(Alice/Bob/Carol 入房,Alice 送 "hi")

**Verify:** `cd go/17_mediator && go test ./...`。Commit: `feat(go): add 17 mediator example`。

### Task 18: 18 Memento(module `memento`)

**Files:** `go/18_memento/{go.mod,README.md,memento.go,example_test.go}`

**Design:** `Editor{ content string }`,`Save() memento` 回快照,`Restore(m memento)` 還原。`memento{ state string }` 欄位不外露(同 package 即可示範封裝)。Caretaker 用一個 `[]memento` 堆疊。角色:Originator=`Editor`,Memento=`memento`,Caretaker=堆疊。

**Example Output(`ExampleEditor`):**
```
v2
v1
```
(打字到 "v1" 存檔、改成 "v2" 印出,Restore 後印回 "v1")— 注意順序:先印當前 "v2"?依實作決定,Example 的 Output 要與程式碼一致。實作者:先 `Type("v1"); Save(); Type("v2"); Println(content)` → v2,再 `Restore(第一個快照); Println(content)` → v1。

**Verify:** `cd go/18_memento && go test ./...`。Commit: `feat(go): add 18 memento example`。

### Task 19: 19 Observer(module `observer`)

**Files:** `go/19_observer/{go.mod,README.md,observer.go,example_test.go}`

**Design:** `type Observer interface { Notify(temp int) }`;`Station struct { obs []Observer }`,`Subscribe(o Observer)`、`SetTemp(t int)` 通知所有 observer(依訂閱順序)。具體 observer `display{ name string }` 收到就印。角色:Subject=`Station`,Observer 介面 + 具體 display。

**Example Output(`ExampleStation`):**
```
phone: 25
window: 25
```
(phone、window 訂閱後 `SetTemp(25)`)

**Verify:** `cd go/19_observer && go test ./...`。Commit: `feat(go): add 19 observer example`。

### Task 20: 20 State(module `state`)

**Files:** `go/20_state/{go.mod,README.md,state.go,example_test.go}`

**Design:** 訂單狀態機。`type State interface { Next(*Order) string; Name() string }`;狀態 `created`、`paid`、`shipped`。`Order{ state State }`,`Advance()` 呼叫目前狀態的 `Next` 轉移到下一個並回傳新狀態名。角色:Context=`Order`,State 介面 + 各具體狀態。

**Example Output(`ExampleOrder`):**
```
created -> paid
paid -> shipped
```
(從 created 起 `Advance()` 兩次)

**Verify:** `cd go/20_state && go test ./...`。Commit: `feat(go): add 20 state example`。

### Task 21: 21 Strategy(現有,只驗證不重寫)

**Files:** 無新增。目錄 `go/21_strategy/` 已存在(由 `14_strategy` 改名)。

- [ ] **Step 1:** Run `cd go/21_strategy && go test ./...` → Expected: `ok`(含三個 Example)。
- [ ] **Step 2:** 開 `go/21_strategy/README.md`,確認標題與內文沒有殘留 "14" 之類舊編號;若有提到目錄編號則更新為 21。若無改動則跳過。
- [ ] **Step 3:** 若有改動則 commit:`docs(go): fix strategy dir number references`;否則不 commit。

### Task 22: 22 Template Method(module `templatemethod`)

**Files:** `go/22_template_method/{go.mod,README.md,template_method.go,example_test.go}`

**Design:** Go 無繼承,用「骨架函式 + 可替換 hook」。`type Steps struct { Read func() []int; Process func([]int) []int; Write func([]int) string }`;`Run(s Steps) string` 固定呼叫順序 Read→Process→Write。提供預設 hook,可替換 Process(例如過濾偶數 vs 全部)。角色:Template=`Run`,可變步驟=hook 函式。README 說明 Go 慣用「函式欄位」取代抽象方法覆寫。

**Example Output(`ExampleRun`):**
```
[2 4]
```
(Read 回 1..4,Process 過濾偶數,Write 印 slice)

**Verify:** `cd go/22_template_method && go test ./...`。Commit: `feat(go): add 22 template method example`。

### Task 23: 23 Visitor(module `visitor`)

**Files:** `go/23_visitor/{go.mod,README.md,visitor.go,example_test.go}`

**Design:** `type Shape interface { Accept(Visitor) float64 }`;`Circle{ r float64 }`、`Rectangle{ w, h float64 }`。`type Visitor interface { VisitCircle(Circle) float64; VisitRectangle(Rectangle) float64 }`;具體 `AreaVisitor` 算面積。`Accept` 內回呼對應 Visit(double dispatch)。角色:Element=`Shape`,ConcreteElement=Circle/Rectangle,Visitor 介面 + AreaVisitor。

**Example Output(`ExampleAreaVisitor`):** 用固定值避免浮點格式問題,例如 Circle{r:1} 用 `math.Pi`,Rectangle{2,3}:
```
rectangle area: 6
```
(至少印矩形面積 6;圓面積若印,用 `%.2f` 固定格式)

**Verify:** `cd go/23_visitor && go test ./...`。Commit: `feat(go): add 23 visitor example`。

---

## Batch 4:並發 24–27

### Task 24: 24 Worker Pool(module `workerpool`)

**Files:** `go/24_worker_pool/{go.mod,README.md,worker_pool.go,example_test.go}`

**Design:** `Run(jobs []int, workers int) int`:啟 `workers` 個 goroutine 從 `jobs` channel 取值,各自平方後送 `results` channel,主 goroutine 加總回傳。用 `sync.WaitGroup` 等 workers 收工後 close results。角色:固定數量 worker 消費同一 job channel。Example 用可加總的結果避免順序問題。

**Example Output(`ExampleRun`):**
```
55
```
(jobs = 1..5 各平方 = 1+4+9+16+25 = 55,3 個 worker)

**Verify:** `cd go/24_worker_pool && go test -race ./...` → `ok`。Commit: `feat(go): add 24 worker pool example`。

### Task 25: 25 Pipeline(module `pipeline`)

**Files:** `go/25_pipeline/{go.mod,README.md,pipeline.go,example_test.go}`

**Design:** 三階段 channel:`gen(nums...)` 產出 → `square(in)` 平方 → `sum(in)` 加總。每階段回傳 `<-chan int`,用 goroutine + close。角色:多階段 channel 串接。

**Example Output(`ExamplePipeline`):**
```
14
```
(gen 1,2,3 → square → sum = 1+4+9 = 14)

**Verify:** `cd go/25_pipeline && go test -race ./...`。Commit: `feat(go): add 25 pipeline example`。

### Task 26: 26 Fan-in / Fan-out(module `faninout`)

**Files:** `go/26_fan_in_out/{go.mod,README.md,fan_in_out.go,example_test.go}`

**Design:** `gen` 產出一條 channel;fan-out:啟 N 個 `square` goroutine 都讀同一輸入;fan-in:`merge(chans...)` 用 `sync.WaitGroup` 把多條輸出併成一條。主程式加總。角色:一分多、多合一。

**Example Output(`ExampleFanInOut`):**
```
14
```
(1,2,3 平方後 fan-in 加總 = 14)

**Verify:** `cd go/26_fan_in_out && go test -race ./...`。Commit: `feat(go): add 26 fan-in fan-out example`。

### Task 27: 27 Pub/Sub(module `pubsub`)

**Files:** `go/27_pub_sub/{go.mod,README.md,pub_sub.go,example_test.go}`

**Design:** `type Broker struct` 持有 `[]chan string`(訂閱者)。`Subscribe() <-chan string` 新增一條 channel;`Publish(msg)` 送給所有訂閱者;`Close()` 收尾。Example 中每個訂閱者收同一則訊息;為輸出穩定,訂閱者收到後把結果集中排序或用固定順序印。角色:Publisher=`Broker.Publish`,Subscribers=各 channel。

**Example Output(`ExampleBroker`):**
```
sub got: hello
sub got: hello
```
(兩個訂閱者各收到 "hello";用 WaitGroup 等兩者都收到再結束)

**Verify:** `cd go/27_pub_sub && go test -race ./...`。Commit: `feat(go): add 27 pub sub example`。

---

## Self-Review

**Spec coverage:** 27 個 pattern 全數對應到 Task 1–27(21 為驗證任務)。並發 4 種、GoF 23 種齊全,編號照 GoF 書本順序,strategy=21。✓

**Placeholder scan:** 無 TBD/TODO。每個任務都有明確 module 名、檔案、角色設計、Example 預期 Output、驗證指令與 commit 訊息。實作碼僅在 Task 1 全寫(範本),其餘給足規格(介面簽名 + 預期 Output)供實作者寫慣用碼——這是本計畫刻意的取捨(避免把 27 份實作預先寫進計畫)。

**Type consistency:** 各 pattern 獨立 module、互不 import,無跨任務型別依賴。每個任務內的介面/型別自洽。並發任務用 `-race` 驗證。✓

**注意事項給實作者:** 凡是輸出多筆且順序可能不穩(map 迭代、goroutine 完成序、pub/sub 廣播)的 Example,務必在程式碼中排序或用 WaitGroup/固定順序收斂,確保 `// Output:` 逐字比對穩定通過。
