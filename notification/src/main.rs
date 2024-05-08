//! Main module of the gRPC server application.
//!
//! It uses the `dotenv` crate to load environment variables from a `.env` file.
//! The `main` function is the entry point of the application. It calls the `start_server` function from the `grpc` module to start the gRPC server.
//! The server listens on the host and port defined by the `NOTIFICATION_HOST` and `NOTIFICATION_PORT` environment variables.
//!
//! If the server starts successfully, the process exits with status code 0. If an error occurs, the process exits with status code 1.
//!
//! # Example
//!
//! To start the server, you can run the following command in the terminal:
//!
//! ```shell
//! NOTIFICATION_HOST="127.0.0.1" NOTIFICATION_PORT=3000 cargo run
//! ```
mod grpc;
mod handler;

use dotenv::dotenv;
use std::process;

#[tokio::main]
/// The main function of the application.
///
/// It loads environment variables using `dotenv`, then starts the gRPC server using the `start_server` function from the `grpc` module.
///
/// # Errors
///
/// If the server fails to start, the function will return an error.
///
/// # Returns
///
/// If the server starts successfully, the function will return `Ok(())`. If an error occurs, the function will return `Err(e)`, where `e` is the error.
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok();
    match grpc::start_server().await {
        Ok(_) => process::exit(0),
        Err(_) => {
            process::exit(1);
        }
    }
}
