import { Inter } from 'next/font/google';
import { Providers } from './provider';

import { ScrollToTop } from '@/components/ScrollToTop';

import '@/styles/globals.css';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { NextIntlClientProvider } from 'next-intl';

const locales = ['vi', 'en'] as const;
export async function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}
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
  const { locale } = await params;
  setRequestLocale(locale);
  const messages = await getMessages({ locale });

  return (
    <html lang={locale}>
      <body className={`${inter.variable} antialiased flex min-h-svh flex-col`}>
        <NextIntlClientProvider locale={locale} messages={messages}>
          <Providers>
            {children}

            <ScrollToTop />
          </Providers>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
