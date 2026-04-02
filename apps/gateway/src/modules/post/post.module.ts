import { Module } from '@nestjs/common';
import { PostService } from './post.service';
import { PostController } from './post.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { POST_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/post_service/posts';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: GRPC_SERVICE.POST.NAME,
        useFactory: (config: ConfigService<Config>) => ({
          transport: Transport.GRPC,
          options: {
            package: POST_SERVICE_PACKAGE_NAME,
            url: config.get('postServiceEndpoint'),
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
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [PostController],
  providers: [PostService],
  exports: [ClientsModule],
})
export class PostModule {}
