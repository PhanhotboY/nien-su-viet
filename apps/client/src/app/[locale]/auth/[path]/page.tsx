import { Metadata } from 'next';
import Link from '@/i18n/navigation';
import { Auth } from '@/components/auth';

import { authLegalCopy as authLegalCopyVi } from '@/localization/vi/auth-localization';
import { authLegalCopy as authLegalCopyEn } from '@/localization/en/auth-localization';
import { getTranslations } from 'next-intl/server';

export const metadata: Metadata = {
  robots: {
    index: false,
  },
};

export default async function AuthPage({
  params,
}: {
  params: Promise<{ path: string; locale: string }>;
}) {
  const { path, locale } = await params;
  const tshared = await getTranslations({ locale, namespace: 'Shared' });
  const authLegalCopy = locale === 'vi' ? authLegalCopyVi : authLegalCopyEn;

  if (!path) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
      </div>
    );
  }

  return (
    <main className="h-full container flex grow flex-col items-center justify-center gap-4 self-center p-4 md:p-6">
      <Auth path={path} />

      {!['callback', 'sign-out'].includes(path) && (
        <p className="w-3xs text-center text-muted-foreground text-xs">
          {authLegalCopy.BY_CONTINUING_YOU_AGREE}{' '}
          <Link
            className="text-warning underline"
            href="/terms"
            target="_blank"
          >
            {authLegalCopy.TERMS_OF_SERVICE}
          </Link>{' '}
          {tshared('and')}{' '}
          <Link
            className="text-warning underline"
            href="/privacy"
            target="_blank"
          >
            {authLegalCopy.PRIVACY_POLICY}
          </Link>
          .
        </p>
      )}
    </main>
  );
}
