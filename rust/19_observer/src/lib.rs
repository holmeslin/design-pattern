// 訂閱憑證:欄位私有,只有 subscribe 能產生。
#[derive(Clone, Copy, Debug, PartialEq)]
pub struct SubscriptionId(usize);

// observer 就是閉包,不另外定義 Observer trait。
type Observer = Box<dyn Fn(i32) -> String>;

#[derive(Default)]
pub struct Station {
    observers: Vec<(SubscriptionId, Observer)>,
    next_id: usize,
}

impl Station {
    pub fn subscribe(&mut self, observer: impl Fn(i32) -> String + 'static) -> SubscriptionId {
        let id = SubscriptionId(self.next_id);
        self.next_id += 1;
        self.observers.push((id, Box::new(observer)));
        id
    }

    // 不認得的 id 直接忽略,不 panic。
    pub fn unsubscribe(&mut self, id: SubscriptionId) {
        self.observers.retain(|(existing, _)| *existing != id);
    }

    pub fn set_temp(&mut self, temp: i32) -> Vec<String> {
        self.observers
            .iter()
            .map(|(_, notify)| notify(temp))
            .collect()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 依訂閱順序通知每個 observer,並收集它們的回傳值
    #[test]
    fn notifies_observers_in_subscription_order() {
        let mut station = Station::default();
        station.subscribe(|t| format!("phone: {t}"));
        station.subscribe(|t| format!("window: {t}"));
        assert_eq!(station.set_temp(25), ["phone: 25", "window: 25"]);
    }

    // 退訂後,該 observer 不再收到通知
    #[test]
    fn unsubscribed_observer_is_not_notified() {
        let mut station = Station::default();
        let phone = station.subscribe(|t| format!("phone: {t}"));
        station.subscribe(|t| format!("window: {t}"));
        station.unsubscribe(phone);
        assert_eq!(station.set_temp(30), ["window: 30"]);
    }
}
