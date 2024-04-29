use std::env;
use std::net::SocketAddr;

use tonic::transport::Server;
use tracing::info;

use crate::handler::get_notifier_service;


pub async fn start_server() {
    let host = env::var("NOTIFICATION_HOST")
        .expect("NOTIFICATION_HOST must be set");
    let port = env::var("NOTIFICATION_PORT")
        .expect("NOTIFICATION_PORT must be set")
        .parse().expect("Port must be an unassigned integer");

    let host = host.parse().expect("Invalid host");
    let addr = SocketAddr::new(host, port);

    let notification_service_server = get_notifier_service();

    info!("Starting server at {host}:{port}");

    let result = Server::builder()
        .add_service(notification_service_server)
        .serve(addr)
        .await;
    match result {
        Ok(_) => {}
        Err(err) => { panic!("Unable to start server due to {}", err)}
    }
}