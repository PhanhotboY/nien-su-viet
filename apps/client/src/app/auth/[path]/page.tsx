'use client';

import { authClient } from '@/lib/auth-client';
import { AuthView } from '@daveyplate/better-auth-ui';
import Link from 'next/link';
import { redirect, useSearchParams } from 'next/navigation';
import { useEffect, useState } from 'react';

export default function AuthPage({
  params,
}: {
  params: Promise<{ path: string }>;
}) {
  const searchParams = useSearchParams();
  const { data: session, isPending } = authClient.useSession();
  const [path, setPath] = useState<string>('');

  useEffect(() => {
    params.then((p) => setPath(p.path));
  }, [params]);

  useEffect(() => {
    if (!isPending && session) {
      let redirectTo = searchParams.get('redirectTo');
      if (!redirectTo) {
        switch (session.user.role) {
          case 'admin':
            redirectTo = '/admin';
            break;
          case 'editor':
            redirectTo = '/cmsdesk';
            break;
          default:
            redirectTo = '/';
        }
      }
      redirect(redirectTo);
    }
  }, [session, isPending, searchParams]);

  if (!path) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <main className="h-full container flex grow flex-col items-center justify-center gap-4 self-center p-4 md:p-6">
      <AuthView path={path} />

      {!['callback', 'sign-out'].includes(path) && (
        <p className="w-3xs text-center text-muted-foreground text-xs">
          By continuing, you agree to our{' '}
          <Link
            className="text-warning underline"
            href="/terms"
            target="_blank"
          >
            Terms of Service
          </Link>{' '}
          and{' '}
          <Link
            className="text-warning underline"
            href="/privacy"
            target="_blank"
          >
            Privacy Policy
          </Link>
          .
        </p>
      )}
    </main>
  );
}
