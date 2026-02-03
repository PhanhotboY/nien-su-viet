import { Exclude, Expose, Transform } from 'class-transformer';
import {
  IsString,
  IsOptional,
  IsUUID,
  IsUrl,
  MaxLength,
} from 'class-validator';

// DTO for creating image
@Exclude()
export class ImageBaseDto {
  @Expose()
  @IsUUID('4', { message: 'ID không hợp lệ' })
  id?: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Chú thích phải là chuỗi' })
  @Transform(({ value }) => value?.trim())
  caption?: string;

  @Expose()
  @IsOptional()
  @IsString({ message: 'Mô tả phải là chuỗi' })
  @Transform(({ value }) => value?.trim())
  description?: string;

  @Expose()
  @IsUrl({}, { message: 'URL không hợp lệ' })
  @Transform(({ value }) => value?.trim())
  publicUrl!: string;
}
