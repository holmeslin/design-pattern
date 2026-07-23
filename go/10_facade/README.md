# Facade(門面)

門面模式為一組複雜的子系統提供一個簡化的統一介面,讓使用者不需理解內部細節即可完成常見操作。

## 角色

| 角色 | 對應程式 | 說明 |
| --- | --- | --- |
| Facade(門面) | `HomeTheater` | 提供簡化的 `WatchMovie` 方法,統一協調各子系統 |
| Subsystem(子系統) | `lights`、`projector`、`amplifier` | 各自負責獨立的功能,彼此不相依 |

## 執行

```sh
go test ./...
```
