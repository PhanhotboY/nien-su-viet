import { getEvents } from '@/services/historical-event.service';

export async function generateStaticParams() {
  try {
    const { data: events } = await getEvents({
      page: '1',
      limit: '1000',
    });
    return events.map((event) => ({
      eventId: event.id,
    }));
  } catch (error) {
    console.error('Error fetching events for static params:', error);
    return [];
  }
}

export default function EventDetailLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <>{children}</>;
}
