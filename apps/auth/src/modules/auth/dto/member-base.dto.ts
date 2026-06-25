import { Exclude, Expose, Transform, Type } from 'class-transformer';
import { IsString, IsOptional, MinLength, IsDate } from 'class-validator';

// Base DTO for user creation
@Exclude()
export class MemberBaseDto {
  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  id!: string;

  @Expose()
  @IsOptional()
  @IsDate({ message: 'Date không hợp lệ' })
  createdAt!: Date;

  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  organizationId!: string;

  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  userId!: string;

  @Expose()
  @IsString({
    message: 'Vai trò phải là chuỗi',
  })
  role!: string;
}
