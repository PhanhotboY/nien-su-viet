import { OmitType } from '@nestjs/swagger';
import { Exclude, Expose } from 'class-transformer';
import { IsBoolean, IsOptional } from 'class-validator';
import { PostBaseDto } from './post-base.dto';

@Exclude()
export class PostBaseCreateDto extends OmitType(PostBaseDto, [
  'id',
  'views',
  'likes',
  'createdAt',
  'updatedAt',
  'published',
]) {
  @Expose()
  @IsOptional()
  @IsBoolean({ message: 'Trạng thái đăng không hợp lệ' })
  published?: PostBaseDto['published'];
}
