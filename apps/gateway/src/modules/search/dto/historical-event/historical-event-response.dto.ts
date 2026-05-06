import { Exclude, Expose, Type } from 'class-transformer';
import { OmitType, PickType } from '@nestjs/swagger';
import { HistoricalEventBaseDto } from './historical-event-base.dto';
import { EventCategoriesBriefResponseDto } from '@phanhotboy/nsv-common/dto/event-categories';
import { IsString } from 'class-validator';
import { UserBriefResponseDto } from '@gateway/modules/auth/dto';

// DTO for response historical event
@Exclude()
export class HistoricalEventBriefResponseDto extends PickType(
  HistoricalEventBaseDto,
  [
    'id',
    'name',
    'fromDateType',
    'fromDay',
    'fromMonth',
    'fromYear',
    'toDateType',
    'toDay',
    'toMonth',
    'toYear',
    'thumbnail',
  ],
) {
  @Expose()
  @Type(() => UserBriefResponseDto)
  author!: UserBriefResponseDto;
}

@Exclude()
export class HistoricalEventPreviewResponseDto extends HistoricalEventBriefResponseDto {
  @Expose()
  // @Type(() => OmitType(EventCategoriesBriefResponseDto, ['event']))
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
