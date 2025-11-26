import { Injectable } from '@nestjs/common';
import { auth } from '@auth/lib/auth';
import { PrismaService } from '@auth/database';

@Injectable()
export class AuthService {
  private readonly auth = auth;

  constructor(private readonly prisma: PrismaService) {}

  get api() {
    return this.auth.api;
  }
  get instance() {
    return this.auth;
  }

  async deleteUser(userId: string) {
    await this.prisma.user.delete({ where: { id: userId } });
  }
}
