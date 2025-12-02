import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
// import { middleware } from './app.middleware';
import { RmqService } from '@phanhotboy/nsv-common';
import { initSwagger } from '@phanhotboy/nsv-common/swagger';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const rmqService = app.get(RmqService);
  app.connectMicroservice(rmqService.getOptions('historical_event_queue'));

  //ddleware(app);
  initSwagger(app, 'Historical Event', true);

  // Global prefix
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;

  app.enableShutdownHooks();
  await app.startAllMicroservices();
  await app.listen(port);
  console.log(
    `NestJS Historical Event service is running on port ${port} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
