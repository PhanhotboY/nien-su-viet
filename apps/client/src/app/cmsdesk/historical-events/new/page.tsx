'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { createEvent } from '@/services/historical-event.service';
import { EventForm } from '@/components/cmsdesk/events/event-form';

import { toast } from 'sonner';
import { components } from '@nsv-interfaces/historical-event';

export default function NewEventPage() {
  const router = useRouter();
  const [categories, setCategories] = useState<any[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    loadCategories();
  }, []);

  const loadCategories = async () => {
    try {
      // const cats = await getCategories();
      setCategories([]);
    } catch (error) {
      toast.error('Failed to load categories');
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
      await createEvent(data);
      router.push('/cmsdesk/historical-events');
      toast.success('Event created successfully!');
    } catch (error) {
      toast.error('Failed to create event');
      console.error(error);
    } finally {
      setIsSubmitting(false);
    }
  };

  if (isLoading) {
    return (
      <div className="container mx-auto py-6">
        <div className="mb-6 flex items-center gap-4">
          <div className="h-8 w-8 animate-pulse rounded-md bg-muted" />
          <div className="h-8 w-48 animate-pulse rounded-md bg-muted" />
        </div>
        <div className="space-y-4">
          <div className="h-64 animate-pulse rounded-lg bg-muted" />
          <div className="h-64 animate-pulse rounded-lg bg-muted" />
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto space-y-6 py-6">
      {/* Header with Breadcrumb */}
      <div className="space-y-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">
            Create New Event
          </h1>
          <p className="text-muted-foreground">
            Add a new historical event to the timeline
          </p>
        </div>
      </div>

      {/* Form */}
      <EventForm
        categories={[]}
        onSubmit={handleSubmit}
        isSubmitting={isSubmitting}
        submitLabel="Create Event"
      />
    </div>
  );
}
