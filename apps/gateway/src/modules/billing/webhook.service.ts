import { type ClientGrpc } from '@nestjs/microservices';
import { Inject, Injectable, Logger } from '@nestjs/common';

import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  ZaloPayWebhookServiceClient,
  ZALO_PAY_WEBHOOK_SERVICE_NAME,
} from '@phanhotboy/genproto/billing_service/webhook';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { catchError, firstValueFrom, throwError, timeout } from 'rxjs';
import { WebhookRequestDto } from './dto/webhook.req.dto';

@Injectable()
export class WebhookService {
  private readonly microserviceErrorHandler: MicroserviceErrorHandler;
  private webhookClient: ZaloPayWebhookServiceClient;

  constructor(
    @Inject(GRPC_SERVICE.BILLING.NAME)
    private readonly client: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(logger);
    this.webhookClient = this.client.getService<ZaloPayWebhookServiceClient>(
      ZALO_PAY_WEBHOOK_SERVICE_NAME,
    );
  }

  async handleCallback(payload: WebhookRequestDto) {
    this.logger.debug(`Received webhook callback with payload`, payload);
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.webhookClient.handleCallback(payload).pipe(
            timeout(10000),
            catchError((error) => throwError(() => error)),
          ),
        ),
      'handle callback',
      WebhookService.name,
    );
  }
}
