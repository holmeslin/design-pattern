# Newtype(新型別)

把既有型別包進**只有一個欄位的 tuple struct**,得到一個全新的型別:編譯器不會把它和原本的型別混為一談。

本範例:`Meters` / `Kilometers` 兩種長度單位,以及替 `Vec<String>` 實作 `Display` 的 `Tags`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Newtype** | `Meters(f64)`、`Kilometers(f64)` | 包住 `f64`,讓「單位不同」變成型別錯誤 |
| **轉換** | `impl From<Kilometers> for Meters` | 唯一的單位換算處(×1000) |
| **運算** | `impl Add for Meters` | 只允許公尺加公尺 |
| **包裝外部型別** | `Tags(Vec<String>)` | 本地型別,可以實作外部 trait `Display` |

## Rust 慣用寫法

**零成本。** `Meters` 加上 `#[repr(transparent)]` 後,記憶體佈局與 ABI 都保證和 `f64` 相同,包裝本身沒有執行期成本。差別只存在於編譯期:`Meters(1.0) + Kilometers(2.0)` 因為沒有 `Add<Kilometers> for Meters` 而編譯失敗。`lib.rs` 的 doctest 放了一對範例,只差右運算元,一個能編譯、一個標成 `compile_fail`。

stable 版 rustdoc 不檢查 `compile_fail` 的錯誤碼,任何編譯錯誤都會讓它「通過」,所以才需要旁邊那個能編譯的對照組,證明失敗確實來自那個運算元。

**繞過 orphan rule。** `Display` 與 `Vec<String>` 都定義在 std,在本 crate 裡直接 `impl Display for Vec<String>` 會被 orphan rule 拒絕。把它包成本地型別 `Tags` 之後,就可以替 `Tags` 實作 `Display`。代價是呼叫端得透過 `.0` 或 `Tags(...)` 存取內部的 `Vec`。

## 執行

```sh
cd rust && cargo test -p newtype
```
