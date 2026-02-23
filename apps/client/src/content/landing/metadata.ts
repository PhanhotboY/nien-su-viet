import { getTranslations } from 'next-intl/server';

export async function getMetadata() {
  const t = await getTranslations('HomePage.Metadata');

  return {
    title: t('title'),
    description: t('description'),
    logo: '/assets/logo/nsv-logo.webp',
    logoDark: '/assets/logo/nsv-logo-dark.webp',
    social: {
      facebook: 'https://facebook.com/niensuviet',
      // youtube: 'https://youtube.com/nien-suu-viet',
      // tiktok: 'https://tiktok.com/nien-suu-viet',
      // zalo: 'https://zalo.me/nien-suu-viet',
    },
    address: {
      street: '123 Nguyen Trai',
      district: 'Thu Duc',
      province: 'Ho Chi Minh City',
    },
    // msisdn: '+84987654321',
    email: 'nsv-contact@phannd.me',
    map: 'https://maps.google.com/?q=123+Nguyen+Trai+Thu+Duc+Ho+Chi+Minh+City',
  };
}
