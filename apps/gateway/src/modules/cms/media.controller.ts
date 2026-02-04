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

// ref: apps/cms/internal/media/controller/http/media.router.go
@Controller('media')
export class MediaController {
  private readonly routePath = '/api/v1/media*';
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
  @Permissions({ media: ['create'] })
  async ProxyPostRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ media: ['delete'] })
  async proxyPutRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ media: ['delete'] })
  async proxyDeleteRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
