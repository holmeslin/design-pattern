# Bridge(橋接)

把「抽象」與「實作」拆成兩條獨立演化的軸,抽象持有實作的引用,兩邊可以自由搭配,不必為每種組合各寫一個子類別。

本範例:遙控器(抽象)控制不同裝置(實作),TV 與 Radio 都能用同一個 `RemoteControl` 調音量。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Abstraction** | `RemoteControl<D>` | 高階操作(`volume_up`、`status`),持有一個 Implementor |
| **Implementor** | `trait Device` | 低階介面(`name`、`volume`、`set_volume`) |
| **ConcreteImplementor** | `Tv`、`Radio` | 各種裝置的具體實作 |

## Rust 慣用寫法

用泛型 `RemoteControl<D: Device>` 持有裝置,是**靜態分派**:編譯器為每種裝置各產生一份程式碼,沒有執行期成本,但單一 `RemoteControl` 在編譯期就綁死了裝置型別。Go 用 interface 型別的欄位(執行期分派),JS 用 duck typing(只要有對應方法就能用)。若需要在執行期才決定裝置,可把欄位改成 `Box<dyn Device>`。

## 執行

```sh
cd rust && cargo test -p bridge
```
