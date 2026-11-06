export class RealImage {
  // 刻意用公開欄位:原生 Proxy 看不到 #private,trap 在類別外也讀不到 #filename。
  filename;

  constructor(filename) {
    this.filename = filename;
  }

  load() {
    return `loading ${this.filename}`;
  }
}

// 原生 Proxy 就是代理:攔截 load,第一次轉給真物件,之後走快取。
export function withCache(image) {
  let loaded = false;
  return new Proxy(image, {
    get(target, prop) {
      if (prop !== 'load') return Reflect.get(target, prop);
      return () => {
        if (loaded) return `cached ${target.filename}`;
        loaded = true;
        // 對 target 呼叫,不用 receiver(proxy):方法內若讀 #private 會 TypeError。
        return target.load();
      };
    },
  });
}
