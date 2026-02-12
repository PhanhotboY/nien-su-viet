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
import { OperationResponseDto } from '@phanhotboy/nsv-common/dto/response/operation-response.dto';

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
    return await this.cmsProxy.proxyRequest<HeaderNavItemDto>(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['create'] })
  @Serialize(OperationResponseDto)
  async proxyPostRequest(
    @Req() req: Request,
    @Body() body: HeaderNavItemCreateDto,
  ): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['update'] })
  @Serialize(OperationResponseDto)
  async proxyPutRequest(
    @Req() req: Request,
    @Body() body: HeaderNavItemUpdateDto,
  ): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ headerNavItem: ['delete'] })
  @Serialize(OperationResponseDto)
  async proxyDeleteRequest(@Req() req: Request): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
