# Template Method(樣板方法)

把演算法的固定流程寫死一次,把其中會變的步驟留成「洞」,交給子類別填入不同實作。

本範例:資料處理工作,固定流程為 read -> process -> write,範例子類別讀入 `[1, 2, 3, 4]`、濾出偶數、輸出 `[2 4]`。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractClass** | `DataJob` | `run()` 是樣板方法,固定呼叫 `write(process(read()))`;三個步驟預設 throw `not implemented` |
| **ConcreteClass** | `EvenFilterJob` | 覆寫 `read`、`process`、`write` 三個 hook,不碰流程順序 |

Go 沒有繼承,所以用函式欄位當 hook;JS 有 `class extends`,這裡用的是 GoF 原版做法:子類別覆寫 hook。

## 執行

```sh
cd javascript && node --test 22_template_method/*.test.js
```
