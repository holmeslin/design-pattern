export class Handler {
  #next = null;

  // 回傳 handler 本身,才能 a.setNext(b).setNext(c) 串成鏈。
  setNext(handler) {
    this.#next = handler;
    return handler;
  }

  handle(req) {
    return this.#next ? this.#next.handle(req) : 'ok';
  }
}

export class AuthHandler extends Handler {
  handle(req) {
    if (!req.user) throw new Error('auth failed');
    return super.handle(req);
  }
}

export class RateLimitHandler extends Handler {
  #limit;
  #count = 0;

  constructor(limit = 3) {
    super();
    this.#limit = limit;
  }

  handle(req) {
    this.#count++;
    if (this.#count > this.#limit) throw new Error('rate limited');
    return super.handle(req);
  }
}

export class BodyHandler extends Handler {
  handle(req) {
    if (!req.body) throw new Error('empty body');
    return super.handle(req);
  }
}
