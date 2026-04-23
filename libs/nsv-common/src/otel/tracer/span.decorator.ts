import {
  CallHandler,
  ExecutionContext,
  NestInterceptor,
  UseInterceptors,
} from '@nestjs/common';
import { trace, type Span } from '@opentelemetry/api';
import { map, Observable } from 'rxjs';

/**
 * Decorator to create a span for a method.
 * Useful for tracing database operations and other slow operations.
 *
 * @param spanName - Name of the span
 * @param attributes - Optional attributes to add to the span
 *
 * @example
 * ```typescript
 * @WithSpan('database.query.historical_events')
 * async getEvents(query: HistoricalEventQueryDto) {
 *   // Your code here
 * }
 * ```
 */
export function WithSpan(
  tracerName: string,
  spanName: string,
  attributes?: Record<string, string | number | boolean>,
) {
  class SpanInterceptor implements NestInterceptor {
    intercept(_ctx: ExecutionContext, next: CallHandler): Observable<any> {
      return trace
        .getTracer(tracerName)
        .startActiveSpan(spanName, (span: Span) => {
          if (attributes) {
            Object.entries(attributes).forEach(([key, value]) =>
              span.setAttribute(key, value),
            );
          }

          return next.handle().pipe(
            map((data) => {
              span.end();
              return data;
            }),
          );
        });
    }
  }
  return UseInterceptors(SpanInterceptor);
}
