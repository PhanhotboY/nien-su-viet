import { Inject, Injectable } from '@nestjs/common';
import { Counter, Histogram, type Meter } from '@opentelemetry/api';
import { METRICS } from '@phanhotboy/nsv-common/metrics/metrics.constant';

@Injectable()
export class GatewayMetrics {
  requests_total: Counter;
  // rate_limit_rejections_total: Counter;
  // cache_hits_total: Counter;
  request_duration_ms: Histogram;
  db_query_duration_seconds: Histogram;

  constructor(@Inject(METRICS.METER) private meter: Meter) {
    this.requests_total = this.meter.createCounter('requests_total', {
      description: 'Total number of requests received by the gateway',
    });
    // this.rate_limit_rejections_total = this.meter.createCounter('rate_limit_rejections_total', {
    //   description: 'Total number of requests rejected due to rate limiting',
    // });
    // this.cache_hits_total = this.meter.createCounter('cache_hits_total', {
    //   description: 'Total number of cache hits in the gateway',
    // });
    this.request_duration_ms = this.meter.createHistogram(
      'request_duration_ms',
      {
        description: 'Duration of requests handled by the gateway in seconds',
        unit: 'ms',
      },
    );
    // this.db_query_duration_seconds = this.meter.createHistogram(
    //   'db_query_duration_seconds',
    //   {
    //     description:
    //       'Duration of database queries made by the gateway in seconds',
    //   },
    // );
  }
}
