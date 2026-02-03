import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { Request } from 'express';
import * as jwt from 'jsonwebtoken';
import { Cookie } from 'cookiejar';

import { ConfigService } from '@phanhotboy/nsv-common';
import { IS_PUBLIC_KEY } from '../decorators/public.decorator';
import { UserBaseDto } from '@gateway/modules/auth/dto';
import { Config } from '@gateway/config';

interface JWTPayload extends jwt.JwtPayload {
  user: UserBaseDto;
  session: any;
  version: string;
  updatedAt: number;
}

declare global {
  namespace Express {
    interface Request {
      user?: UserBaseDto;
      session?: any;
    }
  }
}

@Injectable()
export class BetterAuthGuard implements CanActivate {
  constructor(
    private reflector: Reflector,
    private configService: ConfigService<Config>,
  ) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    // Check if route is marked as public
    const isPublic = this.reflector.getAllAndOverride<boolean>(IS_PUBLIC_KEY, [
      context.getClass(),
      context.getHandler(),
    ]);

    if (isPublic) {
      return true;
    }

    const request = context.switchToHttp().getRequest<Request>();

    try {
      // Extract JWT token from Authorization header or cookie
      const cookies = this.parseCookies(request.headers.cookie || '');
      const token = cookies.get(
        `${this.configService.get('betterAuth.cookiePrefix')}.session_data`,
      );

      if (!token) {
        throw new UnauthorizedException('No authentication token found');
      }

      const secret = this.configService.get('betterAuth.secret');
      if (!secret) {
        throw new Error('JWT secret not configured');
      }

      const decoded = jwt.verify(token, secret, {
        algorithms: ['HS256'],
      }) as JWTPayload;

      // Attach user info to request
      request['user'] = decoded.user;
      request['session'] = decoded.session;

      return true;
    } catch (error) {
      if (error instanceof jwt.JsonWebTokenError) {
        throw new UnauthorizedException('Invalid token');
      }
      if (error instanceof jwt.TokenExpiredError) {
        throw new UnauthorizedException('Token expired');
      }
      throw new UnauthorizedException('Authentication failed');
    }
  }

  private parseCookies(cookieHeader: string): Map<string, string> {
    const cookies = new Map<string, string>();
    const pairs = cookieHeader.split(';');
    for (const pair of pairs) {
      const cookie = new Cookie(pair.trim());
      cookies.set(cookie.name, cookie.value);
    }
    return cookies;
  }
}
