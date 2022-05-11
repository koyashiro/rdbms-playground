use domain::container::{
    entity::Container,
    repository::ContainerRepository,
    value_object::{ContainerImage, MySQLVersion},
};
use infrastructure::docker::ContainerRepositoryImpl;
use shiplift::Docker;

mod application;
mod domain;
mod infrastructure;
mod server;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let docker = Docker::new();

    let repository = ContainerRepositoryImpl::new(docker);
    let container = Container::new(
        "ab3bae32-d70c-46bd-a381-5fb3870c8c0f".to_owned().into(),
        ContainerImage::MySQL(MySQLVersion::V5),
    );
    repository.insert(container).await?;
    let containers = repository.all().await?;
    dbg!(&containers);
    for c in containers {
        repository.delete(c.id().to_owned()).await?;
    }
    let containers = repository.all().await?;
    dbg!(&containers);
    Ok(())
}
