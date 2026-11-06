// 目前正在執行的 effect;signal.get() 靠它知道「是誰在讀我」,這就是自動追蹤相依的來源。
let currentEffect = null;

export function signal(initial) {
  let value = initial;
  const subscribers = new Set();
  return {
    get() {
      if (currentEffect) subscribers.add(currentEffect);
      return value;
    },
    set(next) {
      if (Object.is(next, value)) return;
      value = next;
      // 迭代副本:effect 重跑時會再 add 自己,直接迭代 Set 會無限循環。
      for (const run of [...subscribers]) run();
    },
  };
}

export function effect(fn) {
  const run = () => {
    // 還原成 prev 而不是 null:effect 可能巢狀執行,內層結束後外層仍要繼續追蹤。
    const prev = currentEffect;
    currentEffect = run;
    try {
      fn();
    } finally {
      currentEffect = prev;
    }
  };
  run();
}

// shortcut: computed 是 eager 的,沒有排程與 glitch-free 保證;需要時改用 TC39 Signals polyfill 或框架實作
export function computed(fn) {
  const s = signal(undefined);
  effect(() => s.set(fn()));
  return { get: s.get };
}
