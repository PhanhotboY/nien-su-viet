import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';

export default async function HomeTimePage() {
  return (
    <main className="container mx-auto py-8 px-4">
      <h1 className="text-4xl text-center font-bold mb-8">Lịch sử Việt Nam</h1>

      <HistoricalEventTimeline />
    </main>
  );
}
