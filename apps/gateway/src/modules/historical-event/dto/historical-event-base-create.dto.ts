import { OmitType } from '@nestjs/swagger';
import { HistoricalEventBaseDto } from './historical-event-base.dto';
import { Exclude, Expose } from 'class-transformer';
import { IsOptional, IsUUID } from 'class-validator';

@Exclude()
export class HistoricalEventBaseCreateDto extends OmitType(
  HistoricalEventBaseDto,
  ['id', 'createdAt', 'updatedAt', 'authorId'],
) {
  // Required in schema but optional through API
  @Expose()
  @IsOptional()
  @IsUUID('4', { message: 'ID tác giả không hợp lệ' })
  authorId?: HistoricalEventBaseDto['authorId'];
}
