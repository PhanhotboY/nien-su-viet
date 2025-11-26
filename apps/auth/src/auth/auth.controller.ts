import { All, Controller, Delete, Param, Req, Res } from '@nestjs/common';
import type { Request, Response } from 'express';
import { toNodeHandler } from 'better-auth/node';

import { auth } from '@auth/lib/auth';
import { AuthService } from './auth.service';
import { Public, Roles } from '@phanhotboy/nsv-common';

@Controller()
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Roles('admin')
  @Delete('auth/admin/delete-user/:userId')
  async deleteUser(@Param('userId') userId: string) {
    await this.authService.deleteUser(userId);
  }

  @Public() // Better-Auth is secured by default
  @All('auth/*splat')
  handleAuthRequests(@Req() req: Request, @Res() res: Response) {
    const handler = toNodeHandler(auth);
    handler(req, res);
  }
}
