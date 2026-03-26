import { Exclude, Expose, Transform, Type } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsInt,
  IsUUID,
  Length,
  MinLength,
  IsUrl,
  Min,
  IsBoolean,
  IsDate,
} from 'class-validator';
import { Post } from '@phanhotboy/genproto/post_service/posts';

// DTO for Post
@Exclude()
export class PostBaseDto
  implements Omit<Post, 'publishedAt' | 'createdAt' | 'updatedAt'>
{
  @Expose()
  @IsUUID('4', { message: 'Id không hợp lệ' })
  id!: string;

  @Expose()
  @IsString({ message: 'Tiêu đề không hợp lệ' })
  @Length(1, 255, { message: 'Độ dài tiêu đề phải từ 1 đến 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  title!: string;

  @Expose()
  @IsString({ message: 'Slug không hợp lệ' })
  @Length(1, 255, { message: 'Slug phải từ 1 đến 255 ký tự' })
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
  @Length(1, 500, { message: 'Tóm tắt phải từ 1 đến 500 ký tự' })
  @Transform(({ value }) => value?.trim())
  summary?: string;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Thumbnail không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  thumbnail?: string;

  @Expose()
  @IsString({ message: 'ID tác giả không hợp lệ' })
  authorId!: string;

  @Expose()
  @IsOptional()
  @IsUUID('4', { message: 'ID danh mục không hợp lệ' })
  categoryId?: string;

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
  publishedAt!: string;

  @Expose()
  createdAt!: string;

  @Expose()
  updatedAt!: string;
}
