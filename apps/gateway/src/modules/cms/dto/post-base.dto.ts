import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsInt,
  IsUUID,
  Min,
  MinLength,
  IsBoolean,
  IsUrl,
  MaxLength,
} from 'class-validator';

// DTO for creating post
@Exclude()
export class PostBaseDto {
  @Expose()
  @IsUUID('4', { message: 'Id không hợp lệ' })
  id!: string;

  @Expose()
  @IsString({ message: 'Tiêu đề không hợp lệ' })
  @MinLength(1, { message: 'Tiêu đề là bắt buộc' })
  @MaxLength(255, { message: 'Tiêu đề không được quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  title!: string;

  @Expose()
  @IsString({ message: 'Slug không hợp lệ' })
  @MinLength(1, { message: 'Slug là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  slug!: string;

  @Expose()
  @IsOptional()
  @IsUUID('4', { message: 'ID danh mục không hợp lệ' })
  categoryId?: string;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Thumbnail không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  thumbnail?: string | null;

  @Expose()
  @IsString({ message: 'ID tác giả không hợp lệ' })
  authorId!: string;

  @Expose()
  @IsString({ message: 'Nội dung không hợp lệ' })
  @MinLength(1, { message: 'Nội dung là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  content!: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Tóm tắt không hợp lệ' })
  @MaxLength(5000, { message: 'Tóm tắt không được quá 5000 ký tự' })
  @Transform(({ value }) => value?.trim())
  summary?: string | null;

  @Expose()
  @IsInt({ message: 'Lượt xem không hợp lệ' })
  @Min(0, { message: 'Lượt xem không hợp lệ' })
  @Transform(({ value }) => (value !== undefined ? parseInt(value, 10) : value))
  views!: number;

  @Expose()
  @IsInt({ message: 'Lượt thích không hợp lệ' })
  @Min(0, { message: 'Lượt thích không hợp lệ' })
  @Transform(({ value }) => (value !== undefined ? parseInt(value, 10) : value))
  likes!: number;

  @Expose()
  @IsBoolean({ message: 'Trạng thái đăng không hợp lệ' })
  published!: boolean;

  @Expose()
  createdAt!: string | Date;

  @Expose()
  updatedAt!: string | Date;
}
