import { HISTORICAL_EVENT } from '@phanhotboy/constants';

export class HistoricalEventDbEntity {
  id: string;
  name: string;
  thumbnail?: string | undefined;
  fromDateType: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>;
  fromDay?: number;
  fromMonth?: number;
  fromYear: number;
  toDateType?: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>;
  toDay?: number | undefined;
  toMonth?: number | undefined;
  toYear?: number | undefined;
  authorId?: string;
  categories: string[];
  content: string;
  createdAt: Date;
  updatedAt: Date;

  constructor(partial: Partial<HistoricalEventDbEntity>) {
    Object.assign(this, partial);
  }
}
