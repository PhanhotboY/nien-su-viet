import { Controller, Delete, Get, Post, Put, Req, Res } from '@nestjs/common';
import type { Request } from 'express';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Permissions, Public } from '@gateway/common/decorators';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';

// ref: apps/cms/internal/footerNavItem/controller/http/footerNavItem.router.go
@Controller('footer-nav-items')
export class FooterNavItemController {
  constructor(private readonly cmsProxy: HttpProxyService) {}

  @Get()
  @Public()
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['create'] })
  ProxyPostRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['update'] })
  proxyPutRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['delete'] })
  proxyDeleteRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }
}
