import { All, Controller, Req, Res } from '@nestjs/common';
import type { Request, Response } from 'express';
import { toNodeHandler } from 'better-auth/node';

import { AuthService } from './auth.service';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @All('*splat')
  async andleAuthRequests(@Req() req: Request, @Res() res: Response) {
    console.log('handle auth request');
    const handler = toNodeHandler(this.authService.instance);
    await handler(req, res);
    return;
  }
}
