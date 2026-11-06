# Composite(組合)

把單一物件與物件群組組成**樹狀結構**,並讓呼叫端用同一個介面對待「葉」與「容器」。

本範例:檔案系統,`File` 有大小,`Folder` 可包含檔案與其他資料夾,`size()` 遞迴算出總大小。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Component** | `enum Node` 與 `size()` | 葉與容器共用的型別與操作 |
| **Leaf** | `Node::File(u64)` | 沒有子節點,直接回傳自己的大小 |
| **Composite** | `Node::Folder(Vec<Node>)` | 持有子節點,遞迴加總 |

## Rust 慣用寫法

節點種類固定時,用 `enum` + `match` 是 Rust 最自然的 Composite:編譯器會檢查 `match` 有沒有漏掉任何一種節點,也不需要 `Box<dyn ...>`。只有在種類需要對外開放擴充(讓別的 crate 加新節點)時,才改用 `trait` + `Vec<Box<dyn Component>>`。

## 執行

```sh
cd rust && cargo test -p composite
```
