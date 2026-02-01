import { LoggerModule } from 'nestjs-pino';
import { APP_FILTER, APP_GUARD, APP_INTERCEPTOR } from '@nestjs/core';
import { CacheInterceptor } from '@nestjs/cache-manager';
import { ClassSerializerInterceptor, Module, Scope } from '@nestjs/common';

import { configuration, loggerOptions, RATE_LIMIT } from './config';
import { AuthModule } from './modules/auth/auth.module';
import { CmsModule } from './modules/cms/cms.module';
import { HistoricalEventModule } from './modules/historical-event/historical-event.module';
import { CommonModule } from '@phanhotboy/nsv-common';
import { SerializeResponseInterceptor } from './common/interceptors';
import { ThrottlerGuard, ThrottlerModule } from '@nestjs/throttler';
import { CatchEverythingFilter, HttpExceptionsFilter } from './common/filters';

@Module({
  imports: [
    LoggerModule.forRoot(loggerOptions),
    CommonModule.forRoot({
      cachePrefix: 'gateway',
      configuration,
    }),
    ThrottlerModule.forRoot(RATE_LIMIT.DEFAULT),
    AuthModule,
    // CmsModule,
    HistoricalEventModule,
  ],
  providers: [
    {
      provide: APP_INTERCEPTOR,
      useClass: ClassSerializerInterceptor,
      scope: Scope.REQUEST,
    },
    {
      provide: APP_INTERCEPTOR,
      useClass: SerializeResponseInterceptor,
    },
    {
      provide: APP_INTERCEPTOR,
      useClass: CacheInterceptor,
    },
    {
      provide: APP_GUARD,
      useClass: ThrottlerGuard,
    },
    {
      provide: APP_FILTER,
      useClass: CatchEverythingFilter,
    },
    {
      provide: APP_FILTER,
      useClass: HttpExceptionsFilter,
    },
  ],
})
export class GatewayModule {}
