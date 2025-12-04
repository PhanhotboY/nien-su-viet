import { StringValue } from 'ms';
export const config = {
  db: {
    url: process.env.DATABASE_URL || '',
    directUrl: process.env.DIRECT_DATABASE_URL || '',
    //   type: process.env.DB_TYPE || 'mysql',
    //   synchronize: false,
    //   logging: false,
    //   replication: {
    //     master: {
    //       host: process.env.DB_HOST || 'masterHost',
    //       port: process.env.DB_PORT || 3306,
    //       username: process.env.DB_USER || 'username',
    //       password: process.env.DB_PASSWORD || 'password',
    //       database: process.env.DB_NAME || 'dbname',
    //     },
    //     slaves: [
    //       {
    //         // fix if necessary
    //         host: 'slaveHost',
    //         port: 3306,
    //         username: 'username',
    //         password: process.env.DB_PASSWORD || 'password',
    //         database: 'dbname',
    //       },
    //     ],
    //   },
    //   extra: {
    //     connectionLimit: 30,
    //   },
    //   autoLoadEntities: true,
  },
  port: process.env.NODE_PORT || 3000,
  env: process.env.NODE_ENV,
  clientUrl: process.env.CLIENT_URL || 'http://localhost:5173',
  serverUrl: process.env.SERVER_URL || 'http://localhost:3001',
  betterAuth: {
    secret: process.env.BETTER_AUTH_SECRET || '',
    cookieDomain: process.env.COOKIE_DOMAIN || 'localhost',
    cookiePrefix: process.env.AUTH_COOKIE_PREFIX || 'nsv-auth',
  },
  jwt: {
    privateKey: process.env.JWT_PRIVATE_KEY || '',
    publicKey: process.env.JWT_PUBLIC_KEY || '',
    signOptions: {
      expiresIn: (process.env.JWT_EXPIRES_IN || '1h') as StringValue,
    },
  },
  rabbitmq: process.env.RABBITMQ_URL,
  redis: {
    url: process.env.REDIS_URL || 'redis://localhost:6379',
  },
  throttlers: [
    {
      name: 'default',
      ttl: 60000,
      limit: 100,
    },
  ],
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
};
