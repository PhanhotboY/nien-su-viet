'use client';

import { AuthUIProvider } from '@daveyplate/better-auth-ui';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { ThemeProvider } from 'next-themes';
import { type ReactNode } from 'react';
import { Toaster } from 'sonner';
import { authClient, signInWithGoogle } from '@/lib/auth-client';
import { authLocalization as authLocalizationVi } from '@/localization/vi/auth-localization';
import { authLocalization as authLocalizationEn } from '@/localization/en/auth-localization';
import { useLocale } from 'next-intl';

export function Providers({ children }: { children: ReactNode }) {
  const router = useRouter();
  const locale = useLocale();
  const authLocalization =
    locale === 'vi' ? authLocalizationVi : authLocalizationEn;

  return (
    <ThemeProvider
      attribute="class"
      defaultTheme="system"
      enableSystem
      disableTransitionOnChange
      storageKey="nsv-ui-theme"
    >
      <AuthUIProvider
        authClient={authClient}
        navigate={router.push}
        replace={router.replace}
        basePath={`/${locale}/auth`}
        onSessionChange={() => {
          // Clear router cache (protected routes)
          router.refresh();
        }}
        Link={Link}
        // social={{
        //   providers: ['google'],
        // }}
        localization={authLocalization}
        account={{
          basePath: `/${locale}/account`,
        }}
      >
        {children}

        <Toaster position="top-right" richColors />
      </AuthUIProvider>
    </ThemeProvider>
  );
}
