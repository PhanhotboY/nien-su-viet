import { Exclude, Expose, Type } from 'class-transformer';
import { EventCategoryBriefResponseDto } from '../event-category';
// import { HistoricalEventBriefResponseDto } from '../historical-event';

@Exclude()
export class EventCategoriesBriefResponseDto {
  @Expose()
  eventId!: string;

  @Expose()
  categoryId!: string;

  @Expose()
  @Type(() => EventCategoryBriefResponseDto)
  category!: EventCategoryBriefResponseDto;

  // @Expose()
  // @Type(() => HistoricalEventBriefResponseDto)
  // event!: HistoricalEventBriefResponseDto;
}
