// Clone trait 就是 Prototype:derive 對擁有所有權的欄位(String、Vec)做深拷貝。
// 若欄位換成 Rc<T>,clone 只複製指標,副本與原型會共享同一份資料。
#[derive(Debug, Clone, PartialEq)]
pub struct Config {
    pub name: String,
    pub tags: Vec<String>,
}

#[cfg(test)]
mod tests {
    use super::*;

    // clone 是深拷貝:改副本的 tags 不影響原型
    #[test]
    fn clone_is_independent_of_original() {
        let orig = Config {
            name: "base".to_string(),
            tags: vec!["a".to_string()],
        };
        let mut copy = orig.clone();
        copy.tags.push("b".to_string());

        assert_eq!(orig.tags, ["a"]);
        assert_eq!(copy.tags, ["a", "b"]);
    }
}
