import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { RmqService } from '@phanhotboy/nsv-common';
import { AuthService } from './auth';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, { bodyParser: false });
  const rmqService = app.get(RmqService);
  const auth = app.get(AuthService);
  app.connectMicroservice(rmqService.getOptions('auth_queue'));

  // Global prefix
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;

  app.enableShutdownHooks();
  await app.startAllMicroservices();
  await app.listen(port);
  console.log(
    `NestJS Auth service is running on port ${port} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
