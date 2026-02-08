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

// export class PaginatedResponseDto<T> {
//   @Expose()
//   data: T[];

//   @Expose()
//   @Type(() => PaginationMetaDto)
//   pagination: PaginationMetaDto;
// }

export function PaginatedResponseDto<T extends Object>(ItemDto: T) {
  class PaginatedResponseDtoClass {
    @Expose()
    @Type(() => ItemDto as any)
    data: T[];

    @Expose()
    @Type(() => PaginationMetaDto)
    pagination: PaginationMetaDto;
  }

  return PaginatedResponseDtoClass;
}
