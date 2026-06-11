import { type ClientGrpc } from '@nestjs/microservices';
import { Inject, Injectable, Logger } from '@nestjs/common';

import { MicroserviceErrorHandler } from '@gateway/common/microservice-error.handler';
import {
  PLAN_SERVICE_NAME,
  PlanServiceClient,
} from '@phanhotboy/genproto/billing_service/plans';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { catchError, firstValueFrom, throwError, timeout } from 'rxjs';

@Injectable()
export class PlansService {
  private readonly microserviceErrorHandler: MicroserviceErrorHandler;
  private planClient: PlanServiceClient;

  constructor(
    @Inject(GRPC_SERVICE.BILLING.NAME)
    private readonly client: ClientGrpc,
    private readonly logger: Logger,
  ) {
    this.microserviceErrorHandler = new MicroserviceErrorHandler(logger);
    this.planClient =
      this.client.getService<PlanServiceClient>(PLAN_SERVICE_NAME);
  }

  async createPlan(userId: string, payload: any) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.planClient
            .createPlan({
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
      'create plan',
      PlansService.name,
    );
  }

  async listPlans(query: any) {
    return this.microserviceErrorHandler.handleAsyncCall(
      () =>
        firstValueFrom(
          this.planClient
            .listPlans({
              onlyActive: query.onlyActive === 'true',
              query: {
                ...query,
              },
            })
            .pipe(
              timeout(10000),
              catchError((error) => throwError(() => error)),
            ),
        ),
      'list plans',
      PlansService.name,
    );
  }
}
