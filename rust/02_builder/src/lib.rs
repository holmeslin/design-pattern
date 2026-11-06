// Builder 以 `self -> Self` 消耗並回傳自己,鏈式呼叫不需要 `&mut`。
#[derive(Default)]
pub struct QueryBuilder {
    columns: Vec<String>,
    table: String,
    condition: Option<String>,
}

impl QueryBuilder {
    pub fn new() -> Self {
        Self::default()
    }

    pub fn select(mut self, columns: &[&str]) -> Self {
        self.columns = columns.iter().map(|c| c.to_string()).collect();
        self
    }

    pub fn from(mut self, table: &str) -> Self {
        self.table = table.to_string();
        self
    }

    // `where` 是關鍵字,所以叫 where_clause。
    pub fn where_clause(mut self, condition: &str) -> Self {
        self.condition = Some(condition.to_string());
        self
    }

    pub fn build(&self) -> String {
        let mut sql = format!("SELECT {} FROM {}", self.columns.join(", "), self.table);
        if let Some(condition) = &self.condition {
            sql.push_str(" WHERE ");
            sql.push_str(condition);
        }
        sql
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    // 完整鏈式呼叫:有 WHERE
    #[test]
    fn builds_query_with_where() {
        let sql = QueryBuilder::new()
            .select(&["id", "name"])
            .from("users")
            .where_clause("age > 18")
            .build();
        assert_eq!(sql, "SELECT id, name FROM users WHERE age > 18");
    }

    // 不呼叫 where_clause 就不會有 WHERE 子句
    #[test]
    fn omits_where_when_not_set() {
        let sql = QueryBuilder::new().select(&["id"]).from("users").build();
        assert_eq!(sql, "SELECT id FROM users");
    }
}
