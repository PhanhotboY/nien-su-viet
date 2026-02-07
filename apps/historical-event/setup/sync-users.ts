// Load environment variables BEFORE any imports
require('dotenv').config({ path: 'apps/historical-event/.env' });

import { ConfigService } from '@phanhotboy/nsv-common';
import { configuration, type Config } from '@historical-event/config';
import { PrismaService } from '@historical-event/database';

let prisma: PrismaService;

async function syncUsersFromAuthService() {
  const gatewayEndpoint = process.env.GATEWAY_ENDPOINT;
  const adminEmail = process.env.ADMIN_EMAIL;
  const adminPassword = process.env.ADMIN_PASSWORD;
  if (!gatewayEndpoint || !adminEmail || !adminPassword) {
    throw new Error('Required environment variables are not defined.');
  }

  console.log(`Using Gateway endpoint: ${gatewayEndpoint}`);
  const configService = new ConfigService<Config>(configuration());
  prisma = new PrismaService(configService);

  const signInRes = await fetch(`${gatewayEndpoint}/auth/sign-in/email`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email: adminEmail,
      password: adminPassword,
    }),
  });
  if (!signInRes.ok) {
    throw new Error(
      `Failed to sign in: ${signInRes.status} ${signInRes.statusText}`,
    );
  }
  const token = signInRes.headers
    .get('set-cookie')
    ?.match(/nsv-auth\.session_token=([^;]+);/)?.[1];
  if (!token) {
    throw new Error('No token received from sign-in response.');
  }
  console.log(`Successfully signed in. Token length: ${token.length}`);

  const listUserRes = await fetch(`${gatewayEndpoint}/auth/admin/list-users`, {
    method: 'GET',
    headers: {
      Cookie: `nsv-auth.session_token=${token}`,
    },
  });
  if (!listUserRes.ok) {
    throw new Error(
      `Failed to list users: ${listUserRes.status} ${listUserRes.statusText}`,
    );
  }
  const listUserData = await listUserRes.json();
  const users = listUserData.users;
  if (!Array.isArray(users)) {
    throw new Error('Invalid users data received from list-users response.');
  }
  console.log(`Fetched ${users.length} users from Auth service.`);

  for (const user of users) {
    const exists = await prisma.user.findFirst({
      where: { id: user.id },
    });
    if (exists) {
      console.log(`User already exists: ${user.email}`);
      continue;
    }

    await prisma.user.create({
      data: {
        id: user.id,
        email: user.email,
        name: user.name,
        role: user.role,
        image: user.image,
      },
    });
    console.log(`Created user: ${user.email}`);
  }
}

syncUsersFromAuthService()
  .catch((err) => {
    console.error('Error sync users from Auth service:', err);
    process.exitCode = 1;
  })
  .finally(async () => {
    if (prisma) {
      await prisma.$disconnect();
    }
    process.exit();
  });
