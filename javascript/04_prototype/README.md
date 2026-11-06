# Prototype(原型)

以**現成物件為模板複製出新物件**,而不是從頭 `new` 再逐欄位設定。

本範例:設定物件的深拷貝,改複本的陣列不會動到原型。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Prototype** | 被複製的設定物件;`clone(config)` 是它的 clone 操作 | 當作模板,複製出可獨立修改的複本 |
| **Client** | 呼叫 `clone(config)` 的程式碼 | 拿複本去改,不動原型 |

JS 原生已內建:`structuredClone` 就是深拷貝,不用自己遞迴複製。

注意:`structuredClone` 只處理純資料,class instance 會丟掉 prototype 與方法,含函式的物件會 throw `DataCloneError`。

注意:JS 語言本身的 prototype chain(`Object.create`)是「共享」不是「複製」,和 GoF Prototype 是兩回事。

## 執行

```sh
cd javascript && node --test 04_prototype/*.test.js
```
