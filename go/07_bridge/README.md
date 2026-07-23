# Bridge(橋接)

把「抽象」與「實作」拆成兩個獨立的體系,再用組合(而非繼承)把兩者橋接起來,讓雙方可以各自獨立擴充、互不牽制。

## 角色

| 角色 | 對應本專案 | 職責 |
|------|-----------|------|
| **Abstraction**(抽象) | `RemoteControl` | 對外提供操作入口,內部只透過 `Device` 介面委派工作 |
| **Implementor**(實作介面) | `Device` | 定義裝置必須提供的能力(`Name`、`SetVolume`、`Volume`) |
| **ConcreteImplementor**(具體實作) | `TV`、`Radio` | 各自實作 `Device` 介面的具體裝置 |

`RemoteControl` 只認識 `Device` 介面,完全不知道背後是電視還是收音機;要新增裝置種類時,只需新增一個實作 `Device` 的型別,不必更動 `RemoteControl` 半行程式碼。

## 執行

```sh
go test ./...
```
