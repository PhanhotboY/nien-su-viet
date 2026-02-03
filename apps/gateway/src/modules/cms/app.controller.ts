import type { Request } from 'express';
import { Controller, Get, Put, Req } from '@nestjs/common';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';
import { Permissions, Public } from '@gateway/common/decorators';

@Controller('app')
export class AppController {
  constructor(private readonly cmsProxy: HttpProxyService) {}

  @Get()
  @Public()
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Put()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ app: ['update'] })
  proxyPutRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }
}
