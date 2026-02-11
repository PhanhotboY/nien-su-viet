import { Exclude, Expose, Type } from 'class-transformer';

@Exclude()
class PaginationMetaDto {
  @Expose()
  total!: number;

  @Expose()
  totalPages!: number;

  @Expose()
  page!: number;

  @Expose()
  limit!: number;
}

export class PaginatedResponseDto<T> {
  @Expose()
  data: T[];

  @Expose()
  @Type(() => PaginationMetaDto)
  pagination: PaginationMetaDto;
}
