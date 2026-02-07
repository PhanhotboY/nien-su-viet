import { Exclude, Expose, Type } from 'class-transformer';

@Exclude()
export class OperationResponse {
  @Expose()
  id!: string;

  @Expose()
  success!: boolean;
}

export class OperationResponseDto {
  @Expose()
  @Type(() => Object)
  data: OperationResponse;
}
