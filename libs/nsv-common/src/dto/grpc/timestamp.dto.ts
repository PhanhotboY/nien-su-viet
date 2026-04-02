import { Timestamp } from '@phanhotboy/genproto/google/protobuf/timestamp';
import { IsInt, Min, Max } from 'class-validator';

export class TimestampDto implements Timestamp {
  @IsInt({ message: 'Ngày không hợp lệ' })
  @Min(0, { message: 'Ngày không hợp lệ' })
  seconds: number;

  @IsInt({ message: 'Ngày không hợp lệ' })
  @Min(0, { message: 'Ngày không hợp lệ' })
  @Max(999999999, { message: 'Ngày không hợp lệ' })
  nanos: number;
}
