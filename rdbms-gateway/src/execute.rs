use std::fmt::Debug;

use crate::database::{Database, Value};

#[derive(Clone, Debug, PartialEq)]
pub struct ExecuteResult {
    pub rows: Vec<Vec<Option<Value>>>,
}

#[async_trait::async_trait]
pub trait Execute<T: Database>: Debug + Send + Sync + 'static {
    async fn execute(&self, query: &str) -> Result<ExecuteResult, T::Error>;
}
