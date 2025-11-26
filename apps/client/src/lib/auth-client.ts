import { createAuthClient } from 'better-auth/react';
import { DEFAULT_LOGIN_REDIRECT } from './config';
import { adminClient } from 'better-auth/client/plugins';

const baseURL =
  process.env.NEXT_PUBLIC_BETTER_AUTH_URL ||
  'http://localhost:4000/api/v1/auth';

export const authClient = createAuthClient({
  /** The base URL of the server (optional if you're using the same domain) */
  baseURL,
  plugins: [adminClient()],
});

export const signInWithGoogle = async () => {
  await authClient.signIn.social({
    provider: 'google',
    callbackURL: DEFAULT_LOGIN_REDIRECT,
  });
};
