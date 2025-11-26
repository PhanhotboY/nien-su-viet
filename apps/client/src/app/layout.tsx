import { Inter } from 'next/font/google';
import { Providers } from './provider';

import { ScrollToTop } from '@/components/ScrollToTop';

import '@/styles/globals.css';

const inter = Inter({
  variable: '--font-inter',
  subsets: ['latin'],
});

export default async function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.variable} antialiased flex min-h-svh flex-col`}>
        <Providers>
          {children}

          <ScrollToTop />
        </Providers>
      </body>
    </html>
  );
}
