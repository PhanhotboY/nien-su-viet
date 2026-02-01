import { RATE_LIMIT } from '@gateway/config';
import { Throttle, ThrottlerOptions } from '@nestjs/throttler';

export function ThrottleMap(options: typeof RATE_LIMIT.DEFAULT = []) {
  return Throttle(
    options.reduce(
      (acc, curr) => {
        return {
          [curr.name]: curr,
          ...acc,
        };
      },
      {} as Record<string, ThrottlerOptions>,
    ),
  );
}
