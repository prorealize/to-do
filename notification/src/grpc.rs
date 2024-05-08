//! Responsible for the gRPC server side.
//!
//! It defines the `start_server` function which starts the gRPC server based on environment variables for host and port.
//! It also defines the `get_address` function which retrieves the host and port from environment variables.
//!
//! The `start_server` function uses the `get_notifier_service` function from the `handler` module to get the gRPC service.
//! It then adds this service to the server and starts the server at the specified address.
//!
//! The `get_address` function retrieves the host and port from the `NOTIFICATION_HOST` and `NOTIFICATION_PORT` environment variables respectively.
//! It returns a `SocketAddr` which is used by the `start_server` function to start the server.
//!
//! This module also defines a custom `ServerError` enum which represents the possible errors that can occur when getting the address.
//! This enum is used by the `get_address` function to return a custom error when the environment variables are not set or when the host or port is not valid.
use std::error::Error;
use std::net::SocketAddr;
use std::{env, fmt};

use tonic::transport::Server;
use tracing::{error, info};

use crate::handler::get_notifier_service;

/// Start gRPC server based on environment variables for host and port
///
/// # Errors
///
/// Returns an error if the server cannot be started
///
/// # Returns
///
/// A Result indicating whether the server was started successfully
///
/// # Example
///
/// ```rust
/// let result = start_server().await;
/// ```
pub async fn start_server() -> Result<(), Box<dyn Error>> {
    let addr = match get_address() {
        Ok(addr) => addr,
        Err(err) => {
            error!("{}", err);
            return Err(err);
        }
    };

    info!("Starting server at {}", addr);
    let notification_service_server = get_notifier_service();
    let result = Server::builder()
        .add_service(notification_service_server)
        .serve(addr)
        .await;
    match result {
        Ok(_) => Ok(()),
        Err(err) => {
            error!("Unable to start server due to {}", err);
            return Err(Box::new(err));
        }
    }
}

/// Get host and port from environment variables
///
/// # Errors
///
/// Returns an error if the host or port is missing or invalid
///
/// # Returns
///
/// A valid SocketAddr
///
/// # Example
///
/// ```rust
/// let addr = get_address().unwrap();
/// ```
fn get_address() -> Result<SocketAddr, Box<dyn Error>> {
    let host = env::var("NOTIFICATION_HOST")
        .map_err(|_| ServerError::MissingHost)?
        .parse()
        .map_err(|_| ServerError::InvalidHost)?;
    let port = env::var("NOTIFICATION_PORT")
        .map_err(|_| ServerError::MissingPort)?
        .parse()
        .map_err(|_| ServerError::InvalidPort)?;

    Ok(SocketAddr::new(host, port))
}

#[derive(Debug)]
pub enum ServerError {
    MissingHost,
    MissingPort,
    InvalidHost,
    InvalidPort,
}

impl fmt::Display for ServerError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match *self {
            ServerError::MissingHost => write!(f, "NOTIFICATION_HOST must be set"),
            ServerError::MissingPort => write!(f, "NOTIFICATION_PORT must be set"),
            ServerError::InvalidHost => write!(f, "Host must be a valid IP address"),
            ServerError::InvalidPort => write!(f, "Port must be an unassigned integer"),
        }
    }
}

impl Error for ServerError {}
