import type { Request } from 'express';
import { Controller, Get, Inject, Put, Req } from '@nestjs/common';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';
import { Permissions, Public } from '@gateway/common/decorators';
import { RedisService, type RedisServiceType } from '@phanhotboy/nsv-common';

@Controller('app')
export class AppController {
  private readonly routePath = '/api/v1/app*';
  constructor(
    private readonly cmsProxy: HttpProxyService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
  ) {}

  @Get()
  @Public()
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Put()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ app: ['update'] })
  async proxyPutRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
