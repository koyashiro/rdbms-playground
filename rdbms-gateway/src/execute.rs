use std::fmt::Debug;

use crate::database::Database;

#[async_trait::async_trait]
pub trait Execute<T: Database>: Debug + Send + Sync + 'static {
    async fn execute(&self, query: &str) -> Result<ExecuteResult, T::Error>;
}

#[derive(Clone, Debug, PartialEq)]
pub struct ExecuteResult {
    pub rows: Vec<Vec<Option<Value>>>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum Value {
    Bool(bool),
    I8(i8),
    I16(i16),
    I32(i32),
    I64(i64),
    U8(u8),
    U16(u16),
    U32(u32),
    U64(u64),
    F32(f32),
    F64(f64),
    String(String),
    Bytes(Vec<u8>),
}
