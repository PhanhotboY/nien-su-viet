import { HistoricalEventUpdatedEvent as UpdatedEvent } from '@phanhotboy/genproto/events/historical_event_events';

export class HistoricalEventUpdatedEvent implements UpdatedEvent {
  id: string;
  name?: string | undefined;
  thumbnail?: string | undefined;
  fromDateType?: string | undefined;
  fromDay?: number | undefined;
  fromMonth?: number | undefined;
  fromYear?: number | undefined;
  toDateType?: string | undefined;
  toDay?: number | undefined;
  toMonth?: number | undefined;
  toYear?: number | undefined;
  content?: string | undefined;

  constructor(partial: Partial<UpdatedEvent>) {
    Object.assign(this, partial);
  }
}
