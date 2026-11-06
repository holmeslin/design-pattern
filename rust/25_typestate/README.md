# Typestate(型別狀態)

把物件目前的**狀態編進型別**:不同狀態是不同的型別,每個型別只提供該狀態合法的方法,所以非法的操作順序根本編譯不過。

本範例:訂單 `created -> paid -> shipped`,只能依序 `pay()` 再 `ship()`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **狀態標記** | `Created`、`Paid`、`Shipped` | 沒有欄位的型別,只用來標示狀態 |
| **Context** | `Order<S>` | 帶狀態參數 `S`;`PhantomData<S>` 讓編譯器記住它 |
| **轉換** | `pay(self) -> Order<Paid>`、`ship(self) -> Order<Shipped>` | 取走舊狀態、回傳新狀態 |
| **共用方法** | `impl<S> Order<S>` 的 `id()` | 所有狀態都能呼叫 |

## Rust 慣用寫法

**對照 `20_state`。** 那一版是執行期狀態機:`Order` 持有 `Box<dyn State>`,對已出貨的訂單再推進,要到執行期才收到 `Err(OrderError::AlreadyShipped)`。typestate 版則是根本編譯不過:`Order<Created>` 沒有 `ship` 方法,`Order<Shipped>` 沒有任何轉換方法,所以不需要錯誤型別,也沒有可以回傳的 `Err`。

`pay(self)` 取走 `self`(所有權轉移),呼叫之後舊的 `Order<Created>` 就不能再用,因此同一個狀態不會被重複推進。`lib.rs` 的 doctest 放了兩對範例,每對只差一處:一個能編譯,一個標成 `compile_fail`(`Order::new(1).ship();`,以及 move 之後多出的 `o.pay();`)。stable 版 rustdoc 不檢查 `compile_fail` 的錯誤碼,任何編譯錯誤都會讓它「通過」,所以才需要能編譯的對照組,證明失敗確實來自那一處。

**零大小。** `Created`、`Paid`、`Shipped` 沒有欄位,`PhantomData<S>` 也一樣,大小都是 0:狀態不佔記憶體,也沒有儲存狀態值或在執行期比對狀態的程式碼,狀態只存在於編譯期的型別裡。

**限制。** 狀態必須在編譯期就知道。如果狀態是執行期才從資料庫或網路讀出來的,仍要用 `20_state` 的做法(或 enum)。

## 執行

```sh
cd rust && cargo test -p typestate
```
