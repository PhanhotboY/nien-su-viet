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
