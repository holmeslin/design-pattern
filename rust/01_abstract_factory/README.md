# Abstract Factory(抽象工廠)

一個工廠**成套產出「相關的一族產品」**,呼叫端只認工廠與產品的 trait,換一整族只換一個工廠。

本範例:UI 主題工廠,產出成套的 Button + Checkbox(Dark / Light)。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **AbstractFactory** | `trait GuiFactory` | 宣告成套產出哪些產品 |
| **ConcreteFactory** | `DarkFactory`、`LightFactory` | 產出同一族的具體產品 |
| **Product** | `trait Button`、`trait Checkbox` | 產品介面;具體產品(`DarkButton`…)不對外公開 |

## Rust 慣用寫法

產品以 `Box<dyn Button>` 回傳(trait object,執行期分派),因為主題通常是讀設定才決定。若編譯期就確定用哪一族,可改用泛型 `fn render_ui<F: GuiFactory>(f: &F)`,以靜態分派換取零成本。

## 執行

```sh
cd rust && cargo test -p abstract_factory
```
