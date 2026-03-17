import { Exclude, Expose, Type } from 'class-transformer';
import { ApiProperty, OmitType, PickType } from '@nestjs/swagger';
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
  'published',
]) {
  @Expose()
  @ApiProperty()
  @Type(() => UserBaseResponseDto)
  author!: UserBaseResponseDto;
}

@Exclude()
export class PostDetailResponseDto extends PostBriefResponseDto {
  @Expose() content!: string;
}
