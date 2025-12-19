import { HistoricalEventTimeline } from '@/components/HistoricalEventTimeline';
import { getEvents } from '@/services/historical-event.service';

export default async function TimelineLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <main className="container mx-auto py-8 px-4">
      <h1 className="text-4xl text-center font-bold mb-8">Lịch sử Việt Nam</h1>

      <HistoricalEventTimeline />
      {children}
    </main>
  );
}
