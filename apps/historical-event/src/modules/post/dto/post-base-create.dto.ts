import { OmitType } from '@nestjs/swagger';
import { PostBaseDto } from './post-base.dto';
import { Exclude, Expose } from 'class-transformer';
import { IsOptional, IsUUID } from 'class-validator';

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
  publishedAt?: PostBaseDto['publishedAt'];

  @Expose()
  @IsOptional()
  published?: PostBaseDto['published'];
}
