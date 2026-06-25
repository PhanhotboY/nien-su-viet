import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsEmail,
  IsOptional,
  MinLength,
  IsUUID,
} from 'class-validator';

// Base DTO for user creation
@Exclude()
export class UserBaseDto {
  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  id!: string;

  @Expose()
  @IsEmail({}, { message: 'Email không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  email!: string;

  @Expose()
  @IsString({ message: 'Tên phải là chuỗi' })
  @MinLength(1, { message: 'Tên là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  name!: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Avatar không hợp lệ' })
  image?: string | null;

  @Expose()
  @IsString({
    message: 'Vai trò phải là chuỗi',
  })
  role!: string;
}
