import { Exclude, Expose } from 'class-transformer';
import { PaginationMetadataDto } from './pagination.dto';

@Exclude()
export class SerializedResponseDto<T> {
  @Expose()
  data: T;

  @Expose()
  message: string;

  @Exclude()
  statusCode: number;

  @Exclude()
  timestamp: string;
}
