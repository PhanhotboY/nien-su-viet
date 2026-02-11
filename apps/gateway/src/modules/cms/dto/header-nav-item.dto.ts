import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsInt,
  IsOptional,
  IsBoolean,
  IsEnum,
  IsUUID,
  MinLength,
  MaxLength,
  Min,
  Max,
} from 'class-validator';
import { LinkType } from './common.dto';

// Header Nav Item Data DTO
@Exclude()
export class HeaderNavItemDto {
  @Expose()
  @IsUUID('4', { message: 'ID không hợp lệ' })
  id!: string;

  @Expose()
  @IsInt({ message: 'Thứ tự không hợp lệ' })
  @Min(0, { message: 'Thứ tự phải lớn hơn hoặc bằng 0' })
  @Max(1000, { message: 'Thứ tự không được quá 1000' })
  @Transform(({ value }) => (value !== undefined ? parseInt(value, 10) : value))
  order!: number;

  @Expose()
  @IsOptional()
  @IsBoolean({ message: 'Link new tab không hợp lệ' })
  link_new_tab?: boolean | null;

  @Expose()
  @IsString({ message: 'Link URL không hợp lệ' })
  @MinLength(1, { message: 'Link URL là bắt buộc' })
  @MaxLength(2048, { message: 'Link URL không được quá 2048 ký tự' })
  @Transform(({ value }) => value?.trim())
  link_url!: string;

  @Expose()
  @IsString({ message: 'Link label không hợp lệ' })
  @MinLength(1, { message: 'Link label là bắt buộc' })
  @MaxLength(255, { message: 'Link label không được quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  link_label!: string;

  @Expose()
  @IsEnum(LinkType, { message: 'Link type phải là internal hoặc external' })
  link_type!: LinkType;
}
