import { Module } from '@nestjs/common';

import { RmqModule } from '@phanhotboy/nsv-common';
import { ProcessedEventService } from './processEvent.service';
import { RMQ } from '@phanhotboy/constants';

@Module({
  imports: [RmqModule.register()],
  providers: [ProcessedEventService],
})
export class ProcessedEventModule {}
