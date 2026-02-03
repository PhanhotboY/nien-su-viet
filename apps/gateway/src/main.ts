import { NestFactory } from '@nestjs/core';
import { GatewayModule } from './gateway.module';
import { middleware } from './gateway.middleware';
import { initSwagger } from '@phanhotboy/nsv-common/swagger';

async function bootstrap() {
  const app = await NestFactory.create(GatewayModule, { bodyParser: false });

  middleware(app);

  // Global prefix
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;
  initSwagger(app, 'NSV API Documentation', true);

  app.enableShutdownHooks();
  await app.listen(port);
  console.log(
    `NestJS Gateway service is running on port ${port} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
