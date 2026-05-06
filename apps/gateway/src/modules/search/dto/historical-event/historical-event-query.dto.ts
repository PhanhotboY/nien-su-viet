import { Transform, Type } from 'class-transformer';
import {
  IsArray,
  IsDateString,
  IsIn,
  IsInt,
  IsNumber,
  IsOptional,
  IsPositive,
  IsString,
  IsUUID,
  Max,
  Min,
} from 'class-validator';

// DTO for historical event query parameters
export class HistoricalEventQueryDto {
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

  // Filter by author
  @IsOptional()
  @IsUUID('4', { message: 'ID tác giả không hợp lệ' })
  authorId?: string;

  // Filter by category
  @IsOptional()
  @Transform(({ value }) => {
    if (typeof value === 'string') {
      return value.split(',').map((v) => v.trim());
    }
    return Array.isArray(value) ? value : undefined;
  })
  @IsArray({ message: 'Danh sách danh mục phải là mảng' })
  @IsUUID('4', { message: 'ID danh mục không hợp lệ' })
  categoryIds?: string[];

  // Filter by date range (from)
  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Năm bắt đầu phải là số nguyên' })
  fromYear?: number;

  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Tháng bắt đầu phải là số nguyên' })
  @Min(1, { message: 'Tháng phải từ 1 đến 12' })
  @Max(12, { message: 'Tháng phải từ 1 đến 12' })
  @Transform(({ value }) => (value ? parseInt(value, 10) : null))
  fromMonth?: number;

  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Ngày bắt đầu phải là số nguyên' })
  @Min(1, { message: 'Ngày phải từ 1 đến 31' })
  @Max(31, { message: 'Ngày phải từ 1 đến 31' })
  @Transform(({ value }) => (value ? parseInt(value, 10) : null))
  fromDay?: number;

  // Filter by date range (to)
  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Năm kết thúc phải là số nguyên' })
  @Transform(({ value }) => (value ? parseInt(value, 10) : null))
  toYear?: number;

  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Tháng kết thúc phải là số nguyên' })
  @Min(1, { message: 'Tháng phải từ 1 đến 12' })
  @Max(12, { message: 'Tháng phải từ 1 đến 12' })
  @Transform(({ value }) => (value ? parseInt(value, 10) : null))
  toMonth?: number;

  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Ngày kết thúc phải là số nguyên' })
  @Min(1, { message: 'Ngày phải từ 1 đến 31' })
  @Max(31, { message: 'Ngày phải từ 1 đến 31' })
  toDay?: number;

  // Filter by year range (searching events within a specific year)
  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Năm tìm kiếm phải là số nguyên' })
  searchYear?: number;

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
