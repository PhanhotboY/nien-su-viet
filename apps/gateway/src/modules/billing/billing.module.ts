import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';

import { Config } from '@gateway/config';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { ConfigService } from '@phanhotboy/nsv-common';
import { BILLING_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/billing_service/billing';
import { PurchasesController } from './purchases.controller';
import { PurchasesService } from './purchases.service';
import { PlansController } from './plans.controller';
import { PlansService } from './plans.service';
import { WebhookService } from './webhook.service';
import { WebhookController } from './webhook.controller';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: GRPC_SERVICE.BILLING.NAME,
        useFactory: (config: ConfigService<Config>) => ({
          transport: Transport.GRPC,
          options: {
            url: config.get('billingServiceEndpoint'),
            package: BILLING_SERVICE_PACKAGE_NAME,
            protoPath: GRPC_SERVICE.BILLING.PROTO_PATH,
            loader: {
              // includes imported proto files
              includeDirs: [GRPC_SERVICE.MAIN_PROTO_PATH],
              // since Unmarshal removes falsy values, we need to set default values back to avoid losing data, like empty array, false, etc.
              defaults: true,
              // The type to use to represent long (int64) values. Instead of Long object by default.
              longs: Number,
            },
          },
        }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [PurchasesController, PlansController, WebhookController],
  providers: [PurchasesService, PlansService, WebhookService],
  exports: [ClientsModule],
})
export class BillingModule {}
