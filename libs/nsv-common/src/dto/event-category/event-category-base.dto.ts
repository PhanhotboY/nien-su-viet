import { Exclude, Expose, Transform } from 'class-transformer';
import { IsString, IsUUID, MinLength } from 'class-validator';

// DTO for creating historical event
@Exclude()
export class EventCategoryBaseDto {
  @Expose()
  @IsUUID('4', { message: 'Id không hợp lệ' })
  id!: string;

  @Expose()
  @IsString({ message: 'Tên không hợp lệ' })
  @MinLength(1, { message: 'Tên là bắt buộc' })
  @Transform(({ value }) => value?.trim())
  name!: string;

  @Expose()
  @IsString()
  slug!: string;

  @Expose()
  createdAt!: string | Date;

  @Expose()
  updatedAt!: string | Date;
}
