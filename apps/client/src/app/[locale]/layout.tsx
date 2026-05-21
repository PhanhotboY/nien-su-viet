import { Inter } from 'next/font/google';
import { Providers } from './provider';

import { ScrollToTop } from '@/components/ScrollToTop';

import '../../styles/globals.css';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { NextIntlClientProvider } from 'next-intl';
import { API_ENDPOINT } from '@/lib/config';
import { Suspense } from 'react';

const locales = ['vi', 'en'] as const;
export async function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}
const inter = Inter({
  display: 'swap',
  subsets: ['latin'],
  preload: true,
});

export default async function AdminLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}>) {
  const { locale } = await params;
  setRequestLocale(locale);
  const messages = await getMessages({ locale });

  return (
    <html lang={locale} className={inter.className}>
      <head>
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin="anonymous"
        />
        <link rel="preconnect" href={API_ENDPOINT} />
      </head>
      <body className={`antialiased flex min-h-svh flex-col`}>
        <NextIntlClientProvider locale={locale} messages={messages}>
          <Suspense fallback={null}>
            <Providers>
              {children}

              <ScrollToTop />
            </Providers>
          </Suspense>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
