export class Editor {
  #content = '';

  // 取代整份內容,不是附加。
  type(text) {
    this.#content = text;
  }

  get content() {
    return this.#content;
  }

  // 凍結快照;約定拿到的人只把它原封不動交回 restore(JS 無法強制)。
  save() {
    return Object.freeze({ content: this.#content });
  }

  restore(memento) {
    this.#content = memento.content;
  }
}
