import { getTranslations } from 'next-intl/server';

export async function getHeaderNavItems() {
  const t = await getTranslations('Shared');

  return [
    // {
    //   order: 1,
    //   link_new_tab: false,
    //   link_url: '/gioi-thieu',
    //   link_label: t('about-us'),
    // },
    {
      order: 2,
      link_new_tab: false,
      link_url: '/blog',
      link_label: t('blog'),
    },
    // {
    //   order: 3,
    //   link_new_tab: false,
    //   link_url: '/lien-he',
    //   link_label: t('contact'),
    // },
  ];
}
