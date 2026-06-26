import { Module } from '@nestjs/common';

import { ProcessedEventService } from './processEvent.service';

@Module({
  providers: [ProcessedEventService],
})
export class ProcessedEventModule {}
