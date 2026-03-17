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
import { OperationMetadataDto } from '@phanhotboy/nsv-common/dto/response';

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
  @Serialize(OperationMetadataDto)
  async proxyPostRequest(
    @Req() req: Request,
    @Body() body: FooterNavItemCreateDto,
  ) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Put(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['update'] })
  @Serialize(OperationMetadataDto)
  async proxyPutRequest(
    @Req() req: Request,
    @Body() body: FooterNavItemUpdateDto,
  ) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }

  @Delete(':id')
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ footerNavItem: ['delete'] })
  @Serialize(OperationMetadataDto)
  async proxyDeleteRequest(@Req() req: Request) {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
