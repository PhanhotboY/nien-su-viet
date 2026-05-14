'use server';

import { headers } from 'next/headers';
import { retryFetcher } from '.';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { revalidateTag } from 'next/cache';
import { IPaginatedResponse } from '@/interfaces/response.interface';

export type OperationResponse = components['schemas']['OperationMetadataDto'];
export type OrganizationMember = components['schemas']['MemberBaseDto'];

async function getOrganizationMembers(organizationId: string) {
  const reqHeaders = new Headers(await headers());

  const res = await retryFetcher<any>(
    `/auth/admin/organizations/${organizationId}/members`,
    {
      method: 'GET',
      headers: reqHeaders,
      next: { tags: ['organizations', 'organization-member'] },
    },
  );

  return res as IPaginatedResponse<OrganizationMember>;
}

async function addUserToOrganization(data: {
  userId: string;
  organizationId: string;
}): Promise<OperationResponse> {
  const reqHeaders = new Headers(await headers());
  // Forward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<OperationResponse>(
    `/auth/admin/organizations/${data.organizationId}/members/${data.userId}`,
    {
      method: 'POST',
      headers: reqHeaders,
    },
  );

  revalidateTag('organization-member', 'max');
  return res.data;
}

async function removeUserFromOrganization(data: {
  memberId: string;
  organizationId: string;
}): Promise<OperationResponse> {
  const reqHeaders = new Headers(await headers());
  // Forward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<OperationResponse>(
    `/auth/admin/organizations/${data.organizationId}/members/${data.memberId}`,
    {
      method: 'DELETE',
      headers: reqHeaders,
    },
  );

  revalidateTag('organization-member', 'max');
  return res.data;
}

export {
  getOrganizationMembers,
  addUserToOrganization,
  removeUserFromOrganization,
};
