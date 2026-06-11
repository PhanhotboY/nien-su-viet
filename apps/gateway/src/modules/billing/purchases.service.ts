import { type ClientGrpc } from '@nestjs/microservices';
import { Inject, Injectable, Logger } from '@nestjs/common';

import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  PurchaseServiceClient,
  PURCHASE_SERVICE_NAME,
} from '@phanhotboy/genproto/billing_service/purchases';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { catchError, firstValueFrom, throwError, timeout } from 'rxjs';

@Injectable()
export class PurchasesService {
  private readonly microserviceErrorHandler: MicroserviceErrorHandler;
  private purchaseClient: PurchaseServiceClient;

  constructor(
    @Inject(GRPC_SERVICE.BILLING.NAME)
    private readonly client: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(logger);
    this.purchaseClient = this.client.getService<PurchaseServiceClient>(
      PURCHASE_SERVICE_NAME,
    );
  }

  async createPurchase(userId: string, payload: any) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.purchaseClient
            .createPurchase({
              ...payload,
              amount: {
                amount: payload.amount,
                currency: payload.currency,
              },
              userId,
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'create purchase',
      PurchasesService.name,
    );
  }
}
