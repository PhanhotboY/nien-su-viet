import { Exclude, Expose, Type } from 'class-transformer';

@Exclude()
export class OperationMetadataDto {
  @Expose()
  id!: string;

  @Expose()
  success!: boolean;
}
