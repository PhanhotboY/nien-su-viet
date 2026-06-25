import { PickType } from '@nestjs/swagger';
import { UserBaseDto } from './user-base.dto';

export class UserBriefResponseDto extends PickType(UserBaseDto, [
  'id',
  'name',
  'image',
]) {}
