import { APIError, betterAuth } from 'better-auth';
import { prismaAdapter } from 'better-auth/adapters/prisma';
import {
  admin as adminPlugin,
  createAuthMiddleware,
  openAPI,
} from 'better-auth/plugins';
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
      crossSubDomainCookies: {
        enabled: true,
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
    hooks: {
      after: createAuthMiddleware(async (ctx) => {
        const isAPIError = ctx.context.returned instanceof APIError;
        if (isAPIError) return;

        const returned = ctx.context.returned || ({} as any);
        const user = plainToInstance(UserBaseDto, returned.user);

        switch (ctx.path) {
          case '/sign-up/email':
          case '/admin/create-user':
            if (!user) return;
            rmq.emit(USER_EVENT.REGISTERED, user);
            break;

          case '/admin/update-user':
            if (!user) return;
            rmq.emit(USER_EVENT.UPDATED, user);
            break;

          case '/admin/remove-user':
            const isSuccess = (ctx.context.returned as any)?.success === true;
            const userId = ctx.body?.userId;
            if (!isSuccess || !userId) return;
            rmq.emit(USER_EVENT.DELETED, {
              userId: ctx.body?.userId,
            } as UserDeleteDto);
            break;
        }
      }),
    },
    socialProviders: {
      google: {
        clientId: config.get('google.clientId'),
        clientSecret: config.get('google.clientSecret'),
      },
    },
  });
}
