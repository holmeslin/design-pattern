// 每個階段都是 async generator:下游 next() 才會向上游拉一筆,所以天生有 backpressure。
export async function* filter(source, predicate) {
  for await (const item of source) {
    if (await predicate(item)) yield item;
  }
}

export async function* map(source, fn) {
  for await (const item of source) {
    yield fn(item);
  }
}

export async function* take(source, n) {
  if (n <= 0) return;
  let count = 0;
  for await (const item of source) {
    yield item;
    // 取滿就 return:for await 離開時會呼叫 source.return(),上游一路執行 finally 關閉資源。
    if (++count >= n) return;
  }
}
