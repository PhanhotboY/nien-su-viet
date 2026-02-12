import { OmitType } from '@nestjs/swagger';
import { Exclude } from 'class-transformer';
import { FooterNavItemDto } from './footer-nav-item.dto';

@Exclude()
export class FooterNavItemCreateDto extends OmitType(FooterNavItemDto, [
  'id',
] as const) {}
