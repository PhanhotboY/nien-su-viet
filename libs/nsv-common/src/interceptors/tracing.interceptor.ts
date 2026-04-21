import {
  Injectable,
  NestInterceptor,
  ExecutionContext,
  CallHandler,
  Inject,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { trace, type Tracer } from '@opentelemetry/api';
import { createHttpAttributes, createRpcAttributes } from '../otel';

@Injectable()
export class TracingInterceptor implements NestInterceptor {
  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const span = trace.getActiveSpan();

    const type = context.getType();
    let statusCode = 200;
    if (type !== 'http') {
      if (type === 'rpc') {
        // Not an HTTP request, skip metrics
        return next.handle().pipe(
          tap({
            error: (err) => {
              span?.recordException(err);
              // check if RpcException
              if (err.error) {
                statusCode = err.error.statusCode || 500;
              } else {
                statusCode = err.status || 500;
              }
            },
            finalize: () => {
              const method = context.getHandler().name;
              const service = context.getClass().name.replace('Controller', '');

              span?.setAttributes(
                createRpcAttributes({ service, method, statusCode }),
              );
            },
          }),
        );
      }
      return next.handle();
    }

    const req = context.switchToHttp().getRequest();
    return next.handle().pipe(
      tap({
        error: (err) => {
          statusCode = err.status || 500;
          span?.recordException(err);
        },
        finalize: () => {
          const method = req.method;
          const route = req.route?.path || 'unknown_route';
          const handler = context.getHandler().name;

          span?.setAttributes(
            createHttpAttributes({ method, handler, route, statusCode }),
          );
        },
      }),
    );
  }
}
