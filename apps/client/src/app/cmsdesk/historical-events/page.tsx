import Link from 'next/link';
import { Plus } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { deleteEvent } from '@/services/historical-event.service';

import EventList from '@/components/cmsdesk/events/event-list';

export default function HistoricalEventManagementPage() {
  return (
    <div className="container mx-auto space-y-6 py-6">
      {/* Header */}
      <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">
            Historical Events
          </h1>
          <p className="text-muted-foreground">
            Manage and organize historical events
          </p>
        </div>
        <Button asChild>
          <Link href="/cmsdesk/historical-events/new">
            <Plus className="mr-2 h-4 w-4" />
            Create Event
          </Link>
        </Button>
      </div>

      <EventList handleDelete={deleteEvent} />
    </div>
  );
}
