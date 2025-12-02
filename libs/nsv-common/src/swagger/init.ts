import { existsSync, mkdirSync, writeFileSync } from 'fs';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { INestApplication } from '@nestjs/common';
import * as AuthDto from '../dto';

export function initSwagger(
  app: INestApplication,
  name: string,
  isStartEndpoint = false,
) {
  const config = new DocumentBuilder()
    .setTitle(`NSV - ${name} API`)
    .setDescription(`The ${name} API description`)
    .setVersion('1.0')
    .build();

  const document = SwaggerModule.createDocument(app, config, {
    extraModels: Object.values(AuthDto),
  });

  // Save OpenAPI JSON into monorepo
  if (existsSync('openapi') === false) {
    // Create the directory if it does not exist
    mkdirSync('openapi');
  }
  writeFileSync(
    `openapi/${name.split(' ').join('-').toLowerCase()}.json`,
    JSON.stringify(document, null, 2),
  );
  if (isStartEndpoint) {
    // Setup Swagger endpoint
    SwaggerModule.setup('api', app, document);
  }
}
