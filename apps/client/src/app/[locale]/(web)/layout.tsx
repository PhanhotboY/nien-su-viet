import type { Metadata } from 'next';
import { Navbar } from '@/components/website/Navbar';
import { Footer } from '@/components/website/Footer';
import '@/styles/globals.css';
import { genMetadata } from '@/lib/metadata.lib';
import { getMetadata } from '@/content/landing/metadata';
import { getHeaderNavItems } from '@/content/menus/header-nav-items';
import { getFooterNavItems } from '@/content/menus/footer-nav-items';
import GoogleAnalytics from '@/components/website/GoogleAnalytics';

export const revalidate = 60; // Revalidate every 60 seconds

type Props = {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
};
export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const { locale } = await params;
    const metadata = await getMetadata({ locale });

    const title = metadata.title;
    const description = metadata.description;
    const logo = metadata.logo;

    const keywords =
      locale === 'vi'
        ? [
            'lịch sử Việt Nam',
            'sự kiện lịch sử',
            'niên sử Việt',
            'timeline lịch sử',
            'nhân vật lịch sử',
            'triều đại Việt Nam',
            'văn hóa Việt Nam',
          ]
        : [
            'Vietnamese history',
            'historical events',
            'Vietnam timeline',
            'historical figures',
            'Vietnamese dynasties',
            'Vietnam culture',
          ];

    return genMetadata({
      title,
      description,
      locale,
      logo,
      path: '/',
      keywords,
    });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path: '/' });
  }
}

export default async function RootLayout({
  children,
  params,
}: Readonly<{
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}>) {
  const { locale } = await params;
  const appData = await getMetadata({ locale });
  const headerNavItems = await getHeaderNavItems({ locale });
  const footerNavItems = await getFooterNavItems({ locale });

  return (
    <>
      <Navbar
        appTitle={appData.title}
        appLogo={appData.logo}
        appLogoDark={appData.logoDark}
        navItems={headerNavItems}
      />

      <GoogleAnalytics />
      {children}

      <Footer
        title={appData.title}
        logo={appData.logo}
        logoDark={appData.logoDark}
        description={appData.description}
        social={{
          facebook: appData.social.facebook,
          twitter: appData.social.twitter,
          instagram: appData.social.instagram,
          threads: appData.social.threads,
        }}
        email={appData.email}
        // msisdn={appData.msisdn}
        // address={{
        //   street: appData.address.street,
        //   district: appData.address.district,
        //   province: appData.address.province,
        // }}
        navItems={footerNavItems}
      />
    </>
  );
}
