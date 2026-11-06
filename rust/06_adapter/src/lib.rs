// 第三方 SDK:介面我們改不了,只收「元」。
pub struct StripeSdk;

impl StripeSdk {
    pub fn charge(&self, dollars: f64) -> String {
        format!("charged ${dollars:.2}")
    }
}

// 呼叫端期望的介面:金額單位是「分」。
pub trait PaymentProcessor {
    fn pay(&self, cents: u64) -> String;
}

// Adapter 持有被適配者,把分換成元後委派。
pub struct StripeAdapter {
    sdk: StripeSdk,
}

impl StripeAdapter {
    pub fn new(sdk: StripeSdk) -> Self {
        Self { sdk }
    }
}

impl PaymentProcessor for StripeAdapter {
    fn pay(&self, cents: u64) -> String {
        self.sdk.charge(cents as f64 / 100.0)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 呼叫端只認 PaymentProcessor(分),Adapter 把它轉成 SDK 要的元
    #[test]
    fn adapter_converts_cents_to_dollars() {
        assert_eq!(StripeAdapter::new(StripeSdk).pay(1234), "charged $12.34");
    }
}
