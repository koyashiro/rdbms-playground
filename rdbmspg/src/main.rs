mod domain;
mod infrastructure;
mod presentation;
mod usecase;

use std::net::SocketAddr;

use axum::{routing::post, Router, Server};

#[tokio::main]
async fn main() {
    let ip = "0.0.0.0".parse().unwrap();
    let port = "3000".parse().unwrap();
    let addr = SocketAddr::new(ip, port);

    let app = Router::new().route("/", post(|| async { "todo!" }));

    Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}
