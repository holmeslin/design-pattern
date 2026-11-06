// 方法叫 interpret 而不是 eval,避免和全域 eval 混淆。
export const num = (n) => ({ interpret: () => n });
export const add = (left, right) => ({ interpret: () => left.interpret() + right.interpret() });
export const mul = (left, right) => ({ interpret: () => left.interpret() * right.interpret() });
