import { SetMetadata } from '@nestjs/common';
import { statements } from './permissions';

export const PERMISSIONS_KEY = 'permissions';

export function Permissions(permissions: Partial<typeof statements>) {
  return SetMetadata(PERMISSIONS_KEY, permissions);
}
