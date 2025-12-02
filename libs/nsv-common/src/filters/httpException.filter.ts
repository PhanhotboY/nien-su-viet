import {
  ArgumentsHost,
  Catch,
  HttpException,
  HttpStatus,
  Logger,
  ExceptionFilter,
} from '@nestjs/common';
import { JsonWebTokenError } from '@nestjs/jwt';
import { Response } from 'express';

@Catch(HttpException)
export class HttpExceptionsFilter implements ExceptionFilter {
  private readonly logger: Logger = new Logger();

  public catch(exception: HttpException, host: ArgumentsHost) {
    const req = host.switchToHttp().getRequest<Request>();
    const response = host.switchToHttp().getResponse<Response>();
    // Retrieve the exception from passport Strategy (if any)
    // @ts-ignore
    const err = req.authInfo || exception;
    const args = req.body;

    let status = this.getHttpStatus(err);
    if (status === HttpStatus.INTERNAL_SERVER_ERROR) {
      if (err instanceof JsonWebTokenError) {
        status = HttpStatus.UNAUTHORIZED;
        this.logger.log(err.message, 'JWTError');
      } else {
        // Error Notifications
        this.logger.error('UnhandledException', err);
      }
    }

    return response.status(status).json({
      statusCode: status,
      timestamp: new Date().toISOString(),
      path: req.url,
      message: err.message || null,
    });
  }

  private getHttpStatus(exception: unknown): HttpStatus {
    return exception instanceof HttpException
      ? exception.getStatus()
      : HttpStatus.INTERNAL_SERVER_ERROR;
  }
}
