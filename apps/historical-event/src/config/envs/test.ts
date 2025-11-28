// export * from './development';
export const config = {
  db: {
    url: process.env.DATABASE_URL || '',
    directUrl: process.env.DIRECT_DATABASE_URL || '',
    type: 'postgres',
    synchronize: false,
    logging: false,
    host: process.env.DB_HOST || '127.0.0.1',
    port: process.env.DB_PORT || 3306,
    username: process.env.DB_USER || 'username',
    password: process.env.DB_PASSWORD || 'password',
    database: process.env.DB_NAME || 'dbname',
    extra: {
      connectionLimit: 5,
    },
    autoLoadEntities: true,
  },
  jwt: {
    publicKey: process.env.JWT_PUBLIC_KEY || '',
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
};
