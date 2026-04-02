import { Module } from '@nestjs/common';
import { CommonModule } from '@phanhotboy/nsv-common';
import { configuration } from './config';
import { HistoricalEventModule } from './modules/historical-event/historical-event.module';
import { PrismaModule } from './database';
import { UserModule } from './modules/user';

@Module({
  imports: [
    CommonModule.forRoot({
      cachePrefix: 'historical-event-service',
      configuration,
      global: true,
    }),
    PrismaModule.forRoot(),
    HistoricalEventModule,
    UserModule,
  ],
})
export class AppModule {}
