// 第三方 SDK(Adaptee):收「元」,我們改不了它。
export class StripeSDK {
  charge(dollars) {
    return `charged $${dollars.toFixed(2)}`;
  }
}

// Adapter:對外維持「收分」的介面,內部換算後委派給 SDK。
export class StripeAdapter {
  #sdk;

  constructor(sdk) {
    this.#sdk = sdk;
  }

  pay(cents) {
    return this.#sdk.charge(cents / 100);
  }
}
