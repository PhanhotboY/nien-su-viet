import { HistoricalEventDbEntity } from '../entity/historical-event.db.entity';

export interface IHistoricalEventDbRepo {
  createHistoricalEvent(payload: any): Promise<HistoricalEventDbEntity>;
  getHistoricalEventById(options: any): Promise<HistoricalEventDbEntity | null>;
  updateHistoricalEvent(id: string, updatedData: any): Promise<string>;
  deleteHistoricalEvent(id: string): Promise<void>;
  searchHistoricalEvents(query: any): Promise<HistoricalEventDbEntity[]>;
  countHistoricalEvents(query: any): Promise<number>;
}
