import { OmitType, PartialType } from '@nestjs/swagger';
import {
  PostBaseCreateDto,
  PostBaseCreateGrpcDto,
} from './post-base-create.dto';
import { Expose, Transform } from 'class-transformer';
import { IsOptional } from 'class-validator';
import { PostBaseDto } from './post-base.dto';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { TimestampDto } from '@phanhotboy/nsv-common';

// DTO for updating historical event
export class PostBaseUpdateDto extends PartialType(PostBaseCreateDto) {
  @Expose()
  @IsOptional()
  publishedAt?: PostBaseDto['publishedAt'];
}

export class PostBaseUpdateGrpcDto extends OmitType(PostBaseUpdateDto, [
  'publishedAt',
]) {
  @Expose()
  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  publishedAt?: TimestampDto;
}
