pub mod notification {
    tonic::include_proto!("notification");
}

use notification::{NotificationRequest, NotificationResponse};
use tonic::{Request, Response, Status};

use notification::notification_service_server::{NotificationService, NotificationServiceServer};

pub fn get_notifier_service() -> NotificationServiceServer<Notifier> {
    let notifier = Notifier::default();
    NotificationServiceServer::new(notifier)
}

#[derive(Debug, Default)]
pub struct Notifier {}

#[tonic::async_trait]
impl NotificationService for Notifier {
    async fn send_notification(
        &self,
        request: Request<NotificationRequest>,
    ) -> Result<Response<NotificationResponse>, Status> {
        println!("Got a request: {:?}", request);

        let reply = notification::NotificationResponse {
            success: true,
            status: format!("Hello {}!", request.into_inner().message),
        };

        Ok(Response::new(reply))
    }
}
