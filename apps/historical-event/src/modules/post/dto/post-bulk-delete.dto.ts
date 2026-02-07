import { ArrayMinSize, IsArray, IsUUID } from 'class-validator';

// DTO for bulk deleting historical events
export class PostBulkDeleteDto {
  @IsArray({ message: 'Danh sách ID phải là mảng' })
  @ArrayMinSize(1, { message: 'Cần ít nhất một ID post' })
  @IsUUID('4', { each: true, message: 'ID post không hợp lệ' })
  postIds!: string[];
}
