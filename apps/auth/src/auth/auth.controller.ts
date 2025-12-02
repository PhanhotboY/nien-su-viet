import { All, Controller, Req, Res } from '@nestjs/common';
import type { Request, Response } from 'express';
import { toNodeHandler } from 'better-auth/node';

import { AuthService } from './auth.service';
import { Public } from '@phanhotboy/nsv-common';

@Controller()
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Public() // Better-Auth is secured by default
  @All('auth/*splat')
  async andleAuthRequests(@Req() req: Request, @Res() res: Response) {
    const handler = toNodeHandler(this.authService.instance);
    await handler(req, res);
  }
}
