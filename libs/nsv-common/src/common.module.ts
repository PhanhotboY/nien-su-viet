import { DynamicModule, Global, Module, ValidationPipe } from '@nestjs/common';
import { CacheModule } from '@nestjs/cache-manager';

import * as providers from './providers';
import { ConfigModule } from '@nestjs/config';
import { APP_PIPE } from '@nestjs/core';
import {
  ConfigurableModuleClass,
  MODULE_OPTIONS_TOKEN,
  OPTIONS_TYPE,
} from './common.module-definition';
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
            ttl: 30 * 1000, // 30 secs, short caching time for gateway route auto caching
            max: 100, // Maximum number of items in cache
          }),
          isGlobal: true,
          inject: [providers.ConfigService],
        }),
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
      ],
      exports: [...services, ConfigModule],
    };
  }
}
