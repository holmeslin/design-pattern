# Observer(觀察者)

定義物件間一對多的依賴關係,主題狀態改變時,自動通知所有訂閱的觀察者。

本範例:氣象站,溫度改變時通知手機與窗戶兩個顯示器。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Subject** | `Station` | 溫度改變時發出 `temp` 事件;訂閱清單由 `EventTarget` 管理 |
| **Observer** | `display(name, out)` 回傳的 listener | 收到事件後把 `名稱: 溫度` 推入 `out` |

JS 原生已內建:`EventTarget`(瀏覽器與 Node 皆有)就是 Subject,不用自己維護訂閱清單;Go 版則要自己寫 `Observer` 介面與清單。

## 執行

```sh
cd javascript && node --test 19_observer/*.test.js
```
