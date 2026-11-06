export class Account {
  balance = 0;
}

// JS 慣用:command 就是 { execute, undo } 物件字面值,不需要 Command 類別。
export function deposit(account, amount) {
  return {
    execute: () => { account.balance += amount; },
    undo: () => { account.balance -= amount; },
  };
}

export class History {
  #done = [];

  run(command) {
    command.execute();
    this.#done.push(command);
  }

  undoLast() {
    this.#done.pop()?.undo();
  }
}
