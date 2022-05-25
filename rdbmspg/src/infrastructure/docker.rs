use docker_api::{api::ContainerCreateOpts, Docker};

use crate::domain::{
    docker::{Container, DockerClient},
    rdbms::Rdbms,
};

pub struct DockerClientImpl {
    docker: Docker,
}

impl DockerClientImpl {
    pub fn new(docker: Docker) -> Self {
        Self { docker }
    }
}

#[async_trait::async_trait]
impl DockerClient for DockerClientImpl {
    async fn create_and_start(&self, rdbms: &Rdbms) -> anyhow::Result<Container> {
        let image = rdbms.to_string();
        let env = match rdbms {
            Rdbms::MySQL(_) => ["MYSQL_ADMIN_PASSWORD=password"],
            Rdbms::MariaDB(_) => ["MARIADB_ADMIN_PASSWORD=password"],
            Rdbms::PostgreSQL(_) => ["POSTGRES_PASSWORD=password"],
        };
        let opts = ContainerCreateOpts::builder(image)
            .auto_remove(true)
            .env(env)
            .publish_all_ports()
            .build();
        let container = self.docker.containers().create(&opts).await?;
        container.start().await?;
        let container_details = container.inspect().await?;
        Ok(container_details.into())
    }

    async fn stop_and_delete(&self, container_id: &str) -> anyhow::Result<()> {
        self.docker
            .containers()
            .get(container_id)
            .stop(None)
            .await?;
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use crate::domain::rdbms::{MariaDBVersion, MySQLVersion, PostgreSQLVersion};

    use super::*;

    #[tokio::test]
    async fn docker_client_impl() {
        let docker = Docker::new("unix:///var/run/docker.sock").unwrap();
        let docker_client = DockerClientImpl::new(docker);
        let rdbmss = [
            Rdbms::MySQL(MySQLVersion::V8_0),
            Rdbms::MariaDB(MariaDBVersion::V10_8),
            Rdbms::PostgreSQL(PostgreSQLVersion::V14),
        ];
        for rdbms in rdbmss {
            let container = docker_client.create_and_start(&rdbms).await.unwrap();
            docker_client.stop_and_delete(container.id()).await.unwrap();
        }
    }
}
