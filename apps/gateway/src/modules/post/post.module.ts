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
            loader: {
              includeDirs: [GRPC_SERVICE.MAIN_PROTO_PATH],
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
