export interface CommonOptionDto {
  configuration: () => Record<string, any>;
  global?: boolean;
  throttlerConfigKey: string;
  redisConfigKey: string;
  rabbitmqConfigKey: string;
  cachePrefix: string;
  useInterceptors?: boolean;
  useCache?: boolean;
}
