import { Exclude, Expose, Type } from 'class-transformer';
import { OmitType, PickType } from '@nestjs/swagger';
import { HistoricalEventBaseDto } from './historical-event-base.dto';
import { ImageBriefResponseDto } from '../image';
import { UserBaseResponseDto } from '../user/user-base-response.dto';
import { EventCategoriesBriefResponseDto } from '../event-categories';
import { IsString } from 'class-validator';

// DTO for response historical event
@Exclude()
export class HistoricalEventBriefResponseDto extends PickType(
  HistoricalEventBaseDto,
  [
    'id',
    'name',
    'fromDay',
    'fromMonth',
    'fromYear',
    'toDay',
    'toMonth',
    'toYear',
  ],
) {
  @Type(() => ImageBriefResponseDto)
  thumbnail?: ImageBriefResponseDto | null;

  @Expose()
  @Type(() => UserBaseResponseDto)
  author!: UserBaseResponseDto;
}

@Exclude()
export class HistoricalEventPreviewResponseDto extends HistoricalEventBriefResponseDto {
  @Expose()
  @Type(() => OmitType(EventCategoriesBriefResponseDto, ['event']))
  categories!: Omit<EventCategoriesBriefResponseDto, 'event'>[];

  @Expose()
  @IsString({ message: 'Trích đoạn không hợp lệ' })
  excerpt!: string;
}

@Exclude()
export class HistoricalEventDetailResponseDto extends OmitType(
  HistoricalEventPreviewResponseDto,
  ['excerpt'],
) {
  @Expose() content!: string;
}
