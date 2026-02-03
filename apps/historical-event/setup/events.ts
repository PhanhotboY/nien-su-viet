// Load environment variables BEFORE any imports
require('dotenv').config({ path: 'apps/historical-event/.env' });

import { ConfigService } from '@phanhotboy/nsv-common';
import { configuration, type Config } from '@historical-event/config';
import { PrismaService } from '@historical-event/database';
import { mockEvents } from './mock-events';

let prisma: PrismaService;

async function seedMockHistoricalEvents() {
  const configService = new ConfigService<Config>(configuration());
  prisma = new PrismaService(configService);

  const author = await prisma.user.findFirst();
  if (!author) {
    throw new Error(
      'No users found in the database. Please create an admin user first.',
    );
  }

  for (const event of mockEvents) {
    const exists = await prisma.historicalEvent.findFirst({
      where: { name: event.name },
    });

    if (exists) {
      console.log(`Skipping existing event: ${event.name}`);
      continue;
    }

    await prisma.historicalEvent.create({
      data: {
        ...event,
        authorId: author.id,
      },
    });

    console.log(`Created event: ${event.name}`);
  }
}

seedMockHistoricalEvents()
  .catch((err) => {
    console.error('Error seeding mock historical events:', err);
    process.exitCode = 1;
  })
  .finally(async () => {
    if (prisma) {
      await prisma.$disconnect();
    }
    process.exit();
  });
