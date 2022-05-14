mod execute;

use std::sync::Arc;

use rdbms_gateway::{
    proto::rdbms_gateway_service_server::RdbmsGatewayServiceServer, server::RdbmsGatewayService,
};
use sqlx::mysql::MySqlPoolOptions;
use tonic::transport::Server;

use crate::execute::MySqlExecute;

#[tokio::main]
async fn main() {
    let pool = MySqlPoolOptions::new()
        .connect("mysql://root:password@localhost/mysql")
        .await
        .unwrap();
    let execute = Arc::new(MySqlExecute::new(pool));
    let service = RdbmsGatewayService::new(execute);
    let server = RdbmsGatewayServiceServer::new(service);
    Server::builder()
        .add_service(server)
        .serve("[::1]:50051".parse().unwrap())
        .await
        .unwrap();
}
