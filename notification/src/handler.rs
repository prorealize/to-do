//! Implementation of the NotificationService gRPC service.
//!
//! It defines the Notifier struct and its methods, as well as the get_notifier_service function.
//!
//! The Notifier struct implements the NotificationService trait, which is defined in the
//! notification.proto file. The NotificationService trait has a single method, send_notification,
//! which takes a NotificationRequest and returns a NotificationResponse.
//!
//! The get_notifier_service function returns a new instance of the Notifier struct wrapped in a
//! NotificationServiceServer, which is used to serve the gRPC service.
pub mod notification {
    tonic::include_proto!("notification");
}

use notification::{NotificationRequest, NotificationResponse};
use tonic::{Request, Response, Status};
use tracing::info;

use notification::notification_service_server::{NotificationService, NotificationServiceServer};

/// Returns a new instance of the Notifier struct wrapped in a NotificationServiceServer.
///
/// # Returns
///
/// A new instance of NotificationServiceServer<Notifier>.
pub fn get_notifier_service() -> NotificationServiceServer<Notifier> {
    let notifier = Notifier::default();
    NotificationServiceServer::new(notifier)
}

/// The Notifier struct. This struct is used to implement the NotificationService trait.
#[derive(Debug, Default)]
pub struct Notifier {}

#[tonic::async_trait]
/// The implementation of the NotificationService trait for the Notifier struct.
impl NotificationService for Notifier {
    /// Sends a notification.
    ///
    /// This method takes a NotificationRequest and returns a NotificationResponse.
    ///
    /// # Parameters
    ///
    /// * `request` - A Request<NotificationRequest> representing the notification request.
    ///
    /// # Returns
    ///
    /// A Result<Response<NotificationResponse>, Status> representing the notification response.
    async fn send_notification(
        &self,
        request: Request<NotificationRequest>,
    ) -> Result<Response<NotificationResponse>, Status> {
        info!("Request: {:?}", request);
        let reply = NotificationResponse {
            success: true,
            // TODO: Implement notification logic to send message to user
            // and create enum for status with clearer response messages
            status: format!("Notified user with: {}!", request.into_inner().message),
        };
        Ok(Response::new(reply))
    }
}
