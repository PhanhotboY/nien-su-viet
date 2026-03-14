import { getTranslations } from 'next-intl/server';

export async function getMetadata({ locale }: { locale: string }) {
  const t = await getTranslations({ locale, namespace: 'HomePage.Metadata' });

  return {
    title: t('title'),
    description: t('description'),
    logo: t('logo'),
    logoDark: t('logoDark'),
    social: {
      facebook: t('social.facebook'),
      twitter: t('social.twitter'),
      instagram: t('social.instagram'),
      linkedin: t('social.linkedin'),
      youtube: t('social.youtube'),
      tiktok: t('social.tiktok'),
      zalo: t('social.zalo'),
    },
    address: {
      street: t('address.street'),
      district: t('address.district'),
      province: t('address.province'),
      country: t('address.country'),
    },
    msisdn: t('msisdn'),
    email: t('email'),
    map: t('map'),
  };
}
