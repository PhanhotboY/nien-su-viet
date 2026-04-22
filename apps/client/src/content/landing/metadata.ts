import { getTranslations } from 'next-intl/server';

export async function getMetadata({ locale }: { locale: string }) {
  const t = await getTranslations({ locale, namespace: 'HomePage.Metadata' });

  return {
    title: t('title'),
    description: t('description'),
    logo: t('logo'),
    logoDark: t('logoDark'),
    social: {
      facebook: 'https://facebook.com/niensuviet',
      twitter: 'https://x.com/niensuviet',
      instagram: 'https://instagram.com/niensuviet',
      threads: 'https://threads.com/@niensuviet',
    },
    // address: {
    //   street: t('address.street'),
    //   district: t('address.district'),
    //   province: t('address.province'),
    //   country: t('address.country'),
    // },
    // msisdn: t('msisdn'),
    email: 'niensuviet@gmail.com',
    // map: t('map'),
  };
}
