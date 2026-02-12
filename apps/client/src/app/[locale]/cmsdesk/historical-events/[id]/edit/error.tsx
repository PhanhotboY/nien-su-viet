'use client';

import { Button } from '@/components/ui/button';
import { ArrowLeft } from 'lucide-react';
import Link from 'next/link';

export default function Error({ children }: { children?: React.ReactNode }) {
  return (
    <div className="container mx-auto py-6">
      <div className="flex min-h-[400px] flex-col items-center justify-center rounded-lg border border-dashed bg-muted/20 p-8 text-center">
        <h3 className="mb-2 text-lg font-semibold">
          {children ? children : 'Event not found'}
        </h3>
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
