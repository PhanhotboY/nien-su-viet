import {
  Injectable,
  NestInterceptor,
  ExecutionContext,
  CallHandler,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';
import { SharedMetrics } from '../providers';
import {
  createHttpAttributes,
  createRpcAttributes,
} from '@phanhotboy/nsv-common';

@Injectable()
export class MetricsInterceptor implements NestInterceptor {
  constructor(private readonly sm: SharedMetrics) {}

  intercept(context: ExecutionContext, next: CallHandler): Observable<any> {
    const type = context.getType();
    const req = context.switchToHttp().getRequest();
    let start: number;
    let statusCode = 200;

    if (type !== 'http') {
      if (type === 'rpc') {
        // Not an HTTP request, skip metrics
        return next.handle().pipe(
          tap({
            subscribe: () => {
              start = Date.now();
            },
            error: (err) => {
              // check if RpcException
              if (err.error) {
                statusCode = err.error.statusCode || 500;
              } else {
                statusCode = err.status || 500;
              }
            },
            finalize: () => {
              const duration = Date.now() - start;
              const method = context.getHandler().name;
              const service = context.getClass().name.replace('Controller', '');

              this.sm.requests_total.add(
                1,
                createRpcAttributes({ service, method, statusCode }),
              );

              this.sm.request_duration_ms.record(
                duration, // Convert to seconds
                createRpcAttributes({ method, service, statusCode }),
              );
            },
          }),
        );
      }
      return next.handle();
    }

    return next.handle().pipe(
      tap({
        subscribe: () => {
          start = Date.now();
        },
        error: (err) => {
          statusCode = err.status || 500;
        },
        finalize: () => {
          const duration = Date.now() - start;
          const method = req.method;
          const route = req.route?.path || 'unknown_route';
          const handler = context.getHandler().name;

          this.sm.requests_total.add(
            1,
            createHttpAttributes({ method, handler, route, statusCode }),
          );

          this.sm.request_duration_ms.record(
            duration, // Convert to seconds
            createHttpAttributes({ method, handler, route, statusCode }),
          );
        },
      }),
    );
  }
}
