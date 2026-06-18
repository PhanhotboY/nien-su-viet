import { Inject, Injectable, NotFoundException } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';

import { PrismaService } from '@auth/database';
import { RMQ } from '@phanhotboy/constants';
import { RedisService, type RedisServiceType } from '@phanhotboy/nsv-common';
import { AuthService } from '@auth/auth';

@Injectable()
export class UserService {
  constructor(
    private readonly prisma: PrismaService,
    @Inject(RMQ.TOPIC_EVENTS_EXCHANGE) private readonly rmq: ClientProxy,
    @Inject(RedisService) private readonly redisService: RedisServiceType,
    private readonly authService: AuthService,
  ) {}

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
