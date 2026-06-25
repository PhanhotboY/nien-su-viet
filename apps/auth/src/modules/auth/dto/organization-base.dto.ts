import { Exclude, Expose, Transform, Type } from 'class-transformer';
import {
  IsString,
  IsOptional,
  MinLength,
  IsNumber,
  IsDate,
} from 'class-validator';
import { MemberBaseDto } from './member-base.dto';

// Base DTO for user creation
@Exclude()
export class OrganizationBaseDto {
  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  id!: string;

  @Expose()
  @IsString({ message: 'Tên phải là chuỗi' })
  @MinLength(1, { message: 'Tên là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  name!: string;

  @Expose()
  @IsString({ message: 'Tên phải là chuỗi' })
  @MinLength(1, { message: 'Tên là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  slug: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Logo không hợp lệ' })
  logo?: string;

  @Expose()
  @IsOptional()
  @IsDate({ message: 'Date không hợp lệ' })
  createdAt: Date;

  @Expose()
  @Type(() => MemberBaseDto)
  members: MemberBaseDto[];

  @Expose()
  @IsOptional()
  @IsString({ message: 'Metadata không hợp lệ' })
  metadata?: string;
}
