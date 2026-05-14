'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { headers } from 'next/headers';
import { retryFetcher } from '.';
import type { components } from '@nsv-interfaces/nsv-api-documentation';
import { revalidateTag } from 'next/cache';

export type Organization = components['schemas']['OrganizationBaseDto'];

async function getOrganizations(): Promise<Organization[]> {
  const reqHeaders = new Headers(await headers());
  // Get organizations from Better Auth
  const res = (await retryFetcher<Organization[]>('/auth/admin/organizations', {
    method: 'GET',
    headers: reqHeaders,
    next: { tags: ['organizations'] },
  })) as any;

  return res.data;
}

async function getOrganizationById(
  organizationId: string,
): Promise<Organization> {
  const reqHeaders = new Headers(await headers());
  // Get organization details from Better Auth
  const res = (await retryFetcher<Organization>(
    `/auth/admin/organizations/${organizationId}`,
    {
      method: 'GET',
      headers: reqHeaders,
      next: { tags: ['organizations'] },
    },
  )) as any;

  return res.data;
}

async function createOrganization(data: {
  name: string;
  slug: string;
  userId?: string | null;
  logo?: string | null;
  metadata?: string | null;
  keepCurrentActiveOrganization?: boolean | null;
}): Promise<Organization> {
  const reqHeaders = new Headers(await headers());
  // Forward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<Organization>('/auth/organization/create', {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify(data),
  });

  revalidateTag('organizations', 'max');
  return res.data;
}

async function deleteOrganization(organizationId: string): Promise<string> {
  const reqHeaders = new Headers(await headers());
  // Forward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<string>('/auth/organization/delete', {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify({ organizationId }),
  });

  revalidateTag('organizations', 'max');
  return res.data;
}

async function updateOrganization(
  organizationId: string,
  data: {
    name?: string;
    slug?: string;
    logo?: string | null;
    metadata?: string | null;
  },
): Promise<Organization> {
  const reqHeaders = new Headers(await headers());
  // Forward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<Organization>('/auth/organization/update', {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify({ organizationId, data }),
  });

  revalidateTag('organizations', 'max');
  return res.data;
}

export {
  getOrganizations,
  getOrganizationById,
  createOrganization,
  deleteOrganization,
  updateOrganization,
};
