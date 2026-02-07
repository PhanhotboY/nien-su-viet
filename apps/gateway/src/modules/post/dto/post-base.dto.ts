import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsInt,
  IsUUID,
  MinLength,
  MaxLength,
  IsUrl,
  Min,
  IsDate,
  IsBoolean,
} from 'class-validator';

// DTO for Post
@Exclude()
export class PostBaseDto {
  @Expose()
  @IsUUID('4', { message: 'Id không hợp lệ' })
  id!: string;

  @Expose()
  @IsString({ message: 'Tiêu đề không hợp lệ' })
  @MinLength(1, { message: 'Tiêu đề là bắt buộc' })
  @MaxLength(255, { message: 'Tiêu đề không được vượt quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  title!: string;

  @Expose()
  @IsString({ message: 'Slug không hợp lệ' })
  @MinLength(1, { message: 'Slug là bắt buộc' })
  @MaxLength(255, { message: 'Slug không được vượt quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  slug!: string;

  @Expose()
  @IsString({ message: 'Nội dung không hợp lệ' })
  @MinLength(1, { message: 'Nội dung là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  content!: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Tóm tắt không hợp lệ' })
  @MaxLength(500, { message: 'Tóm tắt không được vượt quá 500 ký tự' })
  @Transform(({ value }) => value?.trim())
  summary?: string | null;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Thumbnail không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  thumbnail?: string | null;

  @Expose()
  @IsUUID('4', { message: 'ID tác giả không hợp lệ' })
  authorId!: string;

  @Expose()
  @IsOptional()
  @IsUUID('4', { message: 'ID danh mục không hợp lệ' })
  categoryId?: string | null;

  @Expose()
  @IsInt({ message: 'Số lượt xem không hợp lệ' })
  @Min(0, { message: 'Số lượt xem không được âm' })
  views!: number;

  @Expose()
  @IsInt({ message: 'Số lượt thích không hợp lệ' })
  @Min(0, { message: 'Số lượt thích không được âm' })
  likes!: number;

  @Expose()
  @IsOptional()
  @IsBoolean({ message: 'Trạng thái xuất bản không hợp lệ' })
  published!: boolean;

  @Expose()
  @IsOptional()
  @IsDate({ message: 'Ngày xuất bản không hợp lệ' })
  publishedAt!: string | Date;

  @Expose()
  createdAt!: string | Date;

  @Expose()
  updatedAt!: string | Date;
}
