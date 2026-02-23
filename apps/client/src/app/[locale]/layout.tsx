import { Inter } from 'next/font/google';
import { Providers } from './provider';

import { ScrollToTop } from '@/components/ScrollToTop';

import '@/styles/globals.css';
import { getMessages } from 'next-intl/server';
import { NextIntlClientProvider } from 'next-intl';

const inter = Inter({
  variable: '--font-inter',
  subsets: ['latin'],
});

export default async function AdminLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}>) {
  const messages = await getMessages();
  const { locale } = await params;

  return (
    <html lang={locale}>
      <head></head>

      <body className={`${inter.variable} antialiased flex min-h-svh flex-col`}>
        <NextIntlClientProvider messages={messages}>
          <Providers>
            {children}

            <ScrollToTop />
          </Providers>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
