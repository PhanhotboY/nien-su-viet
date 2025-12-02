import { PickType } from '@nestjs/swagger';
import { Exclude, Expose } from 'class-transformer';
import { EventCategoryBaseDto } from './event-category-base.dto';

// DTO for creating event category
@Exclude()
export class EventCategoryBriefResponseDto extends PickType(
  EventCategoryBaseDto,
  ['id', 'name', 'slug'],
) {}

@Exclude()
export class EventCategoryResponseDto extends EventCategoryBriefResponseDto {
  @Expose()
  description?: string;
}
