use super::super::{value_object::ContainerId, value_object::ContainerImage};

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct Container {
    id: ContainerId,
    image: ContainerImage,
}

impl Container {
    pub fn new(id: ContainerId, image: ContainerImage) -> Self {
        Self { id, image }
    }

    pub fn id(&self) -> &ContainerId {
        &self.id
    }

    pub fn image(&self) -> &ContainerImage {
        &self.image
    }
}
