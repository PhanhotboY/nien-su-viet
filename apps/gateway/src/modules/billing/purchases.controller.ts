import { Body, Controller, Inject, Post } from '@nestjs/common';

import {
  ConfigService,
  RedisService,
  type RedisServiceType,
} from '@phanhotboy/nsv-common';
import { PurchasesService } from './purchases.service';
import { CurrentUser } from '@gateway/common/decorators';
import { Config } from '@gateway/config';

@Controller('billing/purchases')
export class PurchasesController {
  private readonly routePath = '/api/v1/billing/purchases*';

  constructor(
    @Inject(RedisService) private readonly redis: RedisServiceType,
    private readonly config: ConfigService<Config>,
    private readonly purchasesService: PurchasesService,
  ) {}

  @Post()
  async createPurchase(
    @CurrentUser('id') userId: string,
    @Body() payload: any,
  ) {
    return this.purchasesService.createPurchase(userId, payload);
  }
}
