import { Controller, Get, Req } from '@nestjs/common';
import type { Request } from 'express';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';

@Controller('ping')
export class PingController {
  constructor(private readonly cmsProxy: HttpProxyService) {}

  @Get('100')
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }
}
