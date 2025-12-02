import { RedisClientOptions } from '@keyv/redis';
import { JwtModuleOptions } from '@nestjs/jwt';
import { ThrottlerOptions } from '@nestjs/throttler';
import { Options } from 'amqplib';
import { StringValue } from 'ms';

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
  port: process.env.NODE_PORT || 3000,
  env: process.env.NODE_ENV || 'development',
  clientUrl: process.env.CLIENT_URL || 'http://localhost:5173',
  serverUrl: process.env.SERVER_URL || 'http://localhost:3001',
  betterAuth: {
    secret: process.env.BETTER_AUTH_SECRET || '',
  },
  jwt: {
    privateKey: process.env.JWT_PRIVATE_KEY || '',
    publicKey: process.env.JWT_PUBLIC_KEY || '',
    signOptions: {
      expiresIn: (process.env.JWT_EXPIRES_IN || '1h') as StringValue,
    },
  } as JwtModuleOptions,
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
  mail: {
    host: process.env.SMTP_HOST || 'smtp.example.com',
    port: Number(process.env.SMTP_PORT) || 587,
    secure: false,
    auth: {
      user: process.env.SMTP_USER || 'user@example.com',
      pass: process.env.SMTP_PASSWORD || 'password',
    },
  },
  google: {
    clientId: process.env.GOOGLE_CLIENT_ID || '',
    clientSecret: process.env.GOOGLE_CLIENT_SECRET || '',
  },
  trustedOrigins: process.env.ALLOWED_ORIGINS?.split(',') || [],
} as const;
