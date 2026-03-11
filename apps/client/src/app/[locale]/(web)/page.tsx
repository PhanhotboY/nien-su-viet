import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';
import { EventStatistics } from '@/components/website/EventStatistics';
import { QuickTips } from '@/components/website/QuickTips';
import { FeaturedPeriods } from '@/components/website/FeaturedPeriods';
import { getTranslations } from 'next-intl/server';
import { getMetadata } from '@/content/landing/metadata';
import { getImportantEvents } from '@/content/landing/important-events';
import { CLIENT_HOST } from '@/lib/config';

export default async function HomeTimePage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'EventPage' });
  const metadata = await getMetadata({ locale });
  const importantEvents = await getImportantEvents();

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
    address: {
      '@type': 'PostalAddress',
      streetAddress: metadata.address.street,
      addressLocality: metadata.address.district,
      addressRegion: metadata.address.province,
      addressCountry: 'VN',
    },
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

  return (
    <>
      {/* Structured Data for Rich Snippets */}
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{
          __html: JSON.stringify(websiteJsonLd),
        }}
      />
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{
          __html: JSON.stringify(organizationJsonLd),
        }}
      />
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{
          __html: JSON.stringify(itemListJsonLd),
        }}
      />
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{
          __html: JSON.stringify(breadcrumbJsonLd),
        }}
      />

      <main
        className="min-h-screen"
        itemScope
        itemType="https://schema.org/WebPage"
      >
        {/* Page Header — optimized for featured snippets */}
        <section className="container mx-auto py-6 md:py-10 px-4">
          <div className="text-center space-y-3 max-w-3xl mx-auto">
            <h1 className="text-4xl md:text-5xl font-bold" itemProp="name">
              <span className="bg-gradient-to-r from-primary via-primary/80 to-secondary text-transparent bg-clip-text">
                {t('title')}
              </span>
            </h1>
            <p className="text-lg text-muted-foreground" itemProp="description">
              {t('description')}
            </p>
          </div>
        </section>

        {/* Quick Tips Section */}
        <QuickTips />

        {/* Main Timeline Section */}
        <section className="container mx-auto py-4 md:py-6 px-4">
          <HistoricalEventTimeline />
        </section>

        {/* Statistics Section */}
        <EventStatistics />

        {/* Featured Historical Periods */}
        <FeaturedPeriods locale={locale} />
      </main>
    </>
  );
}
