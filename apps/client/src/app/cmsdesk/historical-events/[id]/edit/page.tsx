import { redirect } from 'next/navigation';
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
import { components } from '@nsv-interfaces/historical-event';

export default async function EditEventPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id: eventId } = await params;

  // const [categories, setCategories] = useState<IEventCategory[]>([]);
  // const [isSubmitting, setIsSubmitting] = useState(false);
  // const [isLoading, setIsLoading] = useState(true);
  // const [error, setError] = useState<string | null>(null);

  const [{ data: event }] = await Promise.all([
    getEvent(eventId),
    // getCategories(),
  ]);

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
          'use server';
          await updateEvent(eventId, data);
        }}
        submitLabel="Update Event"
      />
    </div>
  );
}
