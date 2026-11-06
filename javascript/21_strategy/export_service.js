// Context:只查表 + 轉交,不認識任何具體格式。
export class ExportService {
  #strategies;

  constructor(strategies) {
    this.#strategies = strategies;
  }

  export(rows, format) {
    const strategy = this.#strategies.get(format);
    if (!strategy) throw new Error(`unsupported: "${format}"`);
    return strategy.export(rows);
  }
}
