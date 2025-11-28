import { RedisClientOptions } from '@keyv/redis';
import { ThrottlerOptions } from '@nestjs/throttler';
import { Options } from 'amqplib';

export const config = {
  db: {
    url: process.env.DATABASE_URL || '',
    directUrl: process.env.DIRECT_DATABASE_URL || '',
    //   type: process.env.DB_TYPE || 'mysql',
    //   synchronize: false,
    //   logging: true,
    //   host: process.env.DB_HOST || '127.0.0.1',
    //   port: process.env.DB_PORT || 3306,
    //   username: process.env.DB_USER || 'username',
    //   password: process.env.DB_PASSWORD || 'password',
    //   database: process.env.DB_NAME || 'dbname',
    //   extra: {
    //     connectionLimit: 10,
    //   },
    //   autoLoadEntities: true,
  },
  jwt: {
    publicKey: process.env.JWT_PUBLIC_KEY || '',
  },
  rabbitmq: process.env.RABBITMQ_URL as Options.Connect,
  redis: {
    url: process.env.REDIS_URL || 'redis://localhost:6379',
  } as RedisClientOptions,
  throttlers: [
    {
      // default, 5 requests per 1 milisecond. For testing only
      name: 'default',
      ttl: 1,
      limit: 5,
    },
  ] as ThrottlerOptions[],
};
