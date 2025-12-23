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

export async function generateMetadata(): Promise<Metadata> {
  try {
    const response = await getAppInfo();
    const appData = response.data;

    const title = appData?.title || 'Nien Su Viet';
    const description =
      appData?.description || 'Vietnam history timeline website';
    const logo = appData?.logo;

    return {
      title: title,
      description: description,
      openGraph: {
        title: title,
        description: description,
        type: 'website',
        locale: 'vi_VN',
        siteName: title,
        ...(logo && { images: [{ url: logo, alt: title }] }),
      },
      twitter: {
        card: 'summary_large_image',
        title: title,
        description: description,
        ...(logo && { images: [logo] }),
      },
      ...(logo && { icons: { icon: logo, apple: logo } }),
    };
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return {
      title: title,
      description: description,
      openGraph: {
        title: title,
        description: description,
        type: 'website',
        locale: 'vi_VN',
        siteName: title,
      },
      twitter: {
        card: 'summary_large_image',
        title: title,
        description: description,
      },
    };
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
      />
    </>
  );
}
