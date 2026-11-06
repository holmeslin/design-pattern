use std::error::Error;
use std::fmt;

// Product:呼叫端只依賴這個 trait,不知道具體是 email 還是 sms。
pub trait Notifier {
    fn send(&self, msg: &str) -> String;
}

struct EmailNotifier;
struct SmsNotifier;

impl Notifier for EmailNotifier {
    fn send(&self, msg: &str) -> String {
        format!("email: {msg}")
    }
}

impl Notifier for SmsNotifier {
    fn send(&self, msg: &str) -> String {
        format!("sms: {msg}")
    }
}

#[derive(Debug, PartialEq)]
pub struct UnknownChannel(pub String);

impl fmt::Display for UnknownChannel {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "unknown channel: {:?}", self.0)
    }
}

impl Error for UnknownChannel {}

// Rust 沒有繼承,工廠方法就是一個依參數 match 的函式;未知值回 Err 而不是 panic。
pub fn new_notifier(channel: &str) -> Result<Box<dyn Notifier>, UnknownChannel> {
    match channel {
        "email" => Ok(Box::new(EmailNotifier)),
        "sms" => Ok(Box::new(SmsNotifier)),
        other => Err(UnknownChannel(other.to_string())),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 依 channel 建立對應的 Notifier,呼叫端只看 trait
    #[test]
    fn creates_notifier_by_channel() {
        assert_eq!(new_notifier("email").unwrap().send("hello"), "email: hello");
        assert_eq!(new_notifier("sms").unwrap().send("hello"), "sms: hello");
    }

    // Box<dyn Notifier> 沒有 Debug,所以用 .err() 而不是 unwrap_err()
    #[test]
    fn rejects_unknown_channel() {
        let err = new_notifier("fax").err().unwrap();
        assert_eq!(err.to_string(), "unknown channel: \"fax\"");
    }
}
