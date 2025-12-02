'use client';

import { AuthUIProvider } from '@daveyplate/better-auth-ui';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { ThemeProvider } from 'next-themes';
import type { ReactNode } from 'react';
import { Toaster } from 'sonner';
import { authClient, signInWithGoogle } from '@/lib/auth-client';
import { CLIENT_HOST } from '@/lib/config';

export function Providers({ children }: { children: ReactNode }) {
  const router = useRouter();

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
        baseURL={CLIENT_HOST}
        onSessionChange={() => {
          // Clear router cache (protected routes)
          router.refresh();
        }}
        Link={Link}
        social={{
          providers: ['google'],
        }}
      >
        {children}

        <Toaster position="top-right" richColors />
      </AuthUIProvider>
    </ThemeProvider>
  );
}
