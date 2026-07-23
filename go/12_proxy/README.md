# Proxy(代理)

代理模式為另一個物件提供替身以控制對它的存取,本範例以快取代理延遲並快取昂貴的載入動作。

## 角色

| 角色 | 對應型別 | 說明 |
| --- | --- | --- |
| Subject(抽象主體) | `Image` | 定義 RealSubject 與 Proxy 共同遵守的介面 |
| RealSubject(真實主體) | `realImage` | 實際執行昂貴載入動作的物件 |
| Proxy(代理) | `ProxyImage` | 持有 `*realImage` 並加上快取邏輯的代理 |

## 執行

```sh
go test ./...
```
