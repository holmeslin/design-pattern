// 實作 Symbol.dispose 的物件,離開 `using` 宣告所在的 block 時會自動被呼叫(含丟例外時)。
export class Connection {
  #name;
  #log;

  constructor(name, log) {
    this.#name = name;
    this.#log = log;
  }

  query(sql) {
    return `${this.#name}: ${sql}`;
  }

  [Symbol.dispose]() {
    this.#log.push(`close ${this.#name}`);
  }
}

// 取得與釋放都是非同步:建構靠 static async 工廠,釋放靠 Symbol.asyncDispose,由 `await using` 等待完成。
export class FileLock {
  #name;
  #log;

  constructor(name, log) {
    this.#name = name;
    this.#log = log;
  }

  static async acquire(name, log) {
    log.push(`lock ${name}`);
    return new FileLock(name, log);
  }

  async [Symbol.asyncDispose]() {
    this.#log.push(`unlock ${this.#name}`);
  }
}
