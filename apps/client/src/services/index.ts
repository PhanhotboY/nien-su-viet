'use server';

import { cookies, headers } from 'next/headers';
import { parseCookies } from 'better-call';
import { AUTH_BASE_URL } from '@/lib/config';

export const retryFetcher = async <T = any>(
  url: string,
  options?: RequestInit,
): Promise<{
  data: T;
  message: string;
  statusCode: number;
  timestamp: string;
  pagination?: {
    total: number;
    limit: number;
    totalPages: number;
    page: number;
  };
}> => {
  const reqHeaders = new Headers(await headers());
  const store = await cookies();
  const sessionData = store.get('better-auth.session_data');
  const sessionToken = store.get('better-auth.session_token');

  // Check if user is authenticated but JWT is expired
  // If user is not authenticated yet, just continue fetching and throw error if any
  if (sessionToken && !sessionData) {
    // Refresh session data with sesion token
    const res = await fetch(AUTH_BASE_URL + '/auth/get-session', {
      headers: { Cookie: store.toString() },
    });

    const newCookies = parseCookies(res.headers.get('set-cookie') || '');
    store.set(
      'better-auth.session_data',
      newCookies.get('better-auth.session_data') || '',
    );
    reqHeaders.set('Cookie', store.toString());
  }

  const response = await fetch(url, {
    method: 'GET',
    ...options,
    headers: {
      ...Object.fromEntries(reqHeaders.entries()),
      ...(options?.body instanceof FormData
        ? {}
        : { 'Content-Type': 'application/json' }),
      ...options?.headers,
    },
  }).catch((error) => {
    console.log('fetch error');
    console.error(error);

    throw new Response('Lỗi hệ thống', {
      status: 500,
    });
  });
  if (!response) {
    console.log('fetch error');
    console.error('No response');

    throw new Response('Lỗi hệ thống', {
      status: 500,
    });
  }

  const data = (await response.json()) as {
    data: T;
    message: string;
    path: string;
    statusCode: number;
    timestamp: string;
  };

  if (response.ok) {
    console.log(
      '%s %s \x1b[32m%s\x1b[0m',
      options?.method || 'GET',
      url,
      response.status,
    );
  } else {
    console.log(
      '%s %s \x1b[31m%s\x1b[0m',
      options?.method || 'GET',
      url,
      response.status,
    );
  }
  return data;
};
