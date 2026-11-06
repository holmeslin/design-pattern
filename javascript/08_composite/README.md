# Composite(組合)

把物件組成**樹狀結構**,讓呼叫端用同一種方式對待「單一物件」和「物件的集合」。

本範例:檔案系統,`folder` 可以放 `file` 或其他 `folder`,`size()` 遞迴加總。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Component** | 「有 `size()` 的物件」這個形狀 | 葉與容器共用的介面 |
| **Leaf** | `file(size)` | 沒有子節點,直接回傳自己的大小 |
| **Composite** | `folder()` | 持有子節點,`add` 加入、`size()` 遞迴加總 |

`file`、`folder` 用小寫工廠函式而不是 class,避免和 Node 全域的 `File` 類別撞名。

## 執行

```sh
cd javascript && node --test 08_composite/*.test.js
```
