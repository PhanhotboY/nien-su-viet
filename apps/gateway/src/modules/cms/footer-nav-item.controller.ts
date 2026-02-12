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
import type { Request } from 'express';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Permissions, Public } from '@gateway/common/decorators';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';
import {
  RedisService,
  Serialize,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import {
  FooterNavItemCreateDto,
  FooterNavItemDto,
  FooterNavItemUpdateDto,
} from './dto';
import { OperationResponseDto } from '@phanhotboy/nsv-common/dto/response/operation-response.dto';

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
  // @Serialize(FooterNavItemDto)
  proxyRequest(@Req() req: Request): Promise<FooterNavItemDto> {
    return this.cmsProxy.proxyRequest<FooterNavItemDto>(req);
  }

  @Post()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['create'] })
  @Serialize(OperationResponseDto)
  async proxyPostRequest(
    @Req() req: Request,
    @Body() body: FooterNavItemCreateDto,
  ): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['update'] })
  @Serialize(OperationResponseDto)
  async proxyPutRequest(
    @Req() req: Request,
    @Body() body: FooterNavItemUpdateDto,
  ): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['delete'] })
  @Serialize(OperationResponseDto)
  async proxyDeleteRequest(@Req() req: Request): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
