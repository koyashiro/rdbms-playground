use std::fmt::Debug;

pub trait Database: Debug + Send + Sync + 'static {
    type Error: std::error::Error;
}

#[derive(Clone, Debug, PartialEq)]
pub enum Value {
    Bool(bool),
    I32(i32),
    I64(i64),
    F32(f32),
    F64(f64),
    String(String),
    Bytes(Vec<u8>),
}
