import { Body, Controller, Get, Inject, Post, Query } from '@nestjs/common';

import {
  ConfigService,
  RedisService,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import { PlansService } from './plans.service';
import { CurrentUser, Public } from '@gateway/common/decorators';
import { Config } from '@gateway/config';

@Controller('billing/plans')
export class PlansController {
  private readonly routePath = '/api/v1/billing/plans*';

  constructor(
    @Inject(RedisService) private readonly redis: RedisServiceType,
    private readonly config: ConfigService<Config>,
    private readonly plansService: PlansService,
  ) {}

  @Get()
  @Public()
  async listPlans(@Query() query: any) {
    return this.plansService.listPlans(query);
  }

  @Post()
  async createPlan(@CurrentUser('id') userId: string, @Body() payload: any) {
    return this.plansService.createPlan(userId, payload);
  }
}
