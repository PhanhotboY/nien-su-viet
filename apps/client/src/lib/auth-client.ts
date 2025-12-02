import { createAuthClient } from 'better-auth/react';
import { AUTH_BASE_URL, DEFAULT_LOGIN_REDIRECT } from './config';
import { adminClient } from 'better-auth/client/plugins';
import { ac, roles } from './permissions';

export const authClient = createAuthClient({
  /** The base URL of the server (optional if you're using the same domain) */
  baseURL: AUTH_BASE_URL + '/auth',
  plugins: [adminClient({ ac, roles })],
});

export const signInWithGoogle = async () => {
  await authClient.signIn.social({
    provider: 'google',
    callbackURL: DEFAULT_LOGIN_REDIRECT,
  });
};
