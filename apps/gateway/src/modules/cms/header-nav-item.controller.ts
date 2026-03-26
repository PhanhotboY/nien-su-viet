import type { Request } from 'express';
import { Throttle } from '@nestjs/throttler';
import {
  Body,
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
import {
  RedisService,
  Serialize,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import {
  HeaderNavItemCreateDto,
  HeaderNavItemDto,
  HeaderNavItemUpdateDto,
} from './dto';
import { OperationMetadataDto } from '@phanhotboy/nsv-common/dto/response';

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
  @Serialize(HeaderNavItemDto)
  async proxyRequest(@Req() req: Request): Promise<HeaderNavItemDto> {
    return await this.cmsProxy.makeRequest<HeaderNavItemDto>(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['create'] })
  @Serialize(OperationMetadataDto)
  async proxyPostRequest(
    @Req() req: Request,
    @Body() body: HeaderNavItemCreateDto,
  ) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.makeRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['update'] })
  @Serialize(OperationMetadataDto)
  async proxyPutRequest(
    @Req() req: Request,
    @Body() body: HeaderNavItemUpdateDto,
  ) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.makeRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['delete'] })
  @Serialize(OperationMetadataDto)
  async proxyDeleteRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.makeRequest(req);
  }
}
