import { Exclude, Expose } from 'class-transformer';
import { ArrayMinSize, IsArray, IsString, IsUUID } from 'class-validator';

@Exclude()
export class UserDeleteDto {
  @Expose()
  @IsString({ message: 'ID không hợp lệ' })
  userId!: string;
}

// DTO for bulk deleting users
@Exclude()
export class UserBulkDeleteDto {
  @Expose()
  @IsArray({ message: 'Danh sách ID phải là mảng' })
  @ArrayMinSize(1, { message: 'Cần ít nhất một ID người dùng' })
  @IsString({ each: true, message: 'ID người dùng không hợp lệ' })
  userIds!: string[];
}
