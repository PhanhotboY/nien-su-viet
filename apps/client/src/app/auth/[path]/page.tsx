import Link from 'next/link';
import { redirect } from 'next/navigation';
import { AuthView } from '@daveyplate/better-auth-ui';

import { getAuthSession } from '@/helper/auth.helper';
import { authLocalization } from '@/localization/vi/auth-localization';

export default async function AuthPage({
  params,
  searchParams,
}: {
  params: Promise<{ path: string }>;
  searchParams: Promise<{ redirectTo?: string }>;
}) {
  const { path } = await params;
  const { user, error } = await getAuthSession();

  let { redirectTo } = await searchParams;
  if (!redirectTo) {
    switch (user?.role) {
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

  if (user && !error) {
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
          {authLocalization.BY_CONTINUING_YOU_AGREE}{' '}
          <Link
            className="text-warning underline"
            href="/terms"
            target="_blank"
          >
            {authLocalization.TERMS_OF_SERVICE}
          </Link>{' '}
          and{' '}
          <Link
            className="text-warning underline"
            href="/privacy"
            target="_blank"
          >
            {authLocalization.PRIVACY_POLICY}
          </Link>
          .
        </p>
      )}
    </main>
  );
}
