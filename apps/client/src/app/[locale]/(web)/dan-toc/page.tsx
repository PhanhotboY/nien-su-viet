import { Metadata } from 'next';
import { getTranslations } from 'next-intl/server';

import { getMetadata } from '@/content/landing/metadata';
import { genMetadata } from '@/lib/metadata.lib';
import { ethnicGroups } from '@/data/vi/vietnamese-ethnic-groups';
import { EthnicGroupsGrid } from '@/components/website/ethnic-groups/ethnic-groups-grid';

type Props = {
  params: Promise<{ locale: string }>;
};

export async function generateMetadata({ params }: Props): Promise<Metadata> {
  const path = '/dan-toc';
  try {
    const { locale } = await params;
    const metadata = await getMetadata({ locale });
    const t = await getTranslations({ locale, namespace: 'EthnicGroupPage' });

    return genMetadata({
      title: `${t('title')} - ${metadata.title}`,
      description: t('description'),
      locale,
      path,
      logo: metadata.logo,
    });
  } catch (error) {
    const title = 'Nien Su Viet';
    const description = 'Vietnam history timeline website';

    return genMetadata({ title, description, locale: 'vi', path });
  }
}

export default async function DanTocPage({ params }: Props) {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'EthnicGroupPage' });

  return (
    <div className="min-h-screen bg-background">
      {/* Hero Section */}
      <div className="relative overflow-hidden border-b border-border/40">
        <div className="absolute inset-0 -top-40 bg-linear-to-b from-primary/5 via-transparent to-transparent pointer-events-none" />
        <div className="relative container mx-auto px-4 py-16 md:py-24">
          <div className="max-w-3xl">
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold tracking-tight leading-tight mb-4">
              {t('title')}
            </h1>
            <p className="text-lg text-muted-foreground mb-2">
              {t('description')}
            </p>
            <p className="text-sm text-muted-foreground">
              {t('explore')}{' '}
              <span className="font-semibold text-foreground">
                {ethnicGroups.length} {t('ethnicGroups')}
              </span>{' '}
              {t('richCulture')} {t('ofVietnam')}
            </p>
          </div>
        </div>
      </div>

      {/* Grid with Search */}
      <EthnicGroupsGrid
        groups={ethnicGroups}
        searchPlaceholder={t('searchPlaceholder')}
        foundResults={t('foundResults')}
        results={t('results')}
        noResultsFound={t('noResultsFound')}
        tryAnotherKeyword={t('tryAnotherKeyword')}
      />

      {/* Credit Section */}
      <div className="border-t border-border/40 bg-muted/20">
        <div className="container mx-auto px-4 py-8">
          <div className="mx-auto text-center">
            <p className="text-xs font-semibold text-muted-foreground uppercase tracking-wide mb-2">
              {t('dataSource')}
            </p>
            <p className="text-sm text-muted-foreground leading-relaxed">
              {t('creditText')}{' '}
              <a
                href="http://www.cema.gov.vn/gioi-thieu/cong-dong-54-dan-toc.htm"
                target="_blank"
                rel="noopener noreferrer"
                className="text-primary hover:underline"
              >
                CỔNG THÔNG TIN ĐIỆN TỬ BỘ DÂN TỘC VÀ TÔN GIÁO
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
