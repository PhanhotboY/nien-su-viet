export const GRPC_SERVICE = {
  MAIN_PROTO_PATH: './api/proto',
  POST: {
    NAME: 'POST_SERVICE',
    PROTO_PATH: './api/proto/post_service/posts.proto',
  },
  SEARCH: {
    NAME: 'SEARCH_SERVICE',
    PROTO_PATH: [
      './api/proto/search_service/posts.proto',
      './api/proto/search_service/historical_events.proto',
      './api/proto/search_service/users.proto',
    ] as string[],
  },
  HISTORICAL_EVENT: {
    NAME: 'HISTORICAL_EVENT_SERVICE',
    PROTO_PATH: './api/proto/historical_event_service/historical_events.proto',
  },
} as const;
