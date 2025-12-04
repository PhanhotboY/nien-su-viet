'use server';

import { IPaginatedResponse } from '@/interfaces/response.interface';
import { authClient } from '@/lib/auth-client';
import type { components, operations } from '@nsv-interfaces/auth-service';
import { ApiError } from 'next/dist/server/api-utils';
import { cookies, headers } from 'next/headers';

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

  // Get users from Better Auth
  const { data, error } = await authClient.$fetch<{
    users: components['schemas']['User'][];
    total: number;
  }>('/admin/list-users', {
    method: 'GET',
    query,
    headers: { Cookie: (await cookies()).toString() },
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
  userData: operations['createUser']['requestBody']['content']['application/json'],
) {
  const { data, error } = await authClient.$fetch<
    components['schemas']['User']
  >('/admin/create-user', {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
      Cookie: (await cookies()).toString(),
    },
    body: JSON.stringify(userData),
  });

  if (error) {
    console.log('Error creating user: ', error);
    throw new ApiError(error.status, error.message || 'Failed to create user');
  }

  return data;
}

async function deleteUser(userId: string) {
  const { error } = await authClient.$fetch<void>(`/admin/remove-user`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
      Cookie: (await cookies()).toString(),
    },
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
  const { data, error } = await authClient.$fetch<
    components['schemas']['User']
  >('/admin/update-user', {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
      Cookie: (await cookies()).toString(),
    },
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
