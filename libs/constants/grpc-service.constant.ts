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
} as const;
