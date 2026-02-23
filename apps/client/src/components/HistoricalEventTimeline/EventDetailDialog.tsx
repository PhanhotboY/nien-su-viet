import { VisItem } from '@/interfaces/vis.interface';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from '../ui/dialog';
import useSWR from 'swr';
import { swrFetcher } from '@/helper/swrFetcher';
import { Skeleton } from '../ui/skeleton';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { toEventPeriodString } from '@/helper/date';
import { IApiResponse } from '../../interfaces/response.interface';
import TextRenderer from '../TextRenderer';
import Link from 'next/link';
import { Button } from '../ui/button';
import { useEffect } from 'react';
import { useTranslations } from 'next-intl';

// Custom React component for item details
export function EventDetailDialog({
  eventId,
  open,
  onOpenChange,
}: {
  eventId: string | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  useEffect(() => {
    mutate();
  }, [eventId, open]);
  const t = useTranslations('EventPage');
  const tshared = useTranslations('Shared');

  const { data, error, isLoading, mutate } = useSWR<
    IApiResponse<components['schemas']['HistoricalEventPreviewResponseDto']>
  >(eventId ? `/api/historical-events/${eventId}/preview` : null, swrFetcher);

  if (error) {
    return (
      <Dialog open={open} onOpenChange={onOpenChange}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Error</DialogTitle>
            <DialogDescription>{t('error-loading-details')}</DialogDescription>
          </DialogHeader>
        </DialogContent>
      </Dialog>
    );
  }

  const { data: eventData } = data || {};
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-h-[80vh] overflow-y-auto">
        {!eventData || isLoading ? (
          <>
            <DialogHeader>
              <DialogTitle className="flex items-center gap-2" asChild>
                <Skeleton className="h-6 w-48" />
              </DialogTitle>

              <DialogDescription asChild>
                <Skeleton className="h-4 w-32" />
              </DialogDescription>
            </DialogHeader>

            <div className="space-y-4">
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-full" />
              <Skeleton className="h-4 w-3/4" />
            </div>
          </>
        ) : (
          <>
            <DialogHeader>
              <DialogTitle className="text-2xl flex items-center gap-2">
                {eventData.name}
              </DialogTitle>
              <DialogDescription>
                {toEventPeriodString(eventData, tshared)}
              </DialogDescription>
            </DialogHeader>

            {eventData.excerpt ? (
              <TextRenderer content={eventData.excerpt} />
            ) : (
              t('empty-excerpt')
            )}

            <Button asChild>
              <Link href={`/su-kien/${eventData.id}`}>{t('view-detail')}</Link>
            </Button>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
}
