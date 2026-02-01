import { Module } from '@nestjs/common';
import { LoggerModule } from 'nestjs-pino';
import { CommonModule } from '@phanhotboy/nsv-common';

import { configuration } from './config/configuration';
import { AuthModule } from './auth';
import { PrismaModule } from './database';
import { loggerOptions } from './config';

@Module({
  imports: [
    LoggerModule.forRoot(loggerOptions  ),
    CommonModule.forRoot({
      configuration,
      cachePrefix: 'auth-service',
      global: true,
    }),
    PrismaModule.forRoot(),
    AuthModule,
  ],
})
export class AppModule {}
