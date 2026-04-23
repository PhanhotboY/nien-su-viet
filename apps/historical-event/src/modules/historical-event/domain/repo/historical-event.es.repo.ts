import { IndexResponse } from '@elastic/elasticsearch/lib/api/types';
import { HistoricalEventDbEntity } from '../entity/historical-event.db.entity';

export interface IHistoricalEventsEsRepo {
  index(event: HistoricalEventDbEntity): Promise<IndexResponse>;
}
