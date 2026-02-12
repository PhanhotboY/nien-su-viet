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
}: Readonly<{
  children: React.ReactNode;
}>) {
  const messages = await getMessages();

  return (
    <html lang="en">
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
