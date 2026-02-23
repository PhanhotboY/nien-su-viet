import { getTranslations } from 'next-intl/server';

export async function getFooterNavItems() {
  const t = await getTranslations('Shared');

  return [
    {
      order: 1,
      link_new_tab: false,
      link_url: '/lien-he',
      link_label: t('contact'),
    },
    {
      order: 2,
      link_new_tab: false,
      link_url: '/privacy',
      link_label: t('privacy-policy'),
    },
    {
      order: 3,
      link_new_tab: false,
      link_url: '/terms',
      link_label: t('terms-of-service'),
    },
  ];
}
