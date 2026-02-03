import { PartialType } from '@nestjs/swagger';
import { HistoricalEventBaseCreateDto } from './historical-event-base-create.dto';

// DTO for updating historical event
export class HistoricalEventBaseUpdateDto extends PartialType(
  HistoricalEventBaseCreateDto,
) {}
