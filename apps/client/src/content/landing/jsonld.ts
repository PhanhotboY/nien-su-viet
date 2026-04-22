import { CLIENT_HOST } from '@/lib/config';
import { getMetadata } from './metadata';

function genJsonLd({
  metadata,
  importantEvents,
  locale,
}: {
  metadata: Awaited<ReturnType<typeof getMetadata>>;
  importantEvents: any[];
  locale: string;
}) {
  const websiteJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    name: metadata.title,
    url: `${CLIENT_HOST}/${locale}`,
    description: metadata.description,
    inLanguage: locale === 'vi' ? 'vi-VN' : 'en-US',
    potentialAction: {
      '@type': 'SearchAction',
      target: {
        '@type': 'EntryPoint',
        urlTemplate: `${CLIENT_HOST}/${locale}/su-kien?q={search_term_string}`,
      },
      'query-input': 'required name=search_term_string',
    },
  };

  const organizationJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'Organization',
    name: metadata.title,
    url: `${CLIENT_HOST}/${locale}`,
    logo: `${CLIENT_HOST}${metadata.logo}`,
    email: metadata.email,
    sameAs: [metadata.social.facebook],
  };

  const itemListJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'ItemList',
    name:
      locale === 'vi'
        ? 'Các giai đoạn lịch sử quan trọng của Việt Nam'
        : 'Important periods in Vietnamese history',
    numberOfItems: importantEvents.length,
    itemListElement: importantEvents.map((event, index) => ({
      '@type': 'ListItem',
      position: index + 1,
      name: event.name,
      description: event.description,
    })),
  };

  const breadcrumbJsonLd = {
    '@context': 'https://schema.org',
    '@type': 'BreadcrumbList',
    itemListElement: [
      {
        '@type': 'ListItem',
        position: 1,
        name: metadata.title,
        item: `${CLIENT_HOST}/${locale}`,
      },
    ],
  };

  return {
    websiteJsonLd,
    organizationJsonLd,
    itemListJsonLd,
    breadcrumbJsonLd,
  };
}

export { genJsonLd };
