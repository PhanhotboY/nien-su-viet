import { Exclude, Expose, Transform, Type } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsEmail,
  IsUrl,
  MaxLength,
  MinLength,
  ValidateNested,
} from 'class-validator';

// Social Links DTO
@Exclude()
export class SocialLinksDto {
  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Facebook URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  facebook?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Youtube URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  youtube?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Tiktok URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  tiktok?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Zalo URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  zalo?: string | null;
}

// Address DTO
@Exclude()
export class AddressDto {
  @Expose()
  @IsOptional()
  @IsString({ message: 'Tỉnh/Thành phố không hợp lệ' })
  @MaxLength(100, { message: 'Tỉnh/Thành phố không được quá 100 ký tự' })
  @Transform(({ value }) => value?.trim())
  province?: string | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Quận/Huyện không hợp lệ' })
  @MaxLength(100, { message: 'Quận/Huyện không được quá 100 ký tự' })
  @Transform(({ value }) => value?.trim())
  district?: string | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Địa chỉ không hợp lệ' })
  @MaxLength(255, { message: 'Địa chỉ không được quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  street?: string | null;
}

// App Data DTO
@Exclude()
export class AppDto {
  @Expose()
  @IsString({ message: 'Tiêu đề không hợp lệ' })
  @MinLength(1, { message: 'Tiêu đề là bắt buộc' })
  @MaxLength(255, { message: 'Tiêu đề không được quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  title!: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Mô tả không hợp lệ' })
  @MaxLength(1000, { message: 'Mô tả không được quá 1000 ký tự' })
  @Transform(({ value }) => value?.trim())
  description?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Logo URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  logo?: string | null;

  @Expose()
  @IsOptional()
  @ValidateNested()
  @Type(() => SocialLinksDto)
  social?: SocialLinksDto | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Mã số thuế không hợp lệ' })
  @MaxLength(50, { message: 'Mã số thuế không được quá 50 ký tự' })
  @Transform(({ value }) => value?.trim())
  tax_code?: string | null;

  @Expose()
  @IsOptional()
  @ValidateNested()
  @Type(() => AddressDto)
  address?: AddressDto | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Số điện thoại không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  msisdn?: string | null;

  @Expose()
  @IsOptional()
  @IsEmail({}, { message: 'Email không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  email?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Map URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  map?: string | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Head scripts không hợp lệ' })
  @MaxLength(5000, { message: 'Head scripts không được quá 5000 ký tự' })
  head_scripts?: string | null;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Body scripts không hợp lệ' })
  @MaxLength(5000, { message: 'Body scripts không được quá 5000 ký tự' })
  body_scripts?: string | null;
}
