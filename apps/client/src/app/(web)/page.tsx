import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';

export default async function HomeTimePage() {
  return (
    <main className="container mx-auto py-4 md:py-8 px-4">
      <h1 className="text-3xl md:text-5xl text-center font-bold leading-12 md:leading-20 mb-4 md:mb-8">
        Lịch sử Việt Nam
      </h1>

      <HistoricalEventTimeline />
    </main>
  );
}
