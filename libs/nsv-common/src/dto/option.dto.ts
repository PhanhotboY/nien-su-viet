export interface CommonOptionDto {
  configuration: () => Record<string, any>;
  global?: boolean;
  throttlerConfigKey: string;
  redisConfigKey: string;
  rabbitmqConfigKey: string;
  cachePrefix: string;
  useSerializeInterceptors?: boolean;
  useCacheInterceptor?: boolean;
}
