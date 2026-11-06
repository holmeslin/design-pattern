// JS 的工廠就是「函式 + 物件查表」,不需要 Creator 類別階層;新增管道只要多一個 key。

const creators = {
  email: () => ({ send: (msg) => `email: ${msg}` }),
  sms: () => ({ send: (msg) => `sms: ${msg}` }),
};

export function createNotifier(channel) {
  // hasOwn 避免 'constructor'、'toString' 這類繼承來的 key 被誤當成管道
  if (!Object.hasOwn(creators, channel)) {
    throw new Error(`unknown channel: "${channel}"`);
  }
  return creators[channel]();
}
