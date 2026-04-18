import { Inject, Injectable } from '@nestjs/common';
import { Counter, Histogram, type Meter } from '@opentelemetry/api';
import { METRICS } from '../metrics/metrics.constant';

@Injectable()
export class SharedMetrics {
  requests_total: Counter;
  request_duration_ms: Histogram;
  db_query_duration_seconds: Histogram;

  constructor(@Inject(METRICS.METER) private meter: Meter) {
    this.requests_total = this.meter.createCounter('requests_total', {
      description: 'Total number of requests received',
    });

    this.request_duration_ms = this.meter.createHistogram(
      'request_duration_ms',
      {
        description: 'Duration of requests in milliseconds',
        unit: 'ms',
      },
    );
  }
}
