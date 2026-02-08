import { NestFactory } from '@nestjs/core';
import { GatewayModule } from './gateway.module';
import { middleware } from './gateway.middleware';
import { initSwagger } from '@phanhotboy/nsv-common/swagger';
import { HistoricalEventBriefResponseDto } from './modules/historical-event/dto';
import { PaginatedResponseDto } from '@phanhotboy/nsv-common';

async function bootstrap() {
  const app = await NestFactory.create(GatewayModule, { bodyParser: false });

  middleware(app);

  // Global prefix
  app.setGlobalPrefix('/api/v1');

  const port = process.env.NODE_PORT || 3000;
  initSwagger({
    app,
    name: 'NSV API Documentation',
    isStartEndpoint: true,
    extraModels: [PaginatedResponseDto],
  });

  app.enableShutdownHooks();
  await app.listen(port);
  console.log(
    `NestJS Gateway service is running on port ${port} in ${process.env.NODE_ENV} mode`,
  );
}
bootstrap();
