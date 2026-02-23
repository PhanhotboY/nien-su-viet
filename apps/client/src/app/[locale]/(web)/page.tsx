import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';
import { EventStatistics } from '@/components/website/EventStatistics';
import { QuickTips } from '@/components/website/QuickTips';
import { FeaturedPeriods } from '@/components/website/FeaturedPeriods';
import { getTranslations } from 'next-intl/server';

export default async function HomeTimePage() {
  const t = await getTranslations('EventPage');
  return (
    <main className="min-h-screen">
      {/* Page Header */}
      <section className="container mx-auto py-6 md:py-10 px-4">
        <div className="text-center space-y-3 max-w-3xl mx-auto">
          <h1 className="text-4xl md:text-5xl font-bold">
            <span className="bg-gradient-to-r from-primary via-primary/80 to-secondary text-transparent bg-clip-text">
              {t('title')}
            </span>
          </h1>
          <p className="text-lg text-muted-foreground">{t('description')}</p>
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
      <FeaturedPeriods />
    </main>
  );
}
