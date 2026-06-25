import { ClientProxy } from '@nestjs/microservices';
import { Inject, Injectable, NotFoundException } from '@nestjs/common';

import {
  ConfigService,
  OperationMetadataDto,
  RedisService,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
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
    @Inject(RedisService) private readonly redisService: RedisServiceType,
  ) {
    this.auth = createBetterAuthInstance(config, prisma, rmq, redisService);
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

  getAllOrganizations() {
    return this.prisma.organization.findMany({
      select: {
        id: true,
        name: true,
        members: { select: { id: true } },
        slug: true,
        createdAt: true,
        logo: true,
      },
    });
  }

  async getOrganizationById(orgId: string) {
    const org = await this.prisma.organization.findUnique({
      where: { id: orgId },
      select: {
        id: true,
        name: true,
        members: { select: { id: true } },
        slug: true,
        createdAt: true,
        logo: true,
      },
    });
    if (!org) {
      throw new NotFoundException('Organization not found');
    }
    return org;
  }

  async getOrganizationBySlug(slug: string) {
    const org = await this.prisma.organization.findFirst({
      where: { slug },
      select: {
        id: true,
        name: true,
        members: { select: { id: true } },
        slug: true,
        createdAt: true,
        logo: true,
      },
    });

    return org;
  }

  async getAllOrganizationMembers(orgId: string) {
    const org = await this.prisma.organization.findUnique({
      where: { id: orgId },
    });
    if (!org) {
      throw new NotFoundException('Organization not found');
    }

    const members = await this.prisma.member.findMany({
      where: { organizationId: orgId },
      select: {
        id: true,
        createdAt: true,
        userId: true,
        organizationId: true,
        role: true,
        user: {
          select: {
            id: true,
            name: true,
            image: true,
            email: true,
          },
        },
      },
    });

    return {
      data: members,
      pagination: {
        limit: members.length,
        total: members.length,
        page: 1,
        totalPages: 1,
      },
    };
  }

  async addUserToOrganization(userId: string, organizationId: string) {
    await this.auth.api.addMember({
      body: {
        role: ['member'],
        userId,
        organizationId,
      },
    });

    return { success: true } as OperationMetadataDto;
  }

  async removeUserFromOrganization(memberId: string, organizationId: string) {
    await this.prisma.member.delete({
      where: { id: memberId, organizationId },
    });

    return { success: true } as OperationMetadataDto;
  }
}
