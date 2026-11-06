// 同時最多 max 個任務在跑,多的先排隊;任務 settle(成功或失敗)後把名額交給下一個。
export function createLimiter(max) {
  let active = 0;
  const waiting = [];

  return async (task) => {
    if (active < max) {
      active++;
    } else {
      const { promise, resolve } = Promise.withResolvers();
      waiting.push(resolve);
      await promise; // 名額由前一個任務直接交棒,active 不變
    }

    try {
      return await task();
    } finally {
      // reject 時也要走到這裡,否則名額永遠不還,後面的任務會卡死
      const next = waiting.shift();
      if (next) next();
      else active--;
    }
  };
}
