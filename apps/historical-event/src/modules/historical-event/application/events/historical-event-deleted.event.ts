import { HistoricalEventDeletedEvent as DeletedEvent } from '@phanhotboy/genproto/events/historical_event_events';

export class HistoricalEventDeletedEvent implements DeletedEvent {
  id: string;
  authorId: string;

  constructor(partial: Partial<DeletedEvent>) {
    Object.assign(this, partial);
  }
}
