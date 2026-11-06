# Facade(外觀)

替一組複雜的子系統提供**一個簡單的統一入口**,呼叫端不用知道內部有哪些零件、要按什麼順序呼叫。

本範例:家庭劇院,`watch_movie()` 一次依序啟動燈光、投影機、擴大機。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Facade** | `HomeTheater` | 對外唯一入口,協調子系統的呼叫順序 |
| **Subsystem** | `Lights`、`Projector`、`Amplifier` | 實際做事的零件;不對外公開 |

## Rust 慣用寫法

Rust 的可見性規則直接幫 Facade 把關:子系統型別不加 `pub`,外部 crate 根本看不到也建構不了,只能透過 `HomeTheater`。`HomeTheater` 欄位私有,用 `#[derive(Default)]` 提供預設建構,不必手寫 `new()`。

## 執行

```sh
cd rust && cargo test -p facade
```
