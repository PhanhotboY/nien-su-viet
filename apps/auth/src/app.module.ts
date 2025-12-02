import { Module } from '@nestjs/common';
import { CommonModule, ConfigService } from '@phanhotboy/nsv-common';

import { configuration } from './config/configuration';
import { AuthModule } from './auth';
import { PrismaModule } from './database';

@Module({
  imports: [
    CommonModule.forRoot({
      configuration,
      rabbitmqConfigKey: 'rabbitmq',
      redisConfigKey: 'redis',
      throttlerConfigKey: 'throttlers',
      cachePrefix: 'auth-service',
      global: true,
      useInterceptors: false,
    }),
    PrismaModule.forRoot(),
    AuthModule,
  ],
  providers: [ConfigService],
})
export class AppModule {}
