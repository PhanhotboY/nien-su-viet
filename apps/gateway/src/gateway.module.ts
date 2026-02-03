import { LoggerModule } from 'nestjs-pino';
import { APP_FILTER, APP_GUARD, APP_INTERCEPTOR } from '@nestjs/core';
import { CacheInterceptor } from '@nestjs/cache-manager';
import { ClassSerializerInterceptor, Module, Scope } from '@nestjs/common';
import { ThrottlerGuard, ThrottlerModule } from '@nestjs/throttler';

import { configuration, loggerOptions, RATE_LIMIT } from './config';
import { AuthModule } from './modules/auth/auth.module';
import { CmsModule } from './modules/cms/cms.module';
import { HistoricalEventModule } from './modules/historical-event/historical-event.module';
import { CommonModule } from '@phanhotboy/nsv-common';
import { SerializeResponseInterceptor } from './common/interceptors';
import { CatchEverythingFilter, HttpExceptionsFilter } from './common/filters';
import { BetterAuthGuard, RolesGuard } from './common/guards';
import { MicroserviceExceptionFilter } from '@phanhotboy/nsv-common/filters/rpc-exception.filter';

@Module({
  imports: [
    LoggerModule.forRoot(loggerOptions),
    CommonModule.forRoot({
      cachePrefix: 'gateway',
      configuration,
    }),
    ThrottlerModule.forRoot(Object.values(RATE_LIMIT.DEFAULT)),
    AuthModule,
    CmsModule,
    HistoricalEventModule,
  ],
  providers: [
    // Global guards
    {
      provide: APP_GUARD,
      useClass: ThrottlerGuard,
    },
    {
      provide: APP_GUARD,
      useClass: BetterAuthGuard,
    },
    {
      provide: APP_GUARD,
      useClass: RolesGuard,
    },
    // Global interceptors
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
    // Global pipes
    // Global filters
    {
      provide: APP_FILTER,
      useClass: CatchEverythingFilter,
    },
    {
      provide: APP_FILTER,
      useClass: HttpExceptionsFilter,
    },
    {
      provide: APP_FILTER,
      useClass: MicroserviceExceptionFilter,
    },
  ],
})
export class GatewayModule {}
