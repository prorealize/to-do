mod handler;
mod grpc;

use dotenv::dotenv;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok();
    grpc::start_server().await;
    Ok(())
}
