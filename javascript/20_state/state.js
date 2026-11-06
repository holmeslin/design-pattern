// 每個 state 自己決定「下一步轉去哪」;next 是延遲呼叫,所以可以參照後面才宣告的 state。
export const created = { name: 'created', next: () => paid };
export const paid = { name: 'paid', next: () => shipped };
export const shipped = {
  name: 'shipped',
  next() {
    throw new Error('order already shipped');
  },
};

export class Order {
  #state = created;

  advance() {
    const from = this.#state;
    this.#state = from.next();
    return `${from.name} -> ${this.#state.name}`;
  }
}
