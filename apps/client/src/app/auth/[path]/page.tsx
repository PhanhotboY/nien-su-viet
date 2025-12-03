import { authClient } from '@/lib/auth-client';
import { AuthView } from '@daveyplate/better-auth-ui';
import { authViewPaths } from '@daveyplate/better-auth-ui/server';
import { headers } from 'next/headers';
import Link from 'next/link';
import { redirect } from 'next/navigation';

export const dynamicParams = false;

export function generateStaticParams() {
  return Object.values(authViewPaths).map((path) => ({ path }));
}

export default async function AuthPage({
  params,
  searchParams,
}: {
  params: Promise<{ path: string }>;
  searchParams: Promise<{ redirectTo?: string }>;
}) {
  const { path } = await params;
  const reqHeaders = await headers();
  const { data, error } = await authClient.getSession({
    fetchOptions: { headers: { Cookie: reqHeaders.get('cookie') || '' } },
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
