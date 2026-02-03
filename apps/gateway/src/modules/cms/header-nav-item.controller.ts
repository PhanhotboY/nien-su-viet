import type { Request } from 'express';
import { Throttle } from '@nestjs/throttler';
import { Controller, Delete, Get, Post, Put, Req } from '@nestjs/common';

import { RATE_LIMIT } from '@gateway/config';
import { Permissions, Public } from '@gateway/common/decorators';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';

// ref: apps/cms/internal/headerNavItem/controller/http/headerNavItem.router.go
@Controller('header-nav-items')
export class HeaderNavItemController {
  constructor(private readonly cmsProxy: HttpProxyService) {}

  @Get()
  @Public()
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['create'] })
  ProxyPostRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['update'] })
  proxyPutRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['delete'] })
  proxyDeleteRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }
}
