import { Transform, Type } from 'class-transformer';
import {
  IsDateString,
  IsIn,
  IsNumber,
  IsOptional,
  IsPositive,
  IsString,
  IsUUID,
} from 'class-validator';

// DTO for image query parameters
export class QueryImageDto {
  @IsOptional()
  @Type(() => Number)
  @IsNumber({}, { message: 'Trang phải là số' })
  @IsPositive({ message: 'Trang phải là số dương' })
  page?: number;

  @IsOptional()
  @Type(() => Number)
  @IsNumber({}, { message: 'Giới hạn phải là số' })
  @IsPositive({ message: 'Giới hạn phải là số dương' })
  limit?: number;

  @IsOptional()
  @IsString({ message: 'Từ khóa tìm kiếm phải là chuỗi' })
  search?: string;

  @IsOptional()
  @IsString({ message: 'Trường sắp xếp phải là chuỗi' })
  sortBy?: string;

  @IsOptional()
  @IsIn(['asc', 'desc'], { message: 'Thứ tự sắp xếp phải là asc hoặc desc' })
  sortOrder?: 'asc' | 'desc';

  // Filter by name
  @IsOptional()
  @IsString({ message: 'Tên phải là chuỗi' })
  name?: string;

  // Filter by type (thumbnail, banner, avatar, etc.)
  @IsOptional()
  @IsString({ message: 'Loại phải là chuỗi' })
  type?: string;

  // Filter by uploader
  @IsOptional()
  @IsUUID('4', { message: 'ID người tải lên không hợp lệ' })
  uploaderId?: string;

  // Filter by created date range
  @IsOptional()
  @IsDateString({}, { message: 'Ngày tạo từ không hợp lệ' })
  @Transform(({ value }: any) => (value ? new Date(value) : undefined))
  createdAtFrom?: Date;

  @IsOptional()
  @IsDateString({}, { message: 'Ngày tạo đến không hợp lệ' })
  @Transform(({ value }: any) => (value ? new Date(value) : undefined))
  createdAtTo?: Date;

  // Filter by updated date range
  @IsOptional()
  @IsDateString({}, { message: 'Ngày cập nhật từ không hợp lệ' })
  @Transform(({ value }: any) => (value ? new Date(value) : undefined))
  updatedAtFrom?: Date;

  @IsOptional()
  @IsDateString({}, { message: 'Ngày cập nhật đến không hợp lệ' })
  @Transform(({ value }: any) => (value ? new Date(value) : undefined))
  updatedAtTo?: Date;
}
