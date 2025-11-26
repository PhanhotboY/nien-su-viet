import { Module } from '@nestjs/common';
import { APP_GUARD } from '@nestjs/core';
import {
  BetterAuthGuard,
  CommonModule,
  ConfigService,
  RolesGuard,
} from '@phanhotboy/nsv-common';

import { UserModule } from './modules/user';
import { PrismaModule } from './database/prisma.module';
import { Config } from './config';
import { configuration } from './config/configuration';

@Module({
  imports: [
    CommonModule.forRoot({
      global: true,
      configuration,
      redisConfigKey: 'redis',
      throttlerConfigKey: 'throttlers',
      rabbitmqConfigKey: 'rabbitmq',
      cachePrefix: 'user-service',
    }),
    PrismaModule.forRoot(),
    UserModule,
  ],
  providers: [
    {
      provide: APP_GUARD,
      useClass: BetterAuthGuard,
    },
    {
      provide: APP_GUARD,
      useClass: RolesGuard,
    },
  ],
})
export class AppModule {}
