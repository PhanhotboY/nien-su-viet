const isObject = function <T>(value: T): value is T & Object {
  return value != null && typeof value === 'object' && !Array.isArray(value);
};

const merge = function <T extends Object, U extends Object>(
  target: T,
  source: U,
): T & U {
  for (const key of Object.keys(source)) {
    const targetValue = target[key];
    const sourceValue = source[key] as any;
    if (isObject(targetValue) && isObject(sourceValue)) {
      Object.assign(sourceValue, merge(targetValue, sourceValue));
    }
  }

  return { ...target, ...source };
};

const genConfiguration = function <D extends Object, T extends Object>(
  defaultConfig: D,
  configs: Record<string, T>,
): () => D & T {
  return () => {
    const env = process.env.NODE_ENV || 'development';
    const envConfig = configs[env];
    if (!envConfig) {
      throw new Error(`Configuration for environment "${env}" is not defined.`);
    }

    return merge(defaultConfig, envConfig);
  };
};

const toSnakeCase = (str: string) => {
  if (!str) return '';
  return (
    str
      .match(
        /[A-Z]{2,}(?=[A-Z][a-z]+[0-9]*|\b)|[A-Z]?[a-z]+[0-9]*|[A-Z]|[0-9]+/g,
      )
      ?.map((word) => word.toLowerCase())
      .join('_') || ''
  );
};

const getRoutingKey = function (event: Function) {
  return toSnakeCase(event.name);
};

const getQueueName = function (event: Function) {
  return `${toSnakeCase(event.name)}_queue`;
};

export const commonUtils = { genConfiguration, getRoutingKey, getQueueName };
