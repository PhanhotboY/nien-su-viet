'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { authClient } from '@/lib/auth-client';
import type { components, operations } from '@nsv-interfaces/auth-service';
import { ApiError } from 'next/dist/server/api-utils';
import { headers } from 'next/headers';

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
  const { data, error } = await authClient.$fetch<{
    users: components['schemas']['User'][];
    total: number;
  }>('/admin/list-users', {
    method: 'GET',
    query,
    headers: reqHeaders,
  });

  if (error) {
    console.log('Error fetching users: ', error);
    throw new ApiError(error.status, error.message || 'Failed to fetch users');
  }

  return {
    data: data.users,
    pagination: {
      total: data.total ?? data.users.length,
      limit: query.limit,
      page: query.page,
      totalPages: Math.ceil(data.total / query.limit),
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

  const { data, error } = await authClient.$fetch<
    components['schemas']['User']
  >('/admin/create-user', {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify(userData),
  });

  if (error) {
    console.log('Error creating user: ', error);
    throw new ApiError(error.status, error.message || 'Failed to create user');
  }

  return data;
}

async function deleteUser(userId: string) {
  const reqHeaders = new Headers(await headers());
  // Foward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const { error } = await authClient.$fetch<void>(`/admin/remove-user`, {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify({ userId }),
  });

  if (error) {
    console.log('Error deleting user: ', error);
    throw new ApiError(error.status, error.message || 'Failed to delete user');
  }

  return;
}

async function updateUserRole(
  userId: string,
  newRole: string,
): Promise<components['schemas']['User']> {
  const reqHeaders = new Headers(await headers());
  // Foward headers from client with different payload
  reqHeaders.delete('content-length');
  reqHeaders.set('content-type', 'application/json');

  const { data, error } = await authClient.$fetch<
    components['schemas']['User']
  >('/admin/update-user', {
    method: 'POST',
    headers: reqHeaders,
    body: JSON.stringify({ userId, data: { role: newRole } }),
  });

  if (error) {
    console.log('Error updating user role: ', error);
    throw new ApiError(
      error.status,
      error.message || 'Failed to update user role',
    );
  }

  return data;
}

export { getUsers, createUser, deleteUser, updateUserRole };
