// ES module 只會被評估一次,匯出的物件天生就是單例;狀態藏在 module 內,外面只拿得到 get / set。

const store = new Map();

export const config = {
  get: (key) => store.get(key),
  set: (key, value) => {
    store.set(key, value);
  },
};
