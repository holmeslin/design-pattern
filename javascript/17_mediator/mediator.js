export class ChatRoom {
  #users = [];

  join(user) {
    this.#users.push(user);
  }

  // 依加入順序轉發,跳過發送者。
  broadcast(from, msg) {
    for (const user of this.#users) {
      if (user !== from) user.receive(from.name, msg);
    }
  }
}

export class User {
  name;
  inbox = [];
  #room;

  constructor(name, room) {
    this.name = name;
    this.#room = room;
    room.join(this);
  }

  send(msg) {
    this.#room.broadcast(this, msg);
  }

  receive(from, msg) {
    this.inbox.push(`${this.name} received from ${from}: ${msg}`);
  }
}
