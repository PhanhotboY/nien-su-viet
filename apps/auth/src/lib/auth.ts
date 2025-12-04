import { betterAuth } from 'better-auth';
import { prismaAdapter } from 'better-auth/adapters/prisma';
import { admin as adminPlugin, openAPI } from 'better-auth/plugins';
import { ClientProxy } from '@nestjs/microservices';

import { type Config } from '../config';
import {
  ConfigService,
  USER_EVENT,
  UserBaseDto,
  UserDeleteDto,
} from '@phanhotboy/nsv-common';
import { ac, roles } from '@phanhotboy/nsv-common/auth';
import { PrismaService } from '@auth/database';
import { plainToInstance } from 'class-transformer';

export function createBetterAuthInstance(
  config: ConfigService<Config>,
  prisma: PrismaService,
  rmq: ClientProxy,
) {
  return betterAuth({
    database: prismaAdapter(prisma, {
      provider: 'postgresql',
    }),
    secret: config.get('betterAuth.secret'),
    basePath: '/api/v1/auth',
    trustedOrigins: config.get('trustedOrigins'),
    emailAndPassword: { enabled: true },
    account: {
      accountLinking: { enabled: true, trustedProviders: ['google'] },
    },
    session: {
      expiresIn: 60 * 60 * 24, // 1 day
      updateAge: 60 * 60 * 1, // 1 hour (every 1 hour the session expiration is updated)
      cookieCache: {
        enabled: true,
        maxAge: 60 * 15, // 15 minutes
        strategy: 'jwt',
      },
    },
    advanced: {
      cookiePrefix: 'nsv-auth',
      crossSubDomainCookies: {
        enabled: true,
        domain: config.get('betterAuth.cookieDomain'),
      },
    },
    plugins: [
      adminPlugin({
        ac,
        roles,
        adminRoles: ['admin'],
        defaultRole: 'user',
      }),
      openAPI(),
    ],
    databaseHooks: {
      user: {
        create: {
          after: async (user) => {
            const userDto = plainToInstance(UserBaseDto, user);
            rmq.emit(USER_EVENT.REGISTERED, userDto);
          },
        },
        update: {
          after: async (user) => {
            const userDto = plainToInstance(UserBaseDto, user);
            rmq.emit(USER_EVENT.UPDATED, userDto);
          },
        },
        delete: {
          after: async (user) => {
            const userDto = plainToInstance(UserDeleteDto, {
              userId: user.id,
            });
            rmq.emit(USER_EVENT.DELETED, userDto);
          },
        },
      },
    },
    socialProviders: {
      google: {
        clientId: config.get('google.clientId'),
        clientSecret: config.get('google.clientSecret'),
      },
    },
  });
}
