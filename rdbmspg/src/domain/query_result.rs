#[derive(Debug, Eq, PartialEq)]
pub struct QueryResult {
    columns: Vec<String>,
    rows: Vec<Option<String>>,
}

impl QueryResult {
    pub fn new(columns: Vec<String>, rows: Vec<Option<String>>) -> Self {
        Self { columns, rows }
    }

    pub fn columns(&self) -> &[String] {
        &self.columns
    }

    pub fn rows(&self) -> &[Option<String>] {
        &self.rows
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn query_result_new() {
        assert_eq!(
            QueryResult::new(
                vec!["a".to_owned(), "b".to_owned(), "c".to_owned()],
                vec![Some("a value".to_owned()), Some("12345".to_owned()), None],
            ),
            QueryResult {
                columns: vec!["a".to_owned(), "b".to_owned(), "c".to_owned()],
                rows: vec![Some("a value".to_owned()), Some("12345".to_owned()), None],
            }
        );
    }

    #[test]
    fn query_result_columns() {
        assert_eq!(
            QueryResult::new(
                vec!["a".to_owned(), "b".to_owned(), "c".to_owned()],
                vec![Some("a value".to_owned()), Some("12345".to_owned()), None],
            )
            .columns(),
            vec!["a".to_owned(), "b".to_owned(), "c".to_owned()].as_slice()
        );
    }

    #[test]
    fn query_result_rows() {
        assert_eq!(
            QueryResult::new(
                vec!["a".to_owned(), "b".to_owned(), "c".to_owned()],
                vec![Some("a value".to_owned()), Some("12345".to_owned()), None],
            )
            .rows(),
            vec![Some("a value".to_owned()), Some("12345".to_owned()), None].as_slice()
        );
    }
}
