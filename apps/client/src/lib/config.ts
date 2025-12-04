const CLIENT_HOST =
  process.env.NEXT_PUBLIC_CLIENT_HOST || 'http://localhost:3000';
const DEFAULT_LOGIN_REDIRECT = CLIENT_HOST + '/admin';

const AUTH_BASE_URL =
  process.env.NEXT_PUBLIC_BETTER_AUTH_URL || 'http://localhost:4000/api/v1';

const AUTH_COOKIE_PREFIX = process.env.AUTH_COOKIE_PREFIX || 'nsv-auth';

export {
  DEFAULT_LOGIN_REDIRECT,
  AUTH_BASE_URL,
  CLIENT_HOST,
  AUTH_COOKIE_PREFIX,
};
