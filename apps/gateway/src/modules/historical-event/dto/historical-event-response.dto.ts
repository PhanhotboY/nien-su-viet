import { Exclude, Expose, Transform, Type } from 'class-transformer';
import { OmitType, PickType } from '@nestjs/swagger';
import { HistoricalEventBaseDto } from './historical-event-base.dto';
import { EventCategoriesBriefResponseDto } from '@phanhotboy/nsv-common/dto/event-categories';
import { IsIn, IsString } from 'class-validator';
import { UserBriefResponseDto } from '@gateway/modules/auth/dto';
import { toEventDateType } from '@historical-event/helper/dateType.helper';
import { HISTORICAL_EVENT } from 'apps/client/src/constants/historical-event.constant';

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
    'thumbnail',
  ],
) {
  @Expose()
  @Transform(({ value }) => toEventDateType(value))
  @IsIn(Object.values(HISTORICAL_EVENT.EVENT_DATE_TYPE), {
    message: 'Loại ngày bắt đầu không hợp lệ.',
  })
  fromDateType!: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>;

  @Expose()
  @Transform(({ value }) => toEventDateType(value))
  @IsIn(Object.values(HISTORICAL_EVENT.EVENT_DATE_TYPE), {
    message: 'Loại ngày bắt đầu không hợp lệ.',
  })
  toDateType!: Values<typeof HISTORICAL_EVENT.EVENT_DATE_TYPE>;

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
