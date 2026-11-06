# Bridge(橋接)

把「抽象」和「實作」拆成**兩條獨立的階層**,中間用組合(一個欄位)連起來,兩邊可以各自擴充而不互相牽動。

本範例:遙控器(抽象)操作任何裝置(實作),`TV`、`Radio` 都能被同一支 `RemoteControl` 控制音量。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Abstraction** | `RemoteControl` | 高層操作 `volumeUp()`、`status()`,持有一個 Implementor |
| **Implementor** | 「有 `name` / `volume` / `setVolume` 的物件」這個形狀 | 低層操作的約定 |
| **ConcreteImplementor** | `TV`、`Radio` | 各自的裝置實作 |

橋就是 `RemoteControl` 持有 device 的那個欄位:想加新遙控器(例如 `AdvancedRemote`)或新裝置,都只動自己那一側。

## 執行

```sh
cd javascript && node --test 07_bridge/*.test.js
```
