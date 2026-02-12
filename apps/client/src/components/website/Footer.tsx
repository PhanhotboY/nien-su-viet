import Link from 'next/link';
import Image from 'next/image';
import { LogoIcon } from '../Icons';
import { Facebook, Youtube, Mail, Phone, MapPin } from 'lucide-react';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { getTranslations } from 'next-intl/server';

interface FooterProps {
  appTitle?: string;
  appLogo?: string;
  appDescription?: string;
  social?: {
    facebook?: string;
    youtube?: string;
    tiktok?: string;
    zalo?: string;
  };
  email?: string;
  phone?: string;
  address?: {
    street?: string;
    district?: string;
    province?: string;
  };
  navItems?: components['schemas']['HeaderNavItemDto'][] | null;
}

export const Footer = async ({
  appTitle = 'ShadcnUI/React',
  appLogo,
  appDescription,
  social,
  email,
  phone,
  address,
  navItems,
}: FooterProps) => {
  const t = await getTranslations('HomePage.Footer');

  return (
    <footer id="footer">
      <hr className="w-11/12 mx-auto" />

      <section className="container py-20 grid grid-cols-2 md:grid-cols-4 xl:grid-cols-6 gap-x-12 gap-y-8">
        <div className="col-span-full xl:col-span-2">
          <Link
            rel="noreferrer noopener"
            href="/"
            className="font-bold text-xl flex items-center"
          >
            {appLogo ? (
              <Image
                src={appLogo}
                alt={appTitle}
                width={32}
                height={32}
                className="mr-2 h-8 w-8 object-contain"
              />
            ) : (
              <LogoIcon />
            )}
          </Link>
          {appDescription && (
            <p className="mt-4 text-sm text-muted-foreground">
              {appDescription}
            </p>
          )}
        </div>

        {(social?.facebook ||
          social?.youtube ||
          social?.tiktok ||
          social?.zalo) && (
          <div className="flex flex-col gap-2">
            <h3 className="font-bold text-lg">{t('follow-us')}</h3>
            {social?.facebook && (
              <div>
                <Link
                  rel="noreferrer noopener"
                  href={social.facebook}
                  target="_blank"
                  className="opacity-60 hover:opacity-100 flex items-center gap-2"
                >
                  <Facebook className="h-4 w-4" />
                  Facebook
                </Link>
              </div>
            )}

            {social?.youtube && (
              <div>
                <Link
                  rel="noreferrer noopener"
                  href={social.youtube}
                  target="_blank"
                  className="opacity-60 hover:opacity-100 flex items-center gap-2"
                >
                  <Youtube className="h-4 w-4" />
                  YouTube
                </Link>
              </div>
            )}

            {social?.tiktok && (
              <div>
                <Link
                  rel="noreferrer noopener"
                  href={social.tiktok}
                  target="_blank"
                  className="opacity-60 hover:opacity-100"
                >
                  TikTok
                </Link>
              </div>
            )}

            {social?.zalo && (
              <div>
                <Link
                  rel="noreferrer noopener"
                  href={social.zalo}
                  target="_blank"
                  className="opacity-60 hover:opacity-100"
                >
                  Zalo
                </Link>
              </div>
            )}
          </div>
        )}

        <div className="flex flex-col gap-2">
          {/*<h3 className="font-bold text-lg">Platforms</h3>*/}
          {navItems
            ?.sort((a, b) => a.order - b.order)
            .map((item, index) => (
              <div key={index}>
                <Link
                  rel="noreferrer noopener"
                  href={item.link_url}
                  className="opacity-60 hover:opacity-100"
                  target={item.link_new_tab ? '_blank' : ''}
                >
                  {item.link_label}
                </Link>
              </div>
            ))}
        </div>

        {(email ||
          phone ||
          address?.street ||
          address?.district ||
          address?.province) && (
          <div className="flex flex-col gap-2">
            <h3 className="font-bold text-lg">{t('contact')}</h3>
            {email && (
              <div>
                <Link
                  href={`mailto:${email}`}
                  className="opacity-60 hover:opacity-100 flex items-center gap-2 text-sm"
                >
                  <Mail className="h-4 w-4" />
                  {email}
                </Link>
              </div>
            )}

            {phone && (
              <div>
                <Link
                  href={`tel:${phone}`}
                  className="opacity-60 hover:opacity-100 flex items-center gap-2 text-sm"
                >
                  <Phone className="h-4 w-4" />
                  {phone}
                </Link>
              </div>
            )}

            {(address?.street || address?.district || address?.province) && (
              <div className="opacity-60 flex items-start gap-2 text-sm">
                <MapPin className="h-4 w-4 mt-0.5 shrink-0" />
                <span>
                  {address.street && (
                    <>
                      {address.street}
                      <br />
                    </>
                  )}
                  {address.district && (
                    <>
                      {address.district}
                      <br />
                    </>
                  )}
                  {address.province}
                </span>
              </div>
            )}
          </div>
        )}
      </section>

      <section className="container pb-14 text-center">
        <h3>{t('copyright')}</h3>
      </section>
    </footer>
  );
};
