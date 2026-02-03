import { PrismaService } from '@historical-event/database';
import { Inject, Injectable } from '@nestjs/common';
import { ClientProxy } from '@nestjs/microservices';
import { isUUID } from 'class-validator';

import {
  RedisService,
  type RedisServiceType,
  UtilService,
} from '@phanhotboy/nsv-common';
import { UserBaseDto } from './dto';

@Injectable()
export class UserService {
  private readonly cachePrefix = 'user';
  private readonly cacheKey: string;

  constructor(
    private readonly prisma: PrismaService,
    private readonly util: UtilService,
    @Inject(RedisService) private readonly redisService: RedisServiceType,
  ) {
    this.cacheKey = this.util.genCacheKey(this.cachePrefix);
  }

  async createUser(user: UserBaseDto) {
    const newUser = await this.prisma.user.create({
      data: {
        id: user.id,
        email: user.email,
        name: user.name,
        image: user.image,
        role: user.role,
      },
    });

    await this.redisService.del(this.cacheKey);

    return newUser;
  }

  async handleUserRegister(data: UserBaseDto) {
    const user = await this.findUserById(data.id).catch(() => null);

    if (user) {
      return user;
    }
    return await this.createUser(data);
  }

  async findUserById(id: string) {
    const options = { where: { id } } satisfies Parameters<
      typeof this.prisma.user.findUnique
    >[0];

    return this.util.handleHashCachingQuery(
      {
        cacheKey: this.cacheKey,
        hashAttribute: options,
        notFoundMessage: 'Người dùng không tồn tại',
      },
      () => {
        return this.prisma.user.findUnique(options);
      },
    );
  }

  async updateUser(data: UserBaseDto) {
    await this.findUserById(data.id); // Ensure user exists

    const updatedUser = await this.prisma.user.update({
      where: { id: data.id },
      data: {
        email: data.email,
        name: data.name,
        image: data.image,
        role: data.role,
      },
    });

    await this.redisService.del(this.cacheKey);

    return updatedUser;
  }

  async deleteUser(id: string) {
    await this.findUserById(id); // Ensure user exists

    await this.prisma.user.delete({ where: { id } });

    await this.redisService.del(this.cacheKey);

    return { success: true };
  }
}
