import { HttpException, Logger } from '@nestjs/common';

import { mapGrpcCodeToHttpStatusCode } from '@gateway/helpers/grpc.helper';
import {
  cleanupErrorMessage,
  mapErrorMessageToStatusCode,
  mapErrorNameToStatusCode,
} from '@gateway/helpers/error.helper';

/**
 * Shared error handling utility for Gateway services
 * Properly handles RpcException from microservices and converts to HTTP exceptions
 */
export class MicroserviceErrorHandler {
  constructor(private readonly logger: Logger) {}

  /**
   * Handle microservice errors and convert to appropriate HTTP exceptions
   * Properly handles RpcException from microservices
   */
  handleError(
    error: any,
    operation: string,
    serviceName: string = 'Microservice',
  ): never {
    this.logger.error(`${serviceName} ${operation} failed:`, error);

    // Handle RpcException from microservice
    const rpcError = error?.response || error;

    // Extract status code and message from RpcException
    const statusCode = this.extractStatusCode(rpcError);
    const message = this.extractErrorMessage(rpcError, operation);

    this.logger.debug(
      `Throwing HttpException with status: ${statusCode}, message: ${message}`,
    );
    throw new HttpException(message, statusCode);
  }

  /**
   * Extract HTTP status code from RpcException or error object
   */
  extractStatusCode(error: any): number {
    // Priority 1: Direct statusCode from RpcException
    if (error.statusCode && typeof error.statusCode === 'number') {
      return error.statusCode;
    }

    // Priority 2: Go gRPC status code mapping
    if (error.code && typeof error.code === 'number') {
      return mapGrpcCodeToHttpStatusCode(error.code);
    }

    // Priority 3: Legacy status property
    if (error.status && typeof error.status === 'number') {
      return error.status;
    }

    // Priority 4: Check error name for NestJS exception types
    if (error.name) {
      return mapErrorNameToStatusCode(error.name);
    }

    // Priority 5: Parse error message for common patterns
    const errorMessage = this.extractErrorMessage(error).toLowerCase();
    return mapErrorMessageToStatusCode(errorMessage);
  }

  /**
   * Extract meaningful error message from RpcException or error object
   */
  private extractErrorMessage(error: any, operation?: string): string {
    // Priority 1: Go gRPC status.Error
    if (error.details && typeof error.details === 'string') {
      return cleanupErrorMessage(error.details);
    }

    // Priority 2: Direct message from RpcException
    if (error.message && typeof error.message === 'string') {
      return cleanupErrorMessage(error.message);
    }

    // Priority 3: Message array (for validation errors)
    if (Array.isArray(error.message)) {
      return error.message.join(', ');
    }

    // Priority 4: Error details from nested response
    if (error.response && error.response.message) {
      if (Array.isArray(error.response.message)) {
        return error.response.message.join(', ');
      }
      return cleanupErrorMessage(error.response.message);
    }

    // Priority 5: Error property
    if (error.error && typeof error.error === 'string') {
      return error.error;
    }

    // Priority 6: Default operation-specific message
    return 'Operation failed';
  }

  /**
   * Async wrapper for handling microservice calls with error handling
   */
  async handleAsyncCall<T>(
    operation: () => Promise<T>,
    operationName: string,
    serviceName: string = 'Microservice',
  ): Promise<T> {
    try {
      return await operation();
    } catch (error) {
      this.handleError(error, operationName, serviceName);
    }
  }
}
