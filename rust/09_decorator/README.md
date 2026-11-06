# Decorator(裝飾者)

用「包裝」的方式動態替物件加上額外行為,被包裝的物件與包裝者實作同一個介面,所以可以層層疊加。

本範例:請求處理器,`Base` 回傳 `"ok"`,`Logging`、`Auth` 各自在結果前面加上前綴。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Component** | `trait Handler` | 被裝飾者與裝飾者共同的介面 |
| **ConcreteComponent** | `Base` | 最內層、真正處理請求的物件 |
| **Decorator** | `Logging<H>`、`Auth<H>` | 持有一個 `Handler`,委派之後加上自己的行為 |

## Rust 慣用寫法

用泛型包裝(`Logging<H: Handler>`)是零成本的:`Logging<Auth<Base>>` 在編譯期就是一個具體型別,編譯後沒有動態分派。代價是裝飾的組合在編譯期固定;要在執行期動態決定順序,才改用 `Box<dyn Handler>`,並加上轉發實作 `impl<H: Handler + ?Sized> Handler for Box<H>`(否則 `Box<dyn Handler>` 本身不是 `Handler`,傳不進 `Logging<H>`)。生態系的 `tower` crate 的 `Layer` / `Service` 就是這個模式。

## 執行

```sh
cd rust && cargo test -p decorator
```
