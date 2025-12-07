'use client';

import {
  getEvent,
  getCategories,
  updateEvent,
} from '@/services/historical-event.service';
import { EventForm } from '@/components/cmsdesk/events/event-form';
import { ArrowLeft } from 'lucide-react';
import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { useEffect, useState } from 'react';
import { components } from '@nsv-interfaces/historical-event';
import ErrorPage from './error';
import Loading from './loading';
import { useParams } from 'next/navigation';

export default function EditEventPage() {
  const { id: eventId } = useParams() as Record<string, string>;

  const [event, setEvent] = useState<
    components['schemas']['HistoricalEventDetailResponseDto'] | null
  >(null);
  // const [isSubmitting, setIsSubmitting] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  if (!eventId || error) {
    return <ErrorPage>{error}</ErrorPage>;
  }
  // getCategories(),

  useEffect(() => {
    (async () => {
      try {
        setIsLoading(true);
        // fetching from client since calling from server causes issues with cookies
        // ssr can't reset cookies when they expire
        const { data } = await getEvent(eventId);
        setEvent(data);
      } catch (err: any) {
        setError(err.message || 'Failed to load event data.');
      } finally {
        setIsLoading(false);
      }
    })();
  }, [eventId]);

  if (!event || isLoading) {
    return <Loading />;
  }

  return (
    <div className="container mx-auto space-y-6 py-6">
      {/* Header with Breadcrumb */}
      <div className="space-y-4">
        <Button variant="ghost" size="sm" asChild className="mb-2">
          <Link href="/cmsdesk/historical-events">
            <ArrowLeft className="mr-2 h-4 w-4" />
            Back to Events
          </Link>
        </Button>

        <div>
          <h1 className="text-3xl font-bold tracking-tight">Edit Event</h1>
          <p className="text-muted-foreground">
            Update the details of {event.name}
          </p>
        </div>
      </div>

      {/* Form */}
      <EventForm
        initialData={event}
        categories={[]}
        onSubmit={async (data) => {
          await updateEvent(eventId, data);
        }}
        submitLabel="Update Event"
      />
    </div>
  );
}
