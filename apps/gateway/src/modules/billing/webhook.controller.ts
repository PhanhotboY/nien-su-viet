import { Body, Controller, Post, Res } from '@nestjs/common';
import { type Response } from 'express';

import { Public } from '@gateway/common/decorators';
import { WebhookService } from './webhook.service';
import { WebhookRequestDto } from './dto/webhook.req.dto';
import { Throttle } from '@nestjs/throttler';
import { RATE_LIMIT } from '@gateway/config';

@Controller('billing/zalopay')
export class WebhookController {
  private readonly routePath = '/api/v1/billing/zalopay*';

  constructor(private readonly webhookService: WebhookService) {}

  @Post('callback')
  @Public()
  @Throttle(RATE_LIMIT.STRICT)
  async handleCallback(
    @Body() payload: WebhookRequestDto,
    @Res() res: Response,
  ) {
    const result = await this.webhookService.handleCallback(payload);

    res.status(200).json({
      return_code: result.returnCode,
      return_message: result.returnMessage,
    });
  }
}
