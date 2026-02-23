import type { Metadata } from 'next';
import { Navbar } from '@/components/website/Navbar';
import { Footer } from '@/components/website/Footer';
import '@/styles/globals.css';
import { genMetadata } from '@/lib/metadata.lib';
import { getMetadata } from '@/content/landing/metadata';
import { getHeaderNavItems } from '@/content/menus/header-nav-items';
import { getFooterNavItems } from '@/content/menus/footer-nav-items';

type Props = {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
};
export async function generateMetadata({ params }: Props): Promise<Metadata> {
  try {
    const locale = (await params).locale;
    const metadata = await getMetadata();

    const title = metadata.title;
    const description = metadata.description;
    const logo = metadata.logo;

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
  const appData = await getMetadata();
  const headerNavItems = await getHeaderNavItems();
  const footerNavItems = await getFooterNavItems();

  return (
    <>
      <Navbar
        appTitle={appData.title}
        appLogo={appData.logo}
        appLogoDark={appData.logoDark}
        navItems={headerNavItems}
      />

      {children}

      <Footer
        title={appData.title}
        logo={appData.logo}
        logoDark={appData.logoDark}
        description={appData.description}
        social={{
          facebook: appData.social.facebook,
          // youtube: appData.social.youtube,
          // tiktok: appData.social.tiktok,
          // zalo: appData.social.zalo,
        }}
        email={appData.email}
        // msisdn={appData.msisdn}
        address={{
          street: appData.address.street,
          district: appData.address.district,
          province: appData.address.province,
        }}
        navItems={footerNavItems}
      />
    </>
  );
}
