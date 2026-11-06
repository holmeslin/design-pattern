export class DataJob {
  // 固定流程:順序寫死在這裡,子類別只能覆寫步驟,不能改順序。
  run() {
    return this.write(this.process(this.read()));
  }

  read() {
    throw new Error('not implemented');
  }

  process() {
    throw new Error('not implemented');
  }

  write() {
    throw new Error('not implemented');
  }
}

export class EvenFilterJob extends DataJob {
  read() {
    return [1, 2, 3, 4];
  }

  process(data) {
    return data.filter((n) => n % 2 === 0);
  }

  write(data) {
    return `[${data.join(' ')}]`;
  }
}
