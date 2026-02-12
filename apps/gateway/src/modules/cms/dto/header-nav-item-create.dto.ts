import { OmitType } from '@nestjs/swagger';
import { Exclude } from 'class-transformer';
import { HeaderNavItemDto } from './header-nav-item.dto';

@Exclude()
export class HeaderNavItemCreateDto extends OmitType(HeaderNavItemDto, [
  'id',
] as const) {}
