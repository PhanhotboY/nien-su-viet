import { Module } from '@nestjs/common';
import { PostService } from './post.service';
import { PostController } from './post.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { GRPC_SERVICE, RMQ } from '@phanhotboy/constants';
import { POST_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/post_service/posts';
import { RmqModule } from '@phanhotboy/nsv-common';

@Module({
  imports: [
    RmqModule.register({ name: RMQ.TOPIC_EVENTS_EXCHANGE }),
    ClientsModule.registerAsync([
      {
        name: GRPC_SERVICE.POST.NAME,
        useFactory: () => ({
          transport: Transport.GRPC,
          options: {
            package: POST_SERVICE_PACKAGE_NAME,
            url: GRPC_SERVICE.POST.URL,
            protoPath: GRPC_SERVICE.POST.PROTO_PATH,
            // Ref: https://github.com/grpc/grpc-node/blob/master/packages/proto-loader/README.md
            loader: {
              // includes imported proto files
              includeDirs: [GRPC_SERVICE.MAIN_PROTO_PATH],
              // since Unmarshal removes falsy values, we need to set default values back to avoid losing data, like empty array, false, etc.
              defaults: true,
              // The type to use to represent long (int64) values. Instead of Long object by default.
              longs: Number,
            },
          },
        }),
        inject: [],
      },
    ]),
  ],
  controllers: [PostController],
  providers: [PostService],
  exports: [ClientsModule],
})
export class PostModule {}
