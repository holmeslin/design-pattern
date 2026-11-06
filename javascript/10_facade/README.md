# Facade(外觀)

為一組複雜的子系統提供**單一、簡單的入口**,呼叫端不必知道子系統怎麼協作。

本範例:家庭劇院,`watchMovie()` 一個呼叫就依序調暗燈光、開投影機、開擴大機。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Facade** | `HomeTheater` | 對外的簡單入口,知道要依什麼順序驅動哪些子系統 |
| **Subsystem** | `lights`、`projector`、`amplifier` | 各自負責一塊功能,不知道 Facade 的存在 |

子系統沒有匯出,呼叫端只能透過 `HomeTheater` 使用它們。

## 執行

```sh
cd javascript && node --test 10_facade/*.test.js
```
