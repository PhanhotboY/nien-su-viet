import { OmitType } from '@nestjs/swagger';
import { TimestampDto } from '@phanhotboy/nsv-common';
import { TimestampUtil } from '@phanhotboy/nsv-common/util/grpc.util';
import { Transform, Type } from 'class-transformer';
import {
  IsArray,
  IsBooleanString,
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
export class PostQueryDto {
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
    return Array.isArray(value) ? value : [];
  })
  @IsArray({ message: 'Danh sách danh mục phải là mảng' })
  @IsUUID('4', { message: 'ID danh mục không hợp lệ' })
  categoryIds: string[];

  @IsOptional()
  @IsBooleanString({ message: 'Trạng thái xuất bản phải là true hoặc false' })
  published?: string;

  // Filter by created date range
  @IsOptional()
  @IsDateString({}, { message: 'Ngày tạo từ không hợp lệ' })
  createdAtFrom?: string;

  @IsOptional()
  @IsDateString({}, { message: 'Ngày tạo đến không hợp lệ' })
  createdAtTo?: string;

  // Filter by updated date range
  @IsOptional()
  @IsDateString({}, { message: 'Ngày cập nhật từ không hợp lệ' })
  updatedAtFrom?: string;

  @IsOptional()
  @IsDateString({}, { message: 'Ngày cập nhật đến không hợp lệ' })
  updatedAtTo?: string;
}

export class PostQueryGrpcDto extends OmitType(PostQueryDto, [
  'createdAtFrom',
  'createdAtTo',
  'updatedAtFrom',
  'updatedAtTo',
  'published',
]) {
  @IsOptional()
  @Transform(({ value }) => Boolean(value))
  published?: boolean;

  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  createdAtFrom?: TimestampDto;

  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  createdAtTo?: TimestampDto;

  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  updatedAtFrom?: TimestampDto;

  @IsOptional()
  @Transform(({ value }) => TimestampUtil.toTimestamp(value))
  updatedAtTo?: TimestampDto;
}
