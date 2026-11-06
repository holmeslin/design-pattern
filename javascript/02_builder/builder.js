// 每個設定方法回傳 this,呼叫端才能一路鏈式呼叫;build() 才真正組出字串。

export class QueryBuilder {
  #columns = [];
  #table = '';
  #condition = '';

  select(...columns) {
    this.#columns = columns;
    return this;
  }

  from(table) {
    this.#table = table;
    return this;
  }

  where(condition) {
    this.#condition = condition;
    return this;
  }

  build() {
    const sql = `SELECT ${this.#columns.join(', ')} FROM ${this.#table}`;
    return this.#condition ? `${sql} WHERE ${this.#condition}` : sql;
  }
}
