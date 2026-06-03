'use client';

import { AuthProvider } from '@/components/auth';
import Link from 'next/link';
import { useRouter, useSearchParams } from 'next/navigation';
import { ThemeProvider } from 'next-themes';
import { type ReactNode } from 'react';
import { Toaster } from 'sonner';
import { authClient } from '@/lib/auth-client';
import { authLocalization as authLocalizationVi } from '@/localization/vi/auth-localization';
import { authLocalization as authLocalizationEn } from '@/localization/en/auth-localization';
import { useLocale } from 'next-intl';

export function Providers({
  children,
}: {
  children: ReactNode;
}) {
  const router = useRouter();
  const searchParams = useSearchParams();
  const locale = useLocale();
  const authLocalization =
    locale === 'vi' ? authLocalizationVi : authLocalizationEn;
  const basePaths = {
    auth: `/${locale}/auth`,
    settings: `/${locale}/settings`,
    organization: `/${locale}/organization`,
  };
  const redirectTo = searchParams?.get('redirectTo') || `/${locale}`;

  return (
    <ThemeProvider
      attribute="class"
      defaultTheme="system"
      enableSystem
      disableTransitionOnChange
      storageKey="nsv-ui-theme"
    >
      <AuthProvider
        authClient={authClient}
        navigate={({ to, replace }) =>
          replace ? router.replace(to) : router.push(to)
        }
        basePaths={basePaths}
        redirectTo={redirectTo}
        Link={Link}
        // social={{
        //   providers: ['google'],
        // }}
        localization={authLocalization}
        emailAndPassword={{
          minPasswordLength: 1,
        }}
      >
        {children}

        <Toaster position="top-right" richColors />
      </AuthProvider>
    </ThemeProvider>
  );
}
