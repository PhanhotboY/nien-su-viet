'use client';

import { EventDetailDialog } from '@/components/HistoricalEventTimeline/EventDetailDialog';
import { useRouter, useSearchParams } from 'next/navigation';

export default function EventPreview() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const eventId = searchParams.get('eventId');
  if (!eventId) return null;

  return (
    <EventDetailDialog
      eventId={eventId}
      open={true}
      onOpenChange={() => router.back()}
    />
  );
}
