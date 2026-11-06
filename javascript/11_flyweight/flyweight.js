export class CharFormatFactory {
  #cache = new Map();

  get(font, size) {
    const key = `${font}:${size}`;
    if (!this.#cache.has(key)) {
      // 共享物件必須不可變,否則一個使用者改了,所有人都跟著變。
      this.#cache.set(key, Object.freeze({ font, size }));
    }
    return this.#cache.get(key);
  }
}
