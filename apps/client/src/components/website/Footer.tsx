'use client';

import Link from 'next/link';
import Image from 'next/image';
import { Mail, Phone, MapPin } from 'lucide-react';
import { getFooterNavItems } from '@/content/menus/footer-nav-items';
import {
  FacebookIcon,
  YoutubeIcon,
  TiktokIcon,
  ZaloIcon,
} from '@/icons/socials';
import { useTranslations } from 'next-intl';
import { useTheme } from 'next-themes';

interface FooterProps {
  title: string;
  logo: string;
  logoDark: string;
  description: string;
  social: {
    facebook?: string;
    youtube?: string;
    tiktok?: string;
    zalo?: string;
  };
  email: string;
  msisdn?: string;
  address: {
    street?: string;
    district?: string;
    province?: string;
  };
  navItems: Awaited<ReturnType<typeof getFooterNavItems>>;
}

export const Footer = ({
  title,
  logo,
  logoDark,
  description,
  social,
  email,
  msisdn,
  address,
  navItems,
}: FooterProps) => {
  const t = useTranslations('HomePage.Footer');
  const { theme } = useTheme();
  const logoToUse = theme === 'dark' ? logoDark : logo;

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
            <Image
              src={logoToUse}
              alt={title}
              width={100}
              height={100}
              className="mr-2 object-contain"
            />
          </Link>
          <p className="mt-4 text-sm text-muted-foreground">{description}</p>
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
                  <FacebookIcon className="h-6 w-6" />
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
                  <YoutubeIcon className="h-6 w-6" />
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
                  className="opacity-60 hover:opacity-100 flex items-center gap-2"
                >
                  <TiktokIcon className="h-6 w-6" />
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
                  className="opacity-60 hover:opacity-100 flex items-center gap-2"
                >
                  <ZaloIcon className="h-6 w-6" />
                  Zalo
                </Link>
              </div>
            )}
          </div>
        )}

        <div className="flex flex-col gap-2">
          <h3 className="font-bold text-lg">{t('see-more')}</h3>
          {navItems
            .sort((a, b) => a.order - b.order)
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

        <div className="flex flex-col gap-2">
          <h3 className="font-bold text-lg">{t('contact')}</h3>
          <div>
            <Link
              href={`mailto:${email}`}
              className="opacity-60 hover:opacity-100 flex items-center gap-2 text-sm"
            >
              <Mail className="h-4 w-4" />
              {email}
            </Link>
          </div>

          {msisdn && (
            <div>
              <Link
                href={`tel:${msisdn}`}
                className="opacity-60 hover:opacity-100 flex items-center gap-2 text-sm"
              >
                <Phone className="h-4 w-4" />
                {msisdn}
              </Link>
            </div>
          )}

          {(address?.street || address?.district || address?.province) && (
            <div className="opacity-60 flex items-start gap-2 text-sm">
              <MapPin className="h-4 w-4 mt-0.5 shrink-0" />
              <span>
                {address.street && (
                  <>
                    {address.street},
                    <br />
                  </>
                )}
                {address.district && (
                  <>
                    {address.district},
                    <br />
                  </>
                )}
                {address.province}
              </span>
            </div>
          )}
        </div>
      </section>

      <section className="container pb-14 text-center">
        <h3>{t('copyright')}</h3>
      </section>
    </footer>
  );
};
