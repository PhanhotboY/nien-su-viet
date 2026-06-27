import path from 'node:path';
require('dotenv').config({ path: path.resolve(process.cwd(), '.env') });

import { createBetterAuthInstance } from '../src/lib/auth';
import { PrismaService } from '@auth/database';
import {
  RedisService,
  RedisServiceType,
  ConfigService,
  RmqService,
} from '@phanhotboy/nsv-common';
import { NestFactory } from '@nestjs/core';
import { AppModule } from '@auth/app.module';
import { Config } from '@auth/config';
import { RMQ } from '@phanhotboy/constants';
import { ClientProxy } from '@nestjs/microservices';

export async function setupAdmin() {
  const app = await NestFactory.create(AppModule, { bodyParser: false });
  const config = app.get(ConfigService<Config>);
  const prisma = app.get(PrismaService);
  const rmq = app.get(RMQ.TOPIC_EVENTS_EXCHANGE) as ClientProxy;
  const redis = app.get(RedisService);

  const auth = createBetterAuthInstance(
    config,
    prisma,
    rmq,
    redis as RedisServiceType,
  );
  // Dynamic import after env is loaded
  const adminEmail = process.env.ADMIN_EMAIL;
  const adminName = process.env.ADMIN_NAME || 'Admin User';
  const adminPassword = process.env.ADMIN_PASSWORD;
  if (!adminEmail || !adminPassword) {
    return console.error(
      'ADMIN_EMAIL and ADMIN_PASSWORD must be set in environment variables.',
    );
  }

  const admin = await auth.api.createUser({
    body: {
      email: adminEmail,
      name: adminName,
      role: 'admin',
      password: adminPassword,
    },
  });

  console.log('Admin user created:', admin.user);
}
