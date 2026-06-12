// proto path must be relative to the project root
export const GRPC_SERVICE = {
  MAIN_PROTO_PATH: './api/proto',
  POST: {
    NAME: 'POST_SERVICE',
    PROTO_PATH: './api/proto/post_service/posts.proto',
  },
  HISTORICAL_EVENT: {
    NAME: 'HISTORICAL_EVENT_SERVICE',
    PROTO_PATH: './api/proto/historical_event_service/historical_events.proto',
  },
  BILLING: {
    NAME: 'BILLING_SERVICE',
    PROTO_PATH: [
      './api/proto/billing_service/billing.proto',
      './api/proto/billing_service/plans.proto',
      './api/proto/billing_service/subscriptions.proto',
      './api/proto/billing_service/subscription_events.proto',
      './api/proto/billing_service/payment_attempts.proto',
      './api/proto/billing_service/payment_transactions.proto',
      './api/proto/billing_service/outbox_events.proto',
      './api/proto/billing_service/processed_events.proto',
      './api/proto/billing_service/purchases.proto',
      './api/proto/billing_service/webhook.proto',
    ] as string[],
  },
} as const;
