import { OmitType } from '@nestjs/swagger';
import { PostBaseDto } from './post-base.dto';
import { Exclude, Expose, Transform, Type } from 'class-transformer';
import { IsBoolean, IsDate, IsInt, IsOptional } from 'class-validator';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { TimestampDto } from '@phanhotboy/nsv-common';

@Exclude()
export class PostBaseCreateDto extends OmitType(PostBaseDto, [
  'id',
  'authorId',
  'publishedAt',
  'published',
  'views',
  'likes',
  'createdAt',
  'updatedAt',
]) {
  // Required in schema but optional through API
  @Expose()
  @IsOptional()
  authorId?: PostBaseDto['authorId'];

  @Expose()
  @IsOptional()
  @IsInt({ message: 'Ngày xuất bản không hợp lệ' })
  publishedAt?: PostBaseDto['publishedAt'];

  @Expose()
  @IsOptional()
  @IsBoolean({ message: 'Trạng thái xuất bản không hợp lệ' })
  published?: PostBaseDto['published'];
}

@Exclude()
export class PostBaseCreateGrpcDto extends OmitType(PostBaseCreateDto, [
  'publishedAt',
]) {
  @Expose()
  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  publishedAt?: TimestampDto;
}
