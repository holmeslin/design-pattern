pub struct Ring<T> {
    items: Vec<T>,
}

impl<T> Ring<T> {
    pub fn new(items: Vec<T>) -> Self {
        Self { items }
    }

    // 回傳 impl Iterator 隱藏具體型別;呼叫端只知道它產出 &T。
    pub fn iter(&self) -> impl Iterator<Item = &T> {
        self.items.iter()
    }
}

// 實作 IntoIterator for &Ring 後,`for x in &ring` 就能用,且不會消耗 ring。
impl<'a, T> IntoIterator for &'a Ring<T> {
    type Item = &'a T;
    type IntoIter = std::slice::Iter<'a, T>;

    fn into_iter(self) -> Self::IntoIter {
        self.items.iter()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // iter() 依序走訪每個元素
    #[test]
    fn iter_visits_in_order() {
        let ring = Ring::new(vec![10, 20, 30]);
        assert_eq!(ring.iter().copied().collect::<Vec<_>>(), vec![10, 20, 30]);
    }

    // 實作 IntoIterator for &Ring 後,for 迴圈可以直接用
    #[test]
    fn for_loop_over_reference() {
        let ring = Ring::new(vec![10, 20, 30]);
        let mut sum = 0;
        for x in &ring {
            sum += x;
        }
        assert_eq!(sum, 60);
    }

    // 內建的 adapter(map、filter…)自動可用,不必自己實作
    #[test]
    fn adapters_work_for_free() {
        let ring = Ring::new(vec![1, 2, 3]);
        assert_eq!(
            ring.iter().map(|x| x * 2).collect::<Vec<_>>(),
            vec![2, 4, 6]
        );
    }
}
