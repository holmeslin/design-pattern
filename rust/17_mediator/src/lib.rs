// 成員只持有 UserId 這個 handle,不持有聊天室的參考,所以不會形成循環。
#[derive(Clone, Copy, Debug, PartialEq)]
pub struct UserId(usize);

struct User {
    name: String,
    inbox: Vec<String>,
}

// Mediator 擁有所有成員;成員之間不直接溝通,一律經過這裡轉發。
#[derive(Default)]
pub struct ChatRoom {
    users: Vec<User>,
}

impl ChatRoom {
    pub fn join(&mut self, name: &str) -> UserId {
        self.users.push(User {
            name: name.to_string(),
            inbox: Vec::new(),
        });
        UserId(self.users.len() - 1)
    }

    // 用 get 而不是索引:別間聊天室的 id 可能越界,library 不能 panic。
    pub fn send(&mut self, from: UserId, msg: &str) {
        let Some(sender) = self.users.get(from.0).map(|u| u.name.clone()) else {
            return;
        };
        for (i, user) in self.users.iter_mut().enumerate() {
            if i != from.0 {
                user.inbox
                    .push(format!("{} received from {sender}: {msg}", user.name));
            }
        }
    }

    pub fn inbox(&self, id: UserId) -> &[String] {
        self.users.get(id.0).map_or(&[], |u| &u.inbox)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 訊息經由聊天室依加入順序轉發,發送者自己收不到
    #[test]
    fn send_forwards_to_everyone_except_sender() {
        let mut room = ChatRoom::default();
        let alice = room.join("Alice");
        let bob = room.join("Bob");
        let carol = room.join("Carol");

        room.send(alice, "hi");

        assert_eq!(room.inbox(bob), ["Bob received from Alice: hi"]);
        assert_eq!(room.inbox(carol), ["Carol received from Alice: hi"]);
        assert!(room.inbox(alice).is_empty());
    }

    // UserId 只能由 join 產生;拿到別間聊天室的 id 時不 panic,send 忽略、inbox 為空
    #[test]
    fn foreign_id_is_ignored() {
        let mut other = ChatRoom::default();
        other.join("Xavier");
        let stranger = other.join("Yvonne"); // 在 other 是第 2 位,room 只有 1 人

        let mut room = ChatRoom::default();
        let alice = room.join("Alice");
        room.send(stranger, "hi");

        assert!(room.inbox(alice).is_empty());
        assert!(room.inbox(stranger).is_empty());
    }
}
