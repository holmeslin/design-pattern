# Mediator(中介者)

用一個**中介者物件**封裝一組物件之間的互動,成員彼此不直接認識,一律透過中介者溝通。

本範例:聊天室。成員透過 `ChatRoom` 發訊息,聊天室依加入順序轉發給其他人(跳過發送者)。

## 角色

| 角色 | 對應本範例 | 職責 |
|------|-----------|------|
| **Mediator** | `ChatRoom` | 擁有所有成員,負責轉發訊息 |
| **Colleague** | `User`(私有) | 有名字與收件匣,不知道其他成員的存在 |
| **Handle** | `UserId` | 外部用來指稱成員的識別碼 |

## Rust 慣用寫法

成員和 mediator 若互相持有參考會形成循環,在 Rust 要用 `Rc<RefCell>` + `Weak` 才寫得出來。所以改成 mediator **擁有**所有成員(`Vec<User>`),外部只拿 `UserId` 這個 handle,要收發訊息時把 id 交給 `ChatRoom`。這樣所有權是單向的,借用檢查器不需要任何 `Rc` 或 `RefCell` 的協助。

`UserId` 的欄位是私有的,只有 `join` 能產生。一個 `UserId` 只對發出它的 `ChatRoom` 有效:拿別間聊天室的 id 來用,若超出範圍,內部用 `get` 而不是索引,所以 `send` 會忽略未知的發送者、`inbox` 回傳空切片,library 不會 panic;若剛好在範圍內,則會被當成這間聊天室的成員(不會被偵測到)。

## 執行

```sh
cd rust && cargo test -p mediator
```
