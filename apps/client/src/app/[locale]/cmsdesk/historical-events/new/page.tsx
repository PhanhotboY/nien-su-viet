import { createEvent } from '@/services/historical-event.service';
import { EventForm } from '@/components/cmsdesk/events/event-form';

export default function NewEventPage() {
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
        onSubmit={createEvent}
        submitLabel="Create Event"
      />
    </div>
  );
}
