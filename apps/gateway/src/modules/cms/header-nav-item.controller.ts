import type { Request } from 'express';
import { Throttle } from '@nestjs/throttler';
import {
  Controller,
  Delete,
  Get,
  Inject,
  Post,
  Put,
  Req,
} from '@nestjs/common';

import { RATE_LIMIT } from '@gateway/config';
import { Permissions, Public } from '@gateway/common/decorators';
import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { RedisService, type RedisServiceType } from '@phanhotboy/nsv-common';

// ref: apps/cms/internal/headerNavItem/controller/http/headerNavItem.router.go
@Controller('header-nav-items')
export class HeaderNavItemController {
  private readonly routePath = '/api/v1/header-nav-items*';
  constructor(
    private readonly cmsProxy: HttpProxyService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
  ) {}

  @Get()
  @Public()
  proxyRequest(@Req() req: Request) {
    return this.cmsProxy.proxyRequest(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['create'] })
  async ProxyPostRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['update'] })
  async proxyPutRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['delete'] })
  async proxyDeleteRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
