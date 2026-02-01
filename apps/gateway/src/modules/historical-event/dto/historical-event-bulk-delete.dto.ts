import { ArrayMinSize, IsArray, IsUUID } from 'class-validator';

// DTO for bulk deleting historical events
export class HistoricalEventBulkDeleteDto {
  @IsArray({ message: 'Danh sách ID phải là mảng' })
  @ArrayMinSize(1, { message: 'Cần ít nhất một ID sự kiện lịch sử' })
  @IsUUID('4', { each: true, message: 'ID sự kiện lịch sử không hợp lệ' })
  eventIds!: string[];
}
