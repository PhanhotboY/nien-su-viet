import { Exclude, Expose, Transform, Type } from 'class-transformer';
import { ApiProperty, OmitType, PickType } from '@nestjs/swagger';
import { PostBaseDto } from './post-base.dto';
import { TimestampDto } from '@phanhotboy/nsv-common';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';

// DTO for response historical event
@Exclude()
export class PostBriefResponseDto extends PickType(PostBaseDto, [
  'id',
  'thumbnail',
  'title',
  'slug',
  'summary',
  'published',
  'authorId',
]) {
  @Expose()
  @ApiProperty()
  @Transform(({ value }: { value: TimestampDto }) =>
    TimestampUtil.toDate(value)?.toISOString(),
  )
  updatedAt: string;

  @Expose()
  @ApiProperty()
  @Transform(({ value }: { value: TimestampDto }) =>
    TimestampUtil.toDate(value)?.toISOString(),
  )
  createdAt: string;
}

@Exclude()
export class PostDetailResponseDto extends PostBriefResponseDto {
  @Expose() content!: string;
}
