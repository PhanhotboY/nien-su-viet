import { betterAuth } from 'better-auth';
import { prismaAdapter } from 'better-auth/adapters/prisma';
import { admin, createAuthMiddleware, openAPI } from 'better-auth/plugins';
import { PrismaPg } from '@prisma/adapter-pg';
import { Pool } from 'pg';

import { PrismaClient } from '../../generated/prisma';
import { configuration, type Config } from '../config';

const config = configuration() as Config;

const pool = new Pool({
  connectionString: config.db.url,
});
const adapter = new PrismaPg(pool);
const prisma = new PrismaClient({
  adapter,
  log: ['query', 'info', 'warn', 'error'],
});

export const auth = betterAuth({
  database: prismaAdapter(prisma, {
    provider: 'postgresql',
  }),
  secret: config.betterAuth.secret,
  basePath: '/api/v1/auth',
  trustedOrigins: config.trustedOrigins,
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
    admin({
      adminRoles: ['admin'],
      defaultRole: 'user',
    }),
    openAPI(),
  ],
  hooks: {
    after: createAuthMiddleware(async (ctx) => {
      if (ctx.path.startsWith('/sign-up')) {
        console.log('After hook context:', ctx.error);
        return ctx;
      }
    }),
  },
  socialProviders: {
    google: {
      clientId: config.google.clientId,
      clientSecret: config.google.clientSecret,
    },
  },
});
