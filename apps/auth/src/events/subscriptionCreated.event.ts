import { Exclude, Expose, Type } from 'class-transformer';
import { IsString, IsUUID } from 'class-validator';

import { SubscriptionCreated } from '@phanhotboy/genproto/events/billing_event';
import { SubscriptionStatus } from '@phanhotboy/genproto/billing_service/subscriptions';
import { ToDate } from '@phanhotboy/nsv-common/util/grpc.util';
import { EventEnvelope } from './eventEnvelop';

@Exclude()
export class SubscriptionCreatedEvent extends EventEnvelope {
  @Expose()
  @Type(() => SubscriptionCreatedEventData)
  declare data: SubscriptionCreatedEventData | undefined;

  constructor(partial: Partial<SubscriptionCreatedEvent>) {
    super(partial);
    Object.assign(this, { subscription: partial });
  }
}

export class SubscriptionCreatedEventData
  implements
    Omit<
      SubscriptionCreated,
      | 'currentPeriodStart'
      | 'currentPeriodEnd'
      | 'canceledAt'
      | 'expiredAt'
      | 'createdAt'
      | 'updatedAt'
    >
{
  @Expose()
  @IsUUID('4', { message: 'id must be a valid UUID v4' })
  subscriptionId: string;

  @Expose()
  @IsString({ message: 'userId must be a string' })
  userId: string;

  @Expose()
  @IsUUID('4', { message: 'planId must be a valid UUID v4' })
  planId: string;

  @Expose()
  status: SubscriptionStatus;

  @Expose()
  @ToDate()
  currentPeriodStart?: Date;

  @Expose()
  @ToDate()
  currentPeriodEnd?: Date;

  @Expose()
  cancelAtPeriodEnd: boolean;

  @Expose()
  @ToDate()
  canceledAt?: Date;

  @Expose()
  @ToDate()
  expiredAt?: Date;

  @Expose()
  @ToDate()
  createdAt?: Date;

  @Expose()
  @ToDate()
  updatedAt?: Date;

  constructor(partial: Partial<SubscriptionCreatedEventData>) {
    Object.assign(this, { subscription: partial });
  }
}
