import { Module } from '@nestjs/common';
import { CommonModule } from '@phanhotboy/nsv-common';

import { configuration } from './config/configuration';
import { AuthModule } from './auth';
import { PrismaModule } from './database';

@Module({
  imports: [
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
