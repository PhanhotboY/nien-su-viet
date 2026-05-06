import { Module } from '@nestjs/common';
import { SearchPostService } from './post.service';
import { SearchPostController } from './post.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { SEARCH_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/search_service/historical_events';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';
import { GatewayMetrics } from '@gateway/common/contracts';
import { SearchHistoricalEventController } from './historical-event.controller';
import { SearchHistoricalEventService } from './historical-event.service';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: GRPC_SERVICE.SEARCH.NAME,
        useFactory: (config: ConfigService<Config>) => ({
          transport: Transport.GRPC,
          options: {
            package: SEARCH_SERVICE_PACKAGE_NAME,
            url: config.get('searchServiceEndpoint'),
            protoPath: GRPC_SERVICE.SEARCH.PROTO_PATH,
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
  controllers: [SearchPostController, SearchHistoricalEventController],
  providers: [SearchPostService, GatewayMetrics, SearchHistoricalEventService],
  exports: [ClientsModule],
})
export class SearchModule {}
