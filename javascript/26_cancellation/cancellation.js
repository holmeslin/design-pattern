// 可取消的非同步操作:signal 已取消就立刻 reject,途中取消則清掉計時器再 reject。
export function slowApi(ms, { signal } = {}) {
  return new Promise((resolve, reject) => {
    if (signal?.aborted) return reject(signal.reason);

    const onAbort = () => {
      clearTimeout(timer); // 不清掉的話,計時器會讓行程多撐到 ms 後才結束
      reject(signal.reason);
    };
    const timer = setTimeout(() => {
      signal?.removeEventListener('abort', onAbort);
      resolve('done');
    }, ms);
    signal?.addEventListener('abort', onAbort, { once: true });
  });
}

// 逾時與使用者取消合併成同一個 signal;reason 分別是 TimeoutError / AbortError,呼叫端可據此分辨。
export function withTimeout(task, timeoutMs, { signal } = {}) {
  const timeout = AbortSignal.timeout(timeoutMs);
  return task(signal ? AbortSignal.any([timeout, signal]) : timeout);
}
