# Observer(觀察者)

觀察者模式定義物件間一對多的依賴關係,當主題狀態改變時會自動通知所有訂閱的觀察者。

## 角色

| 角色 | 對應程式碼 | 說明 |
| --- | --- | --- |
| Subject(主題) | `Station` | 維護觀察者清單,狀態改變時依訂閱順序發出通知 |
| Observer(觀察者介面) | `Observer` | 定義 `Notify(temp int)` 方法 |
| ConcreteObserver(具體觀察者) | `display` | 實作 `Observer`,收到通知時印出名稱與溫度 |

## 執行

```sh
go test ./...
```
