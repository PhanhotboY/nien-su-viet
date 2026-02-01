import { NestFactory } from '@nestjs/core';
import { TcpOptions, Transport } from '@nestjs/microservices';

import { AppModule } from './app.module';
import { RmqService } from '@phanhotboy/nsv-common';
import { TCP_SERVICE } from '@phanhotboy/constants/tcp-service.constant';

async function bootstrap() {
  const tcpPort = TCP_SERVICE.HISTORICAL_EVENT.PORT;
  const app = await NestFactory.create(AppModule);

  const rmqService = app.get(RmqService);
  app.connectMicroservice(rmqService.getOptions('historical_event_queue'));
  app.connectMicroservice<TcpOptions>({
    transport: Transport.TCP,
    options: {
      port: tcpPort,
    },
  });

  app.enableShutdownHooks();
  await app.startAllMicroservices();
  await app.init(); // Init the Nest application manually since not using .listen()

  console.log(
    `NestJS Historical Event service is running on port ${tcpPort} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
