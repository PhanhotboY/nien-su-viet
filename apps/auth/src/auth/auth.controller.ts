import {
  All,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Req,
  Res,
} from '@nestjs/common';
import type { Request, Response } from 'express';
import { toNodeHandler } from 'better-auth/node';

import { AuthService } from './auth.service';
import { PaginatedResponseDto, Serialize } from '@phanhotboy/nsv-common';
import {
  MemberBaseDto,
  MemberBriefResponseDto,
  UserBriefResponseDto,
} from './dto';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Get('users/:id')
  @Serialize(UserBriefResponseDto)
  async getUserInfo(
    @Param('id') userId: string,
  ): Promise<UserBriefResponseDto> {
    const res = await this.authService.userInfo(userId);

    return res;
  }

  @Get('admin/organizations/:id/members')
  async getAllOrganizationMembers(
    @Param('id') orgId: string,
  ): Promise<PaginatedResponseDto<MemberBriefResponseDto>> {
    return await this.authService.getAllOrganizationMembers(orgId);
  }

  @Get('admin/organizations/:id')
  async getOrganizationById(@Param('id') orgId: string) {
    return await this.authService.getOrganizationById(orgId);
  }

  @Get('admin/organizations')
  async getAllOrganizations() {
    return await this.authService.getAllOrganizations();
  }

  @Post('admin/organizations/:orgId/members/:userId')
  async addUserToOrganization(
    @Param('orgId') organizationId: string,
    @Param('userId') userId: string,
  ) {
    return await this.authService.addUserToOrganization(userId, organizationId);
  }

  @Delete('admin/organizations/:orgId/members/:memberId')
  async removeUserFromOrganization(
    @Param('orgId') organizationId: string,
    @Param('memberId') memberId: string,
  ) {
    return await this.authService.removeUserFromOrganization(
      memberId,
      organizationId,
    );
  }

  @All('*splat')
  async handleAuthRequests(@Req() req: Request, @Res() res: Response) {
    const handler = toNodeHandler(this.authService.instance);
    await handler(req, res);
    return;
  }
}
