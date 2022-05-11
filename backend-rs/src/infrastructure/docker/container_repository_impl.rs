use std::{collections::HashMap, time::Duration};

use async_trait::async_trait;
use http::StatusCode;
use shiplift::{
    rep::Container as ContainerRep, ContainerFilter, ContainerListOptions, ContainerOptions,
    Docker, PullOptions,
};

use crate::domain::container::{
    self,
    entity::Container,
    repository::{ContainerError, ContainerRepository},
    value_object::{ContainerId, ContainerImage},
};

pub struct ContainerRepositoryImpl {
    docker: Docker,
}

impl ContainerRepositoryImpl {
    pub fn new(docker: Docker) -> Self {
        Self { docker }
    }
}

impl From<ContainerRep> for Container {
    fn from(c: shiplift::rep::Container) -> Self {
        Self::new(c.id.into(), c.image.parse().unwrap())
    }
}

#[async_trait]
impl ContainerRepository for ContainerRepositoryImpl {
    async fn find(&self, id: ContainerId) -> Result<Option<Container>, ContainerError> {
        let container = find(&self.docker, id.into()).await?;
        Ok(container.map(|c| c.into()))
    }

    async fn all(&self) -> Result<Vec<Container>, ContainerError> {
        let opts = ContainerListOptions::builder()
            .filter(vec![ContainerFilter::LabelName("rdbmsp".to_owned())])
            .build();
        let containers = self
            .docker
            .containers()
            .list(&opts)
            .await?
            .into_iter()
            .map(|c| c.into())
            .collect();
        Ok(containers)
    }

    async fn insert(&self, container: Container) -> Result<(), ContainerError> {
        let name = container.image().to_string();

        {
            if let Err(shiplift::Error::Fault {
                code: StatusCode::NOT_FOUND,
                message: _,
            }) = self.docker.images().get(&name).inspect().await
            {
                use futures::StreamExt;

                let opts = PullOptions::builder().image(&name).build();
                let mut stream = self.docker.images().pull(&opts);
                while let Some(result) = stream.next().await {
                    result?;
                }
            }
        }

        let envs = {
            match container.image() {
                ContainerImage::MySQL(_) => &["MYSQL_ROOT_PASSWORD=password"],
                ContainerImage::MariaDB(_) => &["MARIADB_ROOT_PASSWORD=password"],
                ContainerImage::PostgreSQL(_) => &["POSTGRES_PASSWORD=password"],
            }
        };
        let labels = HashMap::from([("rdbmsp", "true")]);

        let opts = ContainerOptions::builder(&name)
            .name(container.id().as_str())
            .labels(&labels)
            .env(envs)
            .auto_remove(true)
            .build();
        let info = self.docker.containers().create(&opts).await?;

        let container = self.docker.containers().get(info.id);
        container.start().await?;

        Ok(())
    }

    async fn delete(&self, id: ContainerId) -> Result<(), ContainerError> {
        // let container = find(&self.docker, id.into()).await?;
        // if let Some(c) = container {
        //     c.stop(Some(Duration::from_secs(10))).await?;
        // }
        let container = self.docker.containers().get(id);
        match container.stop(None).await {
            Err(shiplift::Error::Fault {
                code: StatusCode::NOT_FOUND,
                message: _,
            }) => Ok(()),
            Err(e) => Err(e),
            Ok(_) => Ok(()),
        }
    }
}

async fn find(
    docker: &Docker,
    name: String,
) -> Result<Option<shiplift::Container<'_>>, ContainerError> {
    let opts = ContainerListOptions::builder()
        .filter(vec![
            ContainerFilter::LabelName("rdbmsp".to_owned()),
            ContainerFilter::Label("name".to_owned(), name),
        ])
        .build();
    let mut containers = docker.containers().list(&opts).await?;
    let contaienr = if containers.is_empty() {
        None
    } else {
        let cr = containers.swap_remove(0);
        let c = docker.containers().get(c.id);
        (cr, c)
    };
    Ok(contaienr)
}
