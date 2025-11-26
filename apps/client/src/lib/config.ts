const CLIENT_HOST =
  process.env.NEXT_PUBLIC_CLIENT_HOST || 'http://localhost:3000';
const DEFAULT_LOGIN_REDIRECT = CLIENT_HOST + '/admin';

export { DEFAULT_LOGIN_REDIRECT };
