import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { RmqService } from '@phanhotboy/nsv-common';
import { existsSync, mkdirSync, writeFileSync } from 'fs';
import { AuthService } from './auth';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, { bodyParser: false });
  const rmqService = app.get(RmqService);
  const auth = app.get(AuthService);
  app.connectMicroservice(rmqService.getOptions('auth_queue'));

  // Generate Better Auth OpenAPI schema (contains all auth routes)
  const document = await auth.api.generateOpenAPISchema();
  for (const path in document.paths) {
    for (const method in document.paths[path]) {
      const op = document.paths[path][method];
      // rename operationId
      op.operationId = `${method}_${path.replace(/\W+/g, '_')}`;
    }
  }

  // Save OpenAPI JSON into monorepo
  if (existsSync('openapi') === false) {
    mkdirSync('openapi');
  }
  writeFileSync(`openapi/auth-service.json`, JSON.stringify(document, null, 2));

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
