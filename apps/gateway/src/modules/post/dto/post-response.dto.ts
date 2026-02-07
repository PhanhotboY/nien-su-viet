import { Exclude, Expose, Type } from 'class-transformer';
import { OmitType, PickType } from '@nestjs/swagger';
import { PostBaseDto } from './post-base.dto';
import { UserBaseResponseDto } from '@gateway/modules/auth/dto';

// DTO for response historical event
@Exclude()
export class PostBriefResponseDto extends PickType(PostBaseDto, [
  'id',
  'thumbnail',
  'title',
  'slug',
  'summary',
  'publishedAt',
  'published',
  'createdAt',
  'updatedAt',
]) {
  @Expose()
  @Type(() => UserBaseResponseDto)
  author!: UserBaseResponseDto;
}

@Exclude()
export class PostDetailResponseDto extends PostBriefResponseDto {
  @Expose() content!: string;
}
