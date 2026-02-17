import type { Metadata } from 'next';
import { Navbar } from '@/components/website/Navbar';
import { Footer } from '@/components/website/Footer';
import { CustomScripts } from '@/components/website/CustomScripts';
import {
  getAppInfo,
  getFooterNavItems,
  getHeaderNavItems,
} from '@/services/cms.service';
import '@/styles/globals.css';
import { getTranslations } from 'next-intl/server';
import { CLIENT_HOST } from '@/lib/config';
import { genMetadata } from '@/lib/metadata.lib';

type Props = {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
};
export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const locale = (await params).locale;
    const { data: appData } = await getAppInfo();

    const title = appData?.title || 'Nien Su Viet';
    const description =
      appData?.description || 'Vietnam history timeline website';
    const logo = appData?.logo;

    return genMetadata({ title, description, locale, logo, path: '/' });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path: '/' });
  }
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  let appData,
    headerNavItems,
    footerNavItems = null;

  try {
    [{ data: appData }, { data: headerNavItems }, { data: footerNavItems }] =
      await Promise.all([
        getAppInfo(),
        getHeaderNavItems(),
        getFooterNavItems(),
      ]);
    // appData = response.data;
  } catch (error) {
    console.error('Failed to fetch app info:', error);
  }

  console.log('app data: ', appData);
  return (
    <>
      <CustomScripts
        headScripts={appData?.head_scripts || undefined}
        bodyScripts={appData?.body_scripts || undefined}
      />

      <Navbar
        appTitle={appData?.title}
        appLogo={appData?.logo || undefined}
        navItems={headerNavItems}
      />

      {children}

      <Footer
        appTitle={appData?.title}
        appLogo={appData?.logo || undefined}
        appDescription={appData?.description || undefined}
        social={{
          facebook: appData?.social?.facebook || undefined,
          youtube: appData?.social?.youtube || undefined,
          tiktok: appData?.social?.tiktok || undefined,
          zalo: appData?.social?.zalo || undefined,
        }}
        email={appData?.email || undefined}
        phone={appData?.msisdn || undefined}
        address={{
          street: appData?.address?.street || undefined,
          district: appData?.address?.district || undefined,
          province: appData?.address?.province || undefined,
        }}
        navItems={footerNavItems}
      />
    </>
  );
}
