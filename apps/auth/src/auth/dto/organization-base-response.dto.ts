import { PickType } from '@nestjs/swagger';
import { OrganizationBaseDto } from './organization-base.dto';

export class OrganizationBriefResponseDto extends PickType(
  OrganizationBaseDto,
  ['id', 'name', 'members', 'slug', 'createdAt'],
) {}
