import { HISTORICAL_EVENT } from '@phanhotboy/constants';

export class HistoricalEventEsEntity {
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
  createdAt: Date;
  updatedAt: Date;

  constructor(partial: Partial<HistoricalEventEsEntity>) {
    Object.assign(this, partial);
  }
}
