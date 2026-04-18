import { Throttle } from '@nestjs/throttler';
import type { Request, Response } from 'express';
import {
  All,
  Controller,
  Get,
  Inject,
  Logger,
  Param,
  Post,
  Req,
  Res,
} from '@nestjs/common';

import { Serialize } from '@phanhotboy/nsv-common';
import { RATE_LIMIT } from '@gateway/config';
import { Public } from '@gateway/common/decorators';
import { UserBriefResponseDto } from './dto';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';

@Controller('auth')
@Public() // Let Better Auth handle auth internally
export class AuthController {
  constructor(
    private readonly proxyService: HttpProxyService,
    private readonly logger: Logger,
  ) {}

  @Get('users/:id')
  @Serialize(UserBriefResponseDto)
  @Throttle(RATE_LIMIT.INTERNAL)
  async getUserInfo(@Param('id') userId: string, @Req() req: Request) {
    this.logger.debug(`Fetching info for user: ${userId}`);
    return await this.proxyService.makeRequest(req);
  }

  // Forward specific auth routes for better rate limiting, and documentation
  @Post('sign-in/email')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardSignInEmail(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('sign-in/social')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardSignInSocial(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('sign-out')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardSignOut(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('sign-up/email')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardSignUpEmail(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('reset-password')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardResetPassword(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('send-verification-email')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardSendVerificationEmail(
    @Req() req: Request,
    @Res() res: Response,
  ) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('change-email')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardChangeEmail(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('change-password')
  @Throttle(RATE_LIMIT.AUTH)
  async forwardChangePassword(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('update-user')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardUpdateUser(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('delete-user')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardDeleteUser(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('revoke-session')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardRevokeSession(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('revoke-sessions')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardRevokeSessions(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('link-social')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardLinkSocial(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('unlink-social')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardUnlinkSocial(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('refresh-token')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardRefreshToken(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @Post('get-refresh-token')
  @Throttle(RATE_LIMIT.INTERNAL)
  async forwardGetRefreshToken(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }

  @All('*splat')
  @Public()
  async forwardAuthHandler(@Req() req: Request, @Res() res: Response) {
    await this.proxyService.proxyRequest(req, res);
  }
}
