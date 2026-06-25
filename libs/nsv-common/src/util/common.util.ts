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

function toCamelCase(str: string): string {
  return str
    .toLowerCase()
    .replace(/[-_\s]+(.)?/g, (_, char) => (char ? char.toUpperCase() : ''));
}

function keysToCamelCase(obj: any): any {
  if (Array.isArray(obj)) {
    return obj.map(keysToCamelCase);
  }

  if (obj && typeof obj === 'object') {
    return Object.fromEntries(
      Object.entries(obj).map(([key, value]) => [
        toCamelCase(key),
        keysToCamelCase(value),
      ]),
    );
  }

  return obj;
}

function keysToSnakeCase(obj: any): any {
  if (Array.isArray(obj)) {
    return obj.map(keysToSnakeCase);
  }

  if (obj && typeof obj === 'object') {
    return Object.fromEntries(
      Object.entries(obj).map(([key, value]) => [
        toSnakeCase(key),
        keysToSnakeCase(value),
      ]),
    );
  }

  return obj;
}

/**
 *
 * @description Calculate retry delay using exponential backoff with optional jitter.
 * @returns Delay in milliseconds.
 */
const calculateRetryDelay = function ({
  retryCount,
  baseDelay = 1000,
  maxDelay = 30000,
  jitter,
}: {
  retryCount: number;
  baseDelay?: number;
  maxDelay?: number;
  jitter?: boolean;
}) {
  // Calculate exponential delay: 2^retryCount * baseDelay
  let delay = Math.min(maxDelay, Math.pow(2, retryCount) * baseDelay);

  // Add jitter to prevent thundering herd problem
  if (jitter) {
    delay = delay * (0.5 + Math.random());
  }

  return delay;
};

export {
  genConfiguration,
  toSnakeCase,
  toCamelCase,
  keysToCamelCase,
  keysToSnakeCase,
  calculateRetryDelay,
};
