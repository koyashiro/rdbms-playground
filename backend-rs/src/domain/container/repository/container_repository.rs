use async_trait::async_trait;

use super::super::{entity::Container, value_object::ContainerId};

pub type ContainerError = anyhow::Error;

#[async_trait]
pub trait ContainerRepository {
    async fn find(&self, id: ContainerId) -> Result<Option<Container>, ContainerError>;
    async fn all(&self) -> Result<Vec<Container>, ContainerError>;
    async fn insert(&self, playground: Container) -> Result<(), ContainerError>;
    async fn delete(&self, id: ContainerId) -> Result<(), ContainerError>;
}
