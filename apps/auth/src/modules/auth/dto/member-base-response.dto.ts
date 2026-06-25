import { PickType } from '@nestjs/swagger';
import { MemberBaseDto } from './member-base.dto';
import { Expose, Type } from 'class-transformer';
import { UserBriefResponseDto } from './user-base-response.dto';

export class MemberBriefResponseDto extends MemberBaseDto {
  @Expose()
  @Type(() => UserBriefResponseDto)
  user: UserBriefResponseDto;
}
