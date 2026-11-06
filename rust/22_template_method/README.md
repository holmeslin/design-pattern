# Template Method(樣板方法)

在基底定義**演算法的骨架**,把其中幾個步驟留給子類實作;子類只填步驟,不改順序。

本範例:資料處理作業。`run` 固定「讀取 -> 處理 -> 輸出」,`EvenFilterJob` 讀入 `[1, 2, 3, 4]`,過濾出偶數,輸出 `"[2 4]"`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractClass** | `trait DataJob` | `run` 是樣板方法;`read`、`process`、`write` 是待實作的步驟 |
| **ConcreteClass** | `EvenFilterJob` | 實作三個步驟 |

## Rust 慣用寫法

trait 的 default method 就是樣板方法:`run` 在 trait 裡有預設實作,呼叫三個沒有預設實作的步驟。漏實作任何一個步驟,在**編譯期**就會報錯;JS 版要到執行期呼叫時才會丟 `not implemented`。

Rust 沒有 `final`,實作者技術上可以覆寫 `run`,所以「不覆寫 `run`」只是約定,編譯器不會擋。

## 執行

```sh
cd rust && cargo test -p template_method
```
