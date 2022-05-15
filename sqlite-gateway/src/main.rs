use std::sync::Arc;

use rdbms_gateway::{
    proto::rdbms_gateway_service_server::RdbmsGatewayServiceServer, server::RdbmsGatewayService,
    sqlite::SqliteExecute,
};
use sqlx::sqlite::SqlitePoolOptions;
use tonic::transport::Server;

#[tokio::main]
async fn main() {
    let pool = SqlitePoolOptions::new()
        .connect("sqlite::memory:")
        .await
        .unwrap();
    let execute = Arc::new(SqliteExecute::new(pool));
    let service = RdbmsGatewayService::new(execute);
    let server = RdbmsGatewayServiceServer::new(service);
    Server::builder()
        .add_service(server)
        .serve("[::1]:50051".parse().unwrap())
        .await
        .unwrap();
}
