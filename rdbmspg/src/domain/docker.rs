use std::net::Ipv4Addr;

use docker_api::api::ContainerDetails;

use super::rdbms::Rdbms;

#[derive(Debug)]
pub struct ImageParseError;

pub struct Image(String);

impl Image {
    pub fn name(&self) -> &str {
        self.0.split(':').collect::<Vec<&str>>()[0]
    }

    pub fn version(&self) -> &str {
        self.0.split(':').collect::<Vec<&str>>()[1]
    }
}

impl TryFrom<String> for Image {
    type Error = ImageParseError;

    fn try_from(value: String) -> Result<Self, Self::Error> {
        if value.split(':').count() != 2 {
            return Err(ImageParseError);
        }
        Ok(Image(value))
    }
}

pub struct Container {
    id: String,
    ip: Ipv4Addr,
    port: u16,
    image: Image,
}

impl Container {
    pub fn id(&self) -> &str {
        &self.id
    }

    pub fn ip(&self) -> &Ipv4Addr {
        &self.ip
    }

    pub fn port(&self) -> &u16 {
        &self.port
    }

    pub fn image(&self) -> &Image {
        &self.image
    }
}

impl From<ContainerDetails> for Container {
    fn from(cd: ContainerDetails) -> Self {
        let id = cd.id;
        let ip = [0, 0, 0, 0].into();
        let image: Image = cd.config.unwrap().image.try_into().unwrap();
        let mut port_map = cd.network_settings.ports.unwrap();
        let port = match image.name() {
            "mysql" => port_map.remove("3306/tcp"),
            "mariadb" => port_map.remove("3306/tcp"),
            "postgres" => port_map.remove("5432/tcp"),
            _ => unreachable!(),
        }
        .unwrap()
        .unwrap()
        .into_iter()
        .find(|p| p.host_ip == "0.0.0.0")
        .unwrap()
        .host_port
        .parse()
        .unwrap();
        Self {
            id,
            ip,
            port,
            image,
        }
    }
}

#[async_trait::async_trait]
pub trait DockerClient {
    async fn create_and_start(&self, rdbms: &Rdbms) -> anyhow::Result<Container>;
    async fn stop_and_delete(&self, container_id: &str) -> anyhow::Result<()>;
}
