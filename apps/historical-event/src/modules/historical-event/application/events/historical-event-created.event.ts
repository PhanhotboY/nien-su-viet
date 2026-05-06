import { HistoricalEventCreatedEvent as CreatedEvent } from '@phanhotboy/genproto/events/historical_event_events';

export class HistoricalEventCreatedEvent implements CreatedEvent {
  id: string;
  authorId: string;
  name: string;
  thumbnail?: string | undefined;
  fromDateType: string;
  fromDay?: number | undefined;
  fromMonth?: number | undefined;
  fromYear: number;
  toDateType?: string | undefined;
  toDay?: number | undefined;
  toMonth?: number | undefined;
  toYear?: number | undefined;
  content: string;
  createdAt: string;
  updatedAt: string;

  constructor(partial: Partial<CreatedEvent>) {
    Object.assign(this, partial);
  }
}
