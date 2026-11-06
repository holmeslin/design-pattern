// 對自己不擁有的型別(str)加方法:定義 trait 再為該型別實作。
pub trait TitleCase {
    fn to_title_case(&self) -> String;
}

impl TitleCase for str {
    // split_whitespace 會吃掉連續空白,所以輸出單字間一律只有單一空白。
    fn to_title_case(&self) -> String {
        self.split_whitespace()
            .map(|word| {
                let mut chars = word.chars();
                match chars.next() {
                    Some(first) => first.to_uppercase().chain(chars).collect::<String>(),
                    None => String::new(),
                }
            })
            .collect::<Vec<_>>()
            .join(" ")
    }
}

// 方法有預設實作,下面的 blanket impl 一行就讓所有符合條件的 Iterator 都擁有它。
pub trait SumSquares: Iterator<Item = i64> + Sized {
    fn sum_squares(self) -> i64 {
        self.map(|x| x * x).sum()
    }
}

impl<I: Iterator<Item = i64>> SumSquares for I {}

#[cfg(test)]
mod tests {
    // 使用端要 use 這些 trait,方法才看得到。
    use super::*;

    // 替 str 加上不屬於它的方法
    #[test]
    fn title_case_on_str() {
        assert_eq!("hello rust world".to_title_case(), "Hello Rust World");
    }

    // 元素型別是 i64,符合 blanket impl 的條件;答案與 Go worker pool 相同
    #[test]
    fn sum_squares_on_range() {
        assert_eq!((1..=5i64).sum_squares(), 55);
    }

    // 任何 Iterator<Item = i64> 都有這個方法
    #[test]
    fn sum_squares_on_vec_iter() {
        assert_eq!(vec![3i64, 4].into_iter().sum_squares(), 25);
    }
}
