import { HistoricalEventDbEntity } from '@historical-event/modules/historical-event/domain/entity/historical-event.db.entity';
import { HistoricalEventEsEntity } from '@historical-event/modules/historical-event/domain/entity/historical-event.es.entity';
import { IHistoricalEventsEsRepo } from '@historical-event/modules/historical-event/domain/repo/historical-event.es.repo';
import { Injectable } from '@nestjs/common';
import { ElasticsearchService } from '@nestjs/elasticsearch';

@Injectable()
export class HistoricalEventsEsRepo implements IHistoricalEventsEsRepo {
  private readonly documentName = 'historical-events';

  constructor(private readonly esService: ElasticsearchService) {}

  async index(event: HistoricalEventDbEntity) {
    const document = new HistoricalEventEsEntity(event);
    return this.esService.index({
      index: this.documentName,
      id: event.id,
      document,
    });
  }
}
