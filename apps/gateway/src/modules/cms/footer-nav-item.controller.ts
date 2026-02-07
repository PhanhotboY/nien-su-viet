import {
  Controller,
  Delete,
  Get,
  Inject,
  Post,
  Put,
  Req,
} from '@nestjs/common';
import type { Request } from 'express';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Permissions, Public } from '@gateway/common/decorators';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';
import { RedisService, type RedisServiceType } from '@phanhotboy/nsv-common';

// ref: apps/cms/internal/footerNavItem/controller/http/footerNavItem.router.go
@Controller('footer-nav-items')
export class FooterNavItemController {
  private readonly routePath = '/api/v1/footer-nav-items*';
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
  @Permissions({ footerNavItem: ['create'] })
  async ProxyPostRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['update'] })
  async proxyPutRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['delete'] })
  async proxyDeleteRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
