// 函式是一級公民:裝飾者就是「吃 handler、回傳新 handler」的高階函式。

export const base = (req) => 'ok';

export const logging = (handler) => (req) => '[log] ' + handler(req);

export const auth = (handler) => (req) => '[auth] ' + handler(req);
