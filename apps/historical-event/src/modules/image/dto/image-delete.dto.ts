import { ArrayMinSize, IsArray, IsUUID } from 'class-validator';

// DTO for bulk deleting images
export class BulkDeleteImageDto {
  @IsArray({ message: 'Danh sách ID phải là mảng' })
  @ArrayMinSize(1, { message: 'Cần ít nhất một ID hình ảnh' })
  @IsUUID('4', { each: true, message: 'ID hình ảnh không hợp lệ' })
  imageIds!: string[];
}
