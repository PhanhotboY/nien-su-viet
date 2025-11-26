import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { middleware } from './app.middleware';
import { RmqService } from '@phanhotboy/nsv-common';
import { initSwagger } from '@phanhotboy/nsv-common/swagger';
import { auth } from './lib/auth';
import { existsSync, mkdirSync, writeFileSync } from 'fs';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, { bodyParser: false });
  // const rmqService = app.get(RmqService);
  // app.connectMicroservice(rmqService.getOptions('auth_queue'));

  middleware(app);
  // initSwagger(app, 'Auth Service', true);
  const document = await auth.api.generateOpenAPISchema();

  // Save OpenAPI JSON into monorepo
  if (existsSync('openapi') === false) {
    // Create the directory if it does not exist
    mkdirSync('openapi');
  }
  writeFileSync(`openapi/auth-service.json`, JSON.stringify(document, null, 2));

  // Global prefix
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;

  app.enableShutdownHooks();
  // await app.startAllMicroservices();
  await app.listen(port);
  console.log(
    `NestJS Auth service is running on port ${port} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
