'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import type { components, operations } from '@nsv-interfaces/auth-service';
import { headers } from 'next/headers';
import { retryFetcher } from '.';

async function getUsers(
  options: Record<string, string>,
): Promise<IPaginatedResponse<components['schemas']['User']>> {
  // Build query for Better Auth
  const query: Record<string, any> = {
    limit: parseInt(options.limit || '10'),
    offset: parseInt(options.offset || '0'),
  };

  // Sorting
  if (options.sortBy) query.sortBy = options.sortBy;
  if (['asc', 'desc'].includes(options.sortDirection || ''))
    query.sortDirection = options.sortDirection;

  // Filtering by role
  if (options.role) {
    query.filterField = 'role';
    query.filterOperator = 'eq';
    query.filterValue = options.role;
  }

  // Filtering by status (active/banned)
  if (options.status) {
    query.filterField = 'banned';
    query.filterOperator = 'eq';
    query.filterValue = options.status === 'banned' ? true : false;
  }

  // Filtering by email
  if (options.email) {
    query.searchField = 'email';
    query.searchOperator = 'contains';
    query.searchValue = options.email;
  }

  // Filtering by name
  if (options.name) {
    query.searchField = 'name';
    query.searchOperator = 'contains';
    query.searchValue = options.name;
  }

  const reqHeaders = await headers();
  // Get users from Better Auth
  const res = (await retryFetcher<components['schemas']['User'][]>(
    '/auth/admin/list-users?' + new URLSearchParams(query).toString(),
    {
      method: 'GET',
      headers: reqHeaders,
    },
  )) as any;

  return {
    data: res.users as any,
    pagination: {
      total: res.total ?? res.users.length,
      limit: query.limit,
      page: query.page,
      totalPages: Math.ceil(res.total / query.limit),
    },
    statusCode: 200,
    timestamp: new Date().toISOString(),
  };
}

async function createUser(
  userData: operations['post__admin_create_user']['requestBody']['content']['application/json'],
) {
  const reqHeaders = new Headers(await headers());
  // Foward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<components['schemas']['User']>(
    '/auth/admin/create-user',
    {
      method: 'POST',
      headers: reqHeaders,
      body: JSON.stringify(userData),
    },
  );

  return res;
}

async function deleteUser(userId: string) {
  const reqHeaders = new Headers(await headers());
  // Foward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher(`/auth/admin/remove-user`, {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify({ userId }),
  });

  return;
}

async function updateUserRole(userId: string, newRole: string) {
  const reqHeaders = new Headers(await headers());
  // Foward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const res = await retryFetcher<components['schemas']['User']>(
    '/auth/admin/update-user',
    {
      method: 'POST',
      headers: reqHeaders,
      body: JSON.stringify({ userId, data: { role: newRole } }),
    },
  );

  return;
}

export { getUsers, createUser, deleteUser, updateUserRole };
