# gRPC Notification Service

This application is a gRPC server that provides a Notification Service. The service is defined in the `notification.proto` file and implemented in the `handler` module. The server listens on the host and port defined by the `NOTIFICATION_HOST` and `NOTIFICATION_PORT` environment variables.

## Running the Application

You can run the application using Docker and Docker Compose. The `docker-compose.yml` file is configured to start the gRPC server with a specified number of replicas.

To start the application with Docker Compose, run the following command:
```bash
docker command up
```
Alternatively, you can run the application with Docker:
```bash
cargo run --release
```
## Test

```bash
# Make sure the server is running first
apt install grpcurl
grpcurl -plaintext -d '{"message": "Hello, World!"}' <host>:<port> notification.NotificationService/SendNotification
```