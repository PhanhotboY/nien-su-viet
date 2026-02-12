import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';
import { getTranslations } from 'next-intl/server';

export default async function HomeTimePage() {
  const t = await getTranslations('EventPage');
  return (
    <main className="container mx-auto py-4 md:py-8 px-4">
      <h1 className="text-3xl md:text-5xl text-center font-bold leading-12 md:leading-20 mb-4 md:mb-8">
        {t('title')}
      </h1>

      <HistoricalEventTimeline />
    </main>
  );
}
