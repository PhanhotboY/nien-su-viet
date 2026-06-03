'use client';

import { useAuth, useSession } from '@better-auth-ui/react';
import type { ReactNode } from 'react';

type AuthStateProps = {
  children: ReactNode;
};

export function SignedIn({ children }: AuthStateProps) {
  const { authClient } = useAuth();
  const { data: session } = useSession(authClient);

  if (!session) return null;

  return <>{children}</>;
}

export function SignedOut({ children }: AuthStateProps) {
  const { authClient } = useAuth();
  const { data: session, isPending } = useSession(authClient);

  if (session || isPending) return null;

  return <>{children}</>;
}
