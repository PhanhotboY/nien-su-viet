'use client';

import { useState, useEffect } from 'react';
import { useRouter, useParams } from 'next/navigation';
import {
  getEvent,
  getCategories,
  updateEvent,
} from '@/services/historical-event.service';
import { EventForm } from '@/components/cmsdesk/events/event-form';
import { ArrowLeft } from 'lucide-react';
import Link from 'next/link';
import { toast } from 'sonner';
import { Button } from '@/components/ui/button';
import { Skeleton } from '@/components/ui/skeleton';
import { components } from '@nsv-interfaces/historical-event';

export default function EditEventPage() {
  const router = useRouter();
  const params = useParams();
  const eventId = params.id as string;

  const [event, setEvent] = useState<
    components['schemas']['HistoricalEventDetailResponseDto'] | null
  >(null);
  // const [categories, setCategories] = useState<IEventCategory[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    loadData();
  }, [eventId]);

  const loadData = async () => {
    setIsLoading(true);
    try {
      const [eventRes] = await Promise.all([
        getEvent(eventId),
        // getCategories(),
      ]);
      setEvent(eventRes.data);
      // setCategories(categoriesData);
    } catch (error) {
      setError('Failed to load event');
      toast.error('Failed to load event');
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSubmit = async (
    data: components['schemas']['HistoricalEventBaseCreateDto'],
  ) => {
    setIsSubmitting(true);
    try {
      await updateEvent(eventId, data);
      router.push('/cmsdesk/historical-events');
      toast.success('Event updated successfully!');
    } catch (error) {
      toast.error('Failed to update event');
      console.error(error);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <div className="container mx-auto space-y-6 py-6">
        <div className="flex items-center gap-4">
          <Skeleton className="h-8 w-8" />
          <Skeleton className="h-8 w-48" />
        </div>
        <div className="space-y-4">
          <Skeleton className="h-64 w-full" />
          <Skeleton className="h-64 w-full" />
          <Skeleton className="h-48 w-full" />
        </div>
      </div>
    );
  }

  if (error || !event) {
    return (
      <div className="container mx-auto py-6">
        <div className="flex min-h-[400px] flex-col items-center justify-center rounded-lg border border-dashed bg-muted/20 p-8 text-center">
          <h3 className="mb-2 text-lg font-semibold">Event not found</h3>
          <p className="mb-4 text-sm text-muted-foreground">
            The event you're looking for doesn't exist or has been deleted.
          </p>
          <Button asChild>
            <Link href="/cmsdesk/historical-events">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Back to Events
            </Link>
          </Button>
        </div>
      </div>
    );
  }

  const initialData: components['schemas']['HistoricalEventBaseCreateDto'] = {
    name: event.name,
    content: event.content,
    fromYear: event.fromYear,
    fromMonth: event.fromMonth,
    fromDay: event.fromDay,
    toYear: event.toYear,
    toMonth: event.toMonth,
    toDay: event.toDay,
    thumbnailId: event.thumbnail?.id,
    // categoryIds: event.categories?.map((c) => c.id) || [],
  };

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
        initialData={initialData}
        categories={[]}
        onSubmit={handleSubmit}
        isSubmitting={isSubmitting}
        submitLabel="Update Event"
      />
    </div>
  );
}
