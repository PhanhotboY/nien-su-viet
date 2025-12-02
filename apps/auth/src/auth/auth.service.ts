import { ClientProxy } from '@nestjs/microservices';
import { Inject, Injectable } from '@nestjs/common';

import { ConfigService, RMQ } from '@phanhotboy/nsv-common';
import { createBetterAuthInstance } from '@auth/lib/auth';
import { PrismaService } from '@auth/database';
import { Config } from '@auth/config';

@Injectable()
export class AuthService {
  private readonly auth: ReturnType<typeof createBetterAuthInstance>;

  constructor(
    private readonly prisma: PrismaService,
    private readonly config: ConfigService<Config>,
    @Inject(RMQ.TOPIC_EVENTS_EXCHANGE) private readonly rmq: ClientProxy,
  ) {
    this.auth = createBetterAuthInstance(config, prisma, rmq);
  }

  get api() {
    return this.auth.api;
  }
  get instance() {
    return this.auth;
  }
}
