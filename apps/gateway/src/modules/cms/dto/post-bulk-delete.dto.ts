import { ArrayMinSize, IsArray, IsUUID } from 'class-validator';

// DTO for bulk deleting posts
export class PostBulkDeleteDto {
  @IsArray({ message: 'Danh sách ID phải là mảng' })
  @ArrayMinSize(1, { message: 'Cần ít nhất một ID bài viết' })
  @IsUUID('4', { each: true, message: 'ID bài viết không hợp lệ' })
  postIds!: string[];
}
