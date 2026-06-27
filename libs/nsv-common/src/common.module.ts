import {
  DynamicModule,
  Global,
  Logger,
  Module,
  ValidationPipe,
} from '@nestjs/common';
import { CacheModule } from '@nestjs/cache-manager';

import * as providers from './providers';
import { ConfigModule } from '@nestjs/config';
import { APP_INTERCEPTOR, APP_PIPE } from '@nestjs/core';
import {
  ConfigurableModuleClass,
  MODULE_OPTIONS_TOKEN,
  OPTIONS_TYPE,
} from './common.module-definition';
import KeyvRedis from '@keyv/redis';
import { LoggerModule } from './logger';
import { MeterModule } from './otel';
import { MetricsInterceptor, TracingInterceptor } from './interceptors';
import { RetrySystem } from './util/retry.util';

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
        LoggerModule,
        ConfigModule.forRoot({
          load: [options.configuration],
          cache: true,
        }),
        CacheModule.registerAsync({
          useFactory: async (config: providers.ConfigService) => ({
            stores: [new KeyvRedis(config.get('redis.url'))],
            ttl: config.get('NODE_ENV') === 'development' ? -1 : 30 * 1000, // 30 secs, short caching time for gateway route auto caching
            max: 100, // Maximum number of items in cache
          }),
          isGlobal: true,
          inject: [providers.ConfigService],
        }),
        MeterModule,
      ],
      providers: [
        { provide: MODULE_OPTIONS_TOKEN, useValue: options },
        ...services,
        {
          provide: APP_PIPE,
          useValue: new ValidationPipe({
            // disableErrorMessages: true,
            transform: true, // transform object to DTO class
            whitelist: true,
            forbidNonWhitelisted: true,
          }),
        },
        {
          provide: APP_INTERCEPTOR,
          useClass: MetricsInterceptor,
        },
        {
          provide: APP_INTERCEPTOR,
          useClass: TracingInterceptor,
        },
        {
          provide: RetrySystem,
          useFactory: (logger: Logger) => {
            return new RetrySystem({
              logger,
              retry: {
                maxDelay: 30000,
                maxRetries: 5,
                baseDelay: 1000,
                jitter: true,
              },
              circuitBreaker: {
                failureThreshold: 5,
                resetTimeout: 60000, // 1 minute
              },
            });
          },
          inject: [Logger],
        },
      ],
      exports: [
        ...services,
        ConfigModule,
        LoggerModule,
        MeterModule,
        RetrySystem,
      ],
    };
  }
}
