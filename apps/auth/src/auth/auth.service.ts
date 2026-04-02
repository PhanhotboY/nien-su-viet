import { ClientProxy } from '@nestjs/microservices';
import { Inject, Injectable, NotFoundException } from '@nestjs/common';

import { ConfigService } from '@phanhotboy/nsv-common';
import { createBetterAuthInstance } from '@auth/lib/auth';
import { PrismaService } from '@auth/database';
import { Config } from '@auth/config';
import { RMQ } from '@phanhotboy/constants';

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

  async userInfo(userId: string) {
    const user = await this.prisma.user.findUnique({
      where: { id: userId },
      select: {
        id: true,
        name: true,
        image: true,
      },
    });
    if (!user) {
      throw new NotFoundException('User not found');
    }
    return user;
  }
}
