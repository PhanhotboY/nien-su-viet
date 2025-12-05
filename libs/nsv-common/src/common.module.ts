import {
  ClassSerializerInterceptor,
  DynamicModule,
  Global,
  Module,
  Scope,
  ValidationPipe,
} from '@nestjs/common';
import { CacheInterceptor, CacheModule } from '@nestjs/cache-manager';

import * as providers from './providers';
import { ConfigModule } from '@nestjs/config';
import { ThrottlerModule } from '@nestjs/throttler';
import { APP_FILTER, APP_GUARD, APP_INTERCEPTOR, APP_PIPE } from '@nestjs/core';
import { HttpExceptionsFilter } from './filters';
import { SerializeResponseInterceptor } from './interceptors';
import { LoggerModule } from 'nestjs-pino';
import { loggerOptions } from './config/logger.config';
import {
  ConfigurableModuleClass,
  MODULE_OPTIONS_TOKEN,
  OPTIONS_TYPE,
} from './common.module-definition';
import { BetterAuthGuard, RolesGuard } from './auth';
import { createKeyv, RedisClientOptions } from '@keyv/redis';

const { ...prvds } = providers;
const services = Object.values(prvds);

@Global()
@Module({})
export class CommonModule extends ConfigurableModuleClass {
  static forRoot(options: typeof OPTIONS_TYPE): DynamicModule {
    return {
      ...super.forRoot(options),
      module: CommonModule,
      imports: [
        LoggerModule.forRoot(loggerOptions),
        ConfigModule.forRoot({
          load: [options.configuration],
          cache: true,
        }),
        CacheModule.registerAsync({
          useFactory: async (config: providers.ConfigService) => ({
            stores: [
              createKeyv({
                url: config.get('REDIS_URL'),
              } as RedisClientOptions),
            ],
            ttl: 60 * 5 * 1000, // 5 minutes in seconds for NestJS
            max: 100, // Maximum number of items in cache
          }),
          isGlobal: true,
          inject: [providers.ConfigService],
        }),
        ThrottlerModule.forRootAsync({
          useFactory: (config: providers.ConfigService) => ({
            throttlers: config.get(options.throttlerConfigKey),
          }),
          inject: [providers.ConfigService],
        }),
      ],
      providers: [
        { provide: MODULE_OPTIONS_TOKEN, useValue: options },
        ...services,
        ...(options.useCacheInterceptor
          ? [
              {
                provide: APP_INTERCEPTOR,
                useClass: CacheInterceptor,
              },
            ]
          : []),
        { provide: APP_FILTER, useClass: HttpExceptionsFilter },
        ...(options.useSerializeInterceptors
          ? [
              {
                provide: APP_INTERCEPTOR,
                useClass: ClassSerializerInterceptor,
                scope: Scope.REQUEST,
              },
              {
                provide: APP_INTERCEPTOR,
                useClass: SerializeResponseInterceptor,
              },
            ]
          : []),
        { provide: APP_GUARD, useClass: BetterAuthGuard },
        { provide: APP_GUARD, useClass: RolesGuard },
        {
          provide: APP_PIPE,
          useValue: new ValidationPipe({
            // disableErrorMessages: true,
            transform: true, // transform object to DTO class
            whitelist: true,
          }),
        },
      ],
      exports: [...services, ConfigModule],
    };
  }
}
