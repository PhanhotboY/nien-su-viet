import type { Request } from 'express';
import { Body, Controller, Get, Inject, Put, Req } from '@nestjs/common';

import { HttpProxyService } from '@gateway/common/services/http-proxy.service';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';
import { Permissions, Public } from '@gateway/common/decorators';
import {
  RedisService,
  Serialize,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import { AppDto, AppUpdateDto } from './dto';
import { OperationResponseDto } from '@phanhotboy/nsv-common/dto/response/operation-response.dto';

@Controller('app')
export class AppController {
  private readonly routePath = '/api/v1/app*';
  constructor(
    private readonly cmsProxy: HttpProxyService,
    @Inject(RedisService) private readonly redis: RedisServiceType,
  ) {}

  @Get()
  @Public()
  @Serialize(AppDto)
  async proxyRequest(@Req() req: Request): Promise<AppDto> {
    const res = await this.cmsProxy.proxyRequest<AppDto>(req);
    return res as AppDto;
  }

  @Put()
  @Throttle(RATE_LIMIT.INTERNAL)
  @Permissions({ app: ['update'] })
  @Serialize(OperationResponseDto)
  async proxyPutRequest(
    @Req() req: Request,
    @Body() body: AppUpdateDto,
  ): Promise<OperationResponseDto> {
    await this.redis.mdel(this.routePath);
    return await this.cmsProxy.proxyRequest(req);
  }
}
