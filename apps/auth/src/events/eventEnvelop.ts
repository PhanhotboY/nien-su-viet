import { EventEnvelope as EE } from '@phanhotboy/genproto/events/events';
import { ToDate } from '@phanhotboy/nsv-common/util/grpc.util';
import { Exclude, Expose, Type } from 'class-transformer';
import { IsString, IsUUID } from 'class-validator';

@Exclude()
export class EventEnvelope implements Omit<EE, 'occurredAt' | 'data'> {
  @Expose()
  @IsUUID('4', { message: 'Invalid event id' })
  eventId: string;
  @Expose()
  @IsString({ message: 'Invalid aggregate type' })
  aggregateType: string;
  @Expose()
  @IsUUID('4', { message: 'Invalid aggregate id' })
  aggregateId: string;

  @Expose()
  @IsString({ message: 'Invalid event pattern' })
  pattern: string;

  @Expose()
  @ToDate()
  occurredAt?: Date;
  @Expose()
  attributes: { [key: string]: any } | undefined;

  constructor(partial: Partial<EventEnvelope>) {
    Object.assign(this, partial);
  }
}
