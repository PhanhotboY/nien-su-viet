'use server';

import { cookies, headers } from 'next/headers';
import { parse } from 'set-cookie-parser';
import { AUTH_BASE_URL, AUTH_COOKIE_PREFIX } from '@/lib/config';

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
  reqHeaders.delete('content-length'); // UND_ERR_REQ_CONTENT_LENGTH_MISMATCH
  reqHeaders.delete('content-type');
  const store = await cookies();
  const sessionDataCookieKey = `${AUTH_COOKIE_PREFIX}.session_data`;
  const sessionData = store.get(sessionDataCookieKey);
  const sessionToken = store.get(`${AUTH_COOKIE_PREFIX}.session_token`);
  const isGetMethod = (options?.method || 'GET').toUpperCase() === 'GET';

  // Check if user is authenticated but JWT is expired
  // If user is not authenticated yet, just continue fetching and throw error if any
  if (sessionToken && !sessionData) {
    // Refresh session data with sesion token
    const res = await fetch(AUTH_BASE_URL + '/auth/get-session', {
      headers: { Cookie: store.toString() },
    });

    const newCookies = parse(res.headers.get('set-cookie') || '', {
      map: true,
    });
    if (newCookies && newCookies[sessionDataCookieKey]) {
      try {
        const { name, value, ...options } = newCookies[sessionDataCookieKey];
        store.set(sessionDataCookieKey, value, options as any);
        reqHeaders.set('Cookie', store.toString());
      } catch (error) {
        console.error('Error setting refreshed session cookie:', error);
      }
    }
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
    throw new Error(data.message || 'Lỗi hệ thống');
  }
  return data;
};
