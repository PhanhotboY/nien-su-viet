import { Exclude, Expose, Type } from 'class-transformer';

import { EventEnvelope } from './eventEnvelop';
import { UserAddedToOrganization } from '@phanhotboy/genproto/events/user_event';
import { IsBoolean, IsString, IsUUID } from 'class-validator';

@Exclude()
export class UserAddedToOrganizationEvent extends EventEnvelope {
  @Expose()
  @Type(() => UserAddedToOrganizationData)
  data: UserAddedToOrganizationData | undefined = undefined;

  constructor(partial: Partial<UserAddedToOrganizationEvent>) {
    super(partial);
    Object.assign(this, { userRoleUpdated: partial });
  }
}

@Exclude()
export class UserAddedToOrganizationData implements UserAddedToOrganization {
  @Expose()
  @IsString({ message: 'Invalid user id' })
  userId: string;

  @Expose()
  @IsUUID('4', { message: 'Invalid organization id' })
  organizationId: string;

  @Expose()
  @IsBoolean({ message: 'Invalid isPremium value' })
  isPremium: boolean;

  constructor(partial: Partial<UserAddedToOrganizationData>) {
    Object.assign(this, partial);
  }
}
