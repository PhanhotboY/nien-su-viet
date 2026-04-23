import { context, Span, trace } from '@opentelemetry/api';

/**
 * Create a span in code to trace specific operations
 * @example
 * ```typescript
 * const { span, context: spanContext } = createSpan('database.prisma.query');
 * try {
 *   const result = await context.with(spanContext, async () => {
 *     return this.prisma.historicalEvent.findMany();
 *   });
 * } finally {
 *   span.end();
 * }
 * ```
 */
export function createSpan(
  tracerName: string,
  spanName: string,
  attributes?: Record<string, string | number | boolean>,
) {
  const tracer = trace.getTracer(tracerName);
  const span = tracer.startSpan(spanName, { attributes });
  const spanContext = trace.setSpan(context.active(), span);

  return { span, context: spanContext };
}

/**
 * Utility to add child spans within an operation
 * @example
 * ```typescript
 * const span = trace.getActiveSpan();
 * addSpanEvent(span, 'cache_hit', { cacheKey: 'events:all' });
 * ```
 */
export function addSpanAttributes(
  span: Span | undefined,
  attributes: Record<string, string | number | boolean>,
) {
  if (span) {
    span.setAttributes(attributes);
  }
}

/**
 * Utility to record timing of an operation
 * @example
 * ```typescript
 * const span = trace.getActiveSpan();
 * await recordOperationTiming(span, 'prisma.query', async () => {
 *   return this.prisma.historicalEvent.findMany();
 * });
 * ```
 */
export async function recordOperationTiming<T>(
  tracerName: string,
  operationName: string,
  operation: () => Promise<T>,
): Promise<T> {
  const { span } = createSpan(tracerName, operationName);
  try {
    const result = await operation();
    span.end();
    return result;
  } catch (error) {
    span.recordException(error);
    throw error;
  }
}

export function recordOperationTimingSync<T>(
  tracerName: string,
  operationName: string,
  operation: () => T,
): T {
  const { span } = createSpan(tracerName, operationName);
  try {
    const result = operation();
    span.end();
    return result;
  } catch (error) {
    span.recordException(error);
    throw error;
  }
}
