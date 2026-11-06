// EventTarget 是平台內建的 Subject,訂閱清單與通知順序它自己管。
export class Station extends EventTarget {
  setTemp(t) {
    this.dispatchEvent(new CustomEvent('temp', { detail: t }));
  }
}

// 回傳 listener 而不是直接訂閱,這樣呼叫端才拿得到同一個參考去 removeEventListener。
export function display(name, out) {
  return (event) => out.push(`${name}: ${event.detail}`);
}
