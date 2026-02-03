import { Module } from '@nestjs/common';
import { HistoricalEventService } from './historical-event.service';
import { HistoricalEventController } from './historical-event.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { TCP_SERVICE } from '@phanhotboy/constants/tcp-service.constant';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@gateway/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: TCP_SERVICE.HISTORICAL_EVENT.NAME,
        useFactory: (config: ConfigService<Config>) => ({
          transport: Transport.TCP,
          options: {
            host: config.get('historicalEventServiceHost'),
            port: TCP_SERVICE.HISTORICAL_EVENT.PORT,
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
