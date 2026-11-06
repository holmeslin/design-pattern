# Flyweight(享元)

讓大量相似的物件**共享相同的內部狀態**,只存一份,節省記憶體。

本範例:文字編輯器的字元格式,相同 font + size 的格式只建立一次,所有字元共用同一個物件。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Flyweight** | `{ font, size }`(凍結物件) | 被共享的內部狀態,不可變 |
| **FlyweightFactory** | `CharFormatFactory` | 用 `Map` 快取,key 為 `font:size`,已存在就回傳同一個 |
| **Client** | 呼叫 `get(font, size)` 的程式 | 取得共享物件,不自己 new |

共享物件用 `Object.freeze` 凍結:被多處共用的東西一旦可變,一處修改就會影響所有使用者。

## 執行

```sh
cd javascript && node --test 11_flyweight/*.test.js
```
