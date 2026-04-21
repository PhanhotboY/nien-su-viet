import { NestFactory } from '@nestjs/core';
import { GrpcOptions, Transport } from '@nestjs/microservices';

import { AppModule } from './app.module';
import { RmqService } from '@phanhotboy/nsv-common';
import { GRPC_SERVICE } from '@phanhotboy/constants';
import { HISTORICAL_EVENT_SERVICE_PACKAGE_NAME } from '@phanhotboy/genproto/historical_event_service/historical_events';

async function bootstrap() {
  const grpcPort = process.env.GRPC_PORT || 50051;
  const app = await NestFactory.create(AppModule);

  const rmqService = app.get(RmqService);
  app.connectMicroservice(rmqService.getOptions('historical_event_queue'));

  const grpcServerOptions: GrpcOptions = {
    transport: Transport.GRPC,
    options: {
      url: `0.0.0.0:${grpcPort}`,
      package: HISTORICAL_EVENT_SERVICE_PACKAGE_NAME,
      protoPath: GRPC_SERVICE.HISTORICAL_EVENT.PROTO_PATH,
      loader: {
        // includes imported proto files
        includeDirs: [GRPC_SERVICE.MAIN_PROTO_PATH],
        // since Unmarshal removes falsy values, we need to set default values back to avoid losing data, like empty array, false, etc.
        defaults: true,
        // The type to use to represent long (int64) values. Instead of Long object by default.
        longs: Number,
      },
    },
  };

  app.connectMicroservice<GrpcOptions>(grpcServerOptions, {
    inheritAppConfig: true,
  });

  app.enableShutdownHooks();
  await app.startAllMicroservices();
  await app.init(); // Init the Nest application manually since not using .listen()

  console.log(
    `NestJS Historical Event service is running on port ${grpcPort} (gRPC) in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
