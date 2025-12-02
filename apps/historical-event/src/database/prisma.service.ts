import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import { PrismaClient } from '@historical-event-prisma';
import { PrismaPg } from '@prisma/adapter-pg';
import { ConfigService } from '@phanhotboy/nsv-common';
import { Config } from '@historical-event/config';

@Injectable()
export class PrismaService
  extends PrismaClient
  implements OnModuleInit, OnModuleDestroy
{
  constructor(private readonly configService: ConfigService<Config>) {
    const adapter = new PrismaPg({
      connectionString: configService.get('db.url'),
    });
    super({ adapter, log: ['query', 'info', 'warn', 'error'] });
  }

  async onModuleInit() {
    await this.$connect();
    console.log('Prisma connected to the database');
  }

  async onModuleDestroy() {
    await this.$disconnect();
  }
}
