import Link from 'next/link';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import { AuthView } from '@daveyplate/better-auth-ui';

import { authClient } from '@/lib/auth-client';

export default async function AuthPage({
  params,
  searchParams,
}: {
  params: Promise<{ path: string }>;
  searchParams: Promise<{ redirectTo?: string }>;
}) {
  const { path } = await params;
  const { data, error } = await authClient.getSession({
    fetchOptions: { headers: { Cookie: (await cookies()).toString() } },
  });

  let { redirectTo } = await searchParams;
  if (!redirectTo) {
    switch (data?.user.role) {
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

  if (data && !error) {
    redirect(redirectTo);
  }

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
