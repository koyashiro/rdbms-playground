use std::sync::Arc;

use rdbms_gateway::{
    postgres::PgExecute, proto::rdbms_gateway_service_server::RdbmsGatewayServiceServer,
    server::RdbmsGatewayService,
};
use sqlx::postgres::PgPoolOptions;
use tonic::transport::Server;

#[tokio::main]
async fn main() {
    let pool = PgPoolOptions::new()
        .connect("postgres://postgres:password@localhost/postgres")
        .await
        .unwrap();
    let execute = Arc::new(PgExecute::new(pool));
    let service = RdbmsGatewayService::new(execute);
    let server = RdbmsGatewayServiceServer::new(service);
    Server::builder()
        .add_service(server)
        .serve("[::1]:50051".parse().unwrap())
        .await
        .unwrap();
}
