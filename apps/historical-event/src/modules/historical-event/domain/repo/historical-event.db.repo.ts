import { Injectable } from '@nestjs/common';
import { HistoricalEventDbEntity } from '../entity/historical-event.db.entity';

export abstract class IHistoricalEventDbRepo {
  abstract createHistoricalEvent(
    payload: any,
  ): Promise<HistoricalEventDbEntity>;
  abstract getHistoricalEventById(
    options: any,
  ): Promise<HistoricalEventDbEntity | null>;
  abstract updateHistoricalEvent(id: string, updatedData: any): Promise<string>;
  abstract deleteHistoricalEvent(id: string): Promise<void>;
  abstract searchHistoricalEvents(
    query: any,
  ): Promise<HistoricalEventDbEntity[]>;
  abstract countHistoricalEvents(query: any): Promise<number>;
}
