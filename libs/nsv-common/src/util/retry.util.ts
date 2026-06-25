import { Injectable } from '@nestjs/common';
import { RpcException } from '@nestjs/microservices';

import { calculateRetryDelay } from '../util';

interface RetryOptions {
  baseDelay?: number;
  maxDelay?: number;
  maxRetries?: number;
  jitter?: boolean;
}
class ExponentialBackoffRetry {
  private baseDelay: number;
  private maxDelay: number;
  private maxRetries: number;
  private jitter: boolean;

  constructor(options: RetryOptions = {}) {
    this.baseDelay = options.baseDelay || 1000;
    this.maxDelay = options.maxDelay || 30000;
    this.maxRetries = options.maxRetries || 5;
    this.jitter = options.jitter || true;
  }

  async execute(fn: Function) {
    let retries = 0;

    while (true) {
      try {
        return await fn();
      } catch (error) {
        if (retries >= this.maxRetries) {
          throw new Error(`Failed after ${retries} retries: ${error.message}`);
        }

        const delay = calculateRetryDelay({
          retryCount: retries,
          baseDelay: this.baseDelay,
          maxDelay: this.maxDelay,
          jitter: this.jitter,
        });
        await this.wait(delay);
        retries++;
      }
    }
  }

  wait(ms: number) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
}

interface CircuitBreakerOptions {
  failureThreshold?: number;
  resetTimeout?: number;
}
class CircuitBreaker {
  private failureThreshold: number;
  private resetTimeout: number;
  private failures: number;
  private state: 'CLOSED' | 'OPEN' | 'HALF_OPEN';
  private lastFailureTime: number | null;

  constructor(options: CircuitBreakerOptions = {}) {
    this.failureThreshold = options.failureThreshold || 5;
    this.resetTimeout = options.resetTimeout || 60000;
    this.failures = 0;
    this.state = 'CLOSED';
    this.lastFailureTime = null;
  }

  async execute(fn: Function) {
    if (this.state === 'OPEN') {
      if (Date.now() - (this.lastFailureTime || 0) >= this.resetTimeout) {
        this.state = 'HALF_OPEN';
      } else {
        throw new Error('Circuit breaker is OPEN');
      }
    }

    try {
      const result = await fn();
      if (this.state === 'HALF_OPEN') {
        this.state = 'CLOSED';
        this.failures = 0;
      }
      return result;
    } catch (error) {
      this.failures++;
      this.lastFailureTime = Date.now();

      if (this.failures >= this.failureThreshold) {
        this.state = 'OPEN';
      }
      throw error;
    }
  }
}

interface RetrySystemOptions {
  retry?: RetryOptions;
  circuitBreaker?: CircuitBreakerOptions;
  logger?: any;
}
@Injectable()
class RetrySystem {
  private retrier: ExponentialBackoffRetry;
  private circuitBreaker: CircuitBreaker;
  private logger: any;

  constructor(options: RetrySystemOptions = {}) {
    this.retrier = new ExponentialBackoffRetry(options.retry);
    this.circuitBreaker = new CircuitBreaker(options.circuitBreaker);
    this.logger = options.logger || console;
  }

  async execute<T>(fn: () => T, context = {}): Promise<T> {
    const startTime = Date.now();
    let attempts = 0;

    try {
      return await this.circuitBreaker.execute(async () => {
        return await this.retrier.execute(async () => {
          attempts++;
          try {
            const result = await fn();
            this.logSuccess(context, attempts, startTime);
            return result;
          } catch (error) {
            this.logFailure(context, attempts, error);
            throw error;
          }
        });
      });
    } catch (error) {
      throw new RetryException(error, attempts, Date.now() - startTime);
    }
  }

  logSuccess(context: any, attempts: number, startTime: number) {
    this.logger.log({
      event: 'retry_success',
      context,
      attempts,
      duration: Date.now() - startTime,
    });
  }

  logFailure(context: any, attempts: number, error: Error) {
    this.logger.error({
      event: 'retry_failure',
      context,
      attempts,
      error: error.message,
    });
  }
}

class RetryException extends RpcException {
  constructor(
    private originalError: Error,
    private attempts: number,
    private duration: number,
  ) {
    super(originalError.message);
    this.name = 'RetryError';
  }
}

export { ExponentialBackoffRetry, CircuitBreaker, RetrySystem, RetryException };
