import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';

import { Config } from '@gateway/config';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { ConfigService } from '@phanhotboy/nsv-common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { HISTORICAL_EVENT_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/historical_event_service/historical_events';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: GRPC_SERVICE.HISTORICAL_EVENT.NAME,
        useFactory: (config: ConfigService<Config>) => ({
          transport: Transport.GRPC,
          options: {
            url: config.get('historicalEventServiceEndpoint'),
            package: HISTORICAL_EVENT_SERVICE_PACKAGE_NAME,
            protoPath: GRPC_SERVICE.HISTORICAL_EVENT.PROTO_PATH,
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
  controllers: [HistoricalEventController],
  providers: [HistoricalEventService],
  exports: [ClientsModule],
})
export class HistoricalEventModule {}
