export class Ring {
  #items;

  constructor(...items) {
    this.#items = items;
  }

  // generator 回傳的物件繼承 Iterator.prototype,所以 map / toArray 等 helpers 自動可用。
  *[Symbol.iterator]() {
    yield* this.#items;
  }
}
