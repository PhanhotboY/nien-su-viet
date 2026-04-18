import { Module } from '@nestjs/common';
import { resourceFromAttributes } from '@opentelemetry/resources';
import {
  ATTR_SERVICE_NAME,
  ATTR_SERVICE_VERSION,
} from '@opentelemetry/semantic-conventions';
import { metrics } from '@opentelemetry/api';
import {
  MeterProvider,
  PeriodicExportingMetricReader,
} from '@opentelemetry/sdk-metrics';
import { OTLPMetricExporter } from '@opentelemetry/exporter-metrics-otlp-http';

import { METRICS } from './metrics.constant';
import { ConfigService } from '@nestjs/config';

@Module({
  providers: [
    {
      provide: METRICS.METER,
      useFactory: (config: ConfigService) => {
        const exporter = new OTLPMetricExporter({
          url:
            config.get<string>('OTLP_EXPORTER_ENDPOINT') ||
            'http://127.0.0.1:4318/v1/metrics',
        });
        const reader = new PeriodicExportingMetricReader({
          exporter,
          exportIntervalMillis: 3000, // Export every 3 seconds
        });
        const meterProvider = new MeterProvider({
          resource: resourceFromAttributes({
            [ATTR_SERVICE_NAME]:
              config.get<string>('SERVICE_NAME') || 'nsv-service',
            [ATTR_SERVICE_VERSION]:
              config.get<string>('SERVICE_VERSION') || 'latest',
          }),
          readers: [reader],
        });
        metrics.setGlobalMeterProvider(meterProvider);

        return meterProvider.getMeter(
          config.get<string>('METRICS_INSTRUMENTATION_NAME') ||
            'io.opentelemetry.metrics.js',
        );
      },
      inject: [ConfigService],
    },
  ],
  exports: [METRICS.METER],
})
export class MetricsModule {}
