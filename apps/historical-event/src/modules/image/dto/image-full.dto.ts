import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsUUID,
  IsUrl,
  MinLength,
  MaxLength,
} from 'class-validator';
import { ImageBaseDto } from './image-base.dto';

// DTO for creating image
@Exclude()
export class ImageFullDto extends ImageBaseDto {
  @Expose()
  @IsString({ message: 'Tên phải là chuỗi' })
  @MinLength(1, { message: 'Tên là bắt buộc' })
  @MaxLength(255, { message: 'Tên không được quá 255 ký tự' })
  @Transform(({ value }) => value?.trim())
  name!: string;

  @Expose()
  @IsOptional()
  @IsUrl({}, { message: 'Link không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  redirectLink?: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Loại phải là chuỗi' })
  @MaxLength(100, { message: 'Loại không được quá 100 ký tự' })
  @Transform(({ value }) => value?.trim())
  type?: string;

  @Expose()
  @IsUUID('4', { message: 'ID người tải lên không hợp lệ' })
  uploaderId!: string;
}
