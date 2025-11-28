// Load environment variables BEFORE any imports
require('dotenv').config({ path: 'apps/auth/.env' });

import { ConfigService } from '@nestjs/config';
import { createBetterAuthInstance } from '../src/lib/auth';
import { configuration } from '@auth/config';
import { PrismaService } from '@auth/database';
import { ClientProxy, ClientRMQ } from '@nestjs/microservices';
import { RmqUrl } from '@nestjs/microservices/external/rmq-url.interface';

export async function setupAdmin() {
  const config = new ConfigService(configuration());
  const prisma = new PrismaService(config);
  const rmq = new ClientRMQ({
    urls: [config.get('RABBITMQ_URL') as RmqUrl],
    wildcards: true,
    exchange: 'events',
    exchangeType: 'topic',
  });
  const auth = createBetterAuthInstance(config, prisma, rmq);
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
