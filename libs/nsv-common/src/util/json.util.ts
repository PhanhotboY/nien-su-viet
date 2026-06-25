import { RpcException } from '@nestjs/microservices';
import { ClassConstructor, plainToInstance } from 'class-transformer';
import { keysToCamelCase } from './common.util';

function validateJson<T>(dto: ClassConstructor<T>, json: string): T {
  try {
    const parsed = JSON.parse(json);
    return plainToInstance(dto, keysToCamelCase(parsed));
  } catch (error) {
    throw new RpcException(`Invalid JSON: ${error.message}`);
  }
}

export { validateJson };
