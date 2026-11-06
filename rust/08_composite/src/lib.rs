// 節點種類固定,用 enum 表示「葉(File)或容器(Folder)」,不需要 trait object。
pub enum Node {
    File(u64),
    Folder(Vec<Node>),
}

impl Node {
    // 葉與容器共用同一個操作:容器遞迴加總子節點。
    pub fn size(&self) -> u64 {
        match self {
            Node::File(size) => *size,
            Node::Folder(children) => children.iter().map(Node::size).sum(),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 資料夾與檔案用同一個 size() 遞迴加總
    #[test]
    fn folder_sums_nested_sizes() {
        let tree = Node::Folder(vec![Node::File(100), Node::Folder(vec![Node::File(200)])]);
        assert_eq!(tree.size(), 300);
    }
}
