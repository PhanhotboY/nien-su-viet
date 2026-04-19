import { All, Controller, Get, Param, Req, Res } from '@nestjs/common';
import type { Request, Response } from 'express';
import { toNodeHandler } from 'better-auth/node';

import { AuthService } from './auth.service';
import { Serialize } from '@phanhotboy/nsv-common';
import { UserBriefResponseDto } from './dto';

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

  @All('*splat')
  async handleAuthRequests(@Req() req: Request, @Res() res: Response) {
    const handler = toNodeHandler(this.authService.instance);
    await handler(req, res);
    return;
  }
}
