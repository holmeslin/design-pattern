//! 單位不同的數值各包成自己的型別,混用會在編譯期被擋下。
//!
//! 合法:同為 `Meters` 可以相加。
//!
//! ```
//! use newtype::{Kilometers, Meters};
//!
//! assert_eq!(Meters(1.0) + Meters(2.0), Meters(3.0));
//! ```
//!
//! 非法:`Meters` 只實作了 `Add<Meters>`,和 `Kilometers` 相加編譯失敗。
//! 上面那個範例與這個只差右運算元(`Meters(2.0)` 換成 `Kilometers(2.0)`),
//! 證明失敗來自這一處。
//!
//! ```compile_fail
//! use newtype::{Kilometers, Meters};
//!
//! assert_eq!(Meters(1.0) + Kilometers(2.0), Meters(3.0));
//! ```

use std::fmt;
use std::ops::Add;

// repr(transparent):保證與 f64 有相同的記憶體佈局與 ABI,包一層不花任何執行期成本。
#[derive(Debug, Clone, Copy, PartialEq)]
#[repr(transparent)]
pub struct Meters(pub f64);

#[derive(Debug, Clone, Copy, PartialEq)]
#[repr(transparent)]
pub struct Kilometers(pub f64);

impl From<Kilometers> for Meters {
    fn from(km: Kilometers) -> Self {
        Meters(km.0 * 1000.0)
    }
}

impl Add for Meters {
    type Output = Meters;

    fn add(self, rhs: Meters) -> Meters {
        Meters(self.0 + rhs.0)
    }
}

// orphan rule 不允許對 Vec<String> 實作 Display(兩者都是外部的),包成本地型別就可以。
pub struct Tags(pub Vec<String>);

impl fmt::Display for Tags {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "[{}]", self.0.join(", "))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 公里先轉成公尺才能相加,單位換算集中在 From
    #[test]
    fn adds_meters_after_converting_kilometers() {
        assert_eq!(
            Meters(500.0) + Meters::from(Kilometers(1.5)),
            Meters(2000.0)
        );
    }

    // 外部型別 Vec<String> 包成本地的 Tags,才能實作外部 trait Display
    #[test]
    fn tags_display_as_bracketed_list() {
        let tags = Tags(vec!["a".into(), "b".into()]);
        assert_eq!(tags.to_string(), "[a, b]");
    }
}
