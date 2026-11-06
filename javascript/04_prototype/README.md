# Prototype(原型)

以**現成物件為模板複製出新物件**,而不是從頭 `new` 再逐欄位設定。

本範例:設定物件的深拷貝,改複本的陣列不會動到原型。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Prototype** | 被複製的設定物件 | 當作模板 |
| **Client** | `clone(config)` | 複製模板,取得可獨立修改的複本 |

JS 原生已內建:`structuredClone` 就是深拷貝,不用自己遞迴複製。

注意:JS 語言本身的 prototype chain(`Object.create`)是「共享」不是「複製」,和 GoF Prototype 是兩回事。

## 執行

```sh
cd javascript && node --test 04_prototype/*.test.js
```
