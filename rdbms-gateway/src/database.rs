use std::fmt::Debug;

pub trait Database: Debug + Send + Sync + 'static {
    type Error: std::error::Error;
}
