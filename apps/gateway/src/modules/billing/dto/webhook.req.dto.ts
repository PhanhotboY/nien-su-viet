import { IsNumber, IsString } from 'class-validator';

export class WebhookRequestDto {
  @IsString({ message: 'invalid data' })
  data: string;

  @IsString({ message: 'invalid mac' })
  mac: string;

  /** 1: Order, 2: Agreement */
  @IsNumber(
    { allowNaN: false, allowInfinity: false },
    { message: 'invalid type' },
  )
  type: number;
}
