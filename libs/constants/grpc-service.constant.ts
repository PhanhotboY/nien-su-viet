export const GRPC_SERVICE = {
  MAIN_PROTO_PATH: './api/proto',
  POST: {
    NAME: 'POST_SERVICE',
    URL: '0.0.0.0:6005',
    PROTO_PATH: './api/proto/post_service/posts.proto',
  },
} as const;
