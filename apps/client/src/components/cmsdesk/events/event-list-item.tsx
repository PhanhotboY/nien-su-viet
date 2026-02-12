'use client';

import { Button } from '@/components/ui/button';
import { Calendar, User, Edit, Trash2 } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { formatHistoricalEventDate } from '@/helper/date';
import { useTranslations } from 'next-intl';

type IHistoricalEvent =
  components['schemas']['HistoricalEventBriefResponseDto'];
interface EventListItemProps {
  event: IHistoricalEvent;
  onDelete: (id: string) => void;
}

export function EventListItem({ event, onDelete }: EventListItemProps) {
  const t = useTranslations('EventPage');

  return (
    <div className="group flex gap-4 rounded-lg border bg-card p-4 transition-all hover:shadow-md">
      <div className="relative h-24 w-32 flex-shrink-0 overflow-hidden rounded-md bg-muted">
        {event.thumbnail ? (
          <Image
            src={event.thumbnail}
            alt={event.name}
            fill
            className="object-cover transition-transform group-hover:scale-105"
          />
        ) : (
          <div className="flex h-full items-center justify-center bg-gradient-to-br from-primary/10 to-primary/5">
            <Calendar className="h-8 w-8 text-muted-foreground/40" />
          </div>
        )}
      </div>

      <div className="flex flex-1 flex-col justify-between">
        <div>
          <div className="mb-1 flex items-start justify-between gap-2">
            <h3 className="line-clamp-1 text-lg font-semibold">{event.name}</h3>
          </div>

          <div className="mb-2 flex items-center gap-4 text-sm text-muted-foreground">
            <div className="flex items-center gap-1.5">
              <Calendar className="h-4 w-4" />
              <span>
                {formatHistoricalEventDate(
                  event.fromDateType,
                  t('approximate'),
                  event.fromYear,
                  event.fromMonth,
                  event.fromDay,
                )}
                {' - '}
                {formatHistoricalEventDate(
                  event.toDateType,
                  t('approximate'),
                  event.toYear,
                  event.toMonth,
                  event.toDay,
                )}
              </span>
            </div>
            {/* <div className="flex items-center gap-1.5">
              <User className="h-4 w-4" />
              <span>{event.author?.name || 'Unknown'}</span>
            </div> */}
          </div>

          <p className="line-clamp-2 text-sm text-muted-foreground">
            {/* {event.content.replace(/<[^>]*>/g, '')} */}
          </p>
        </div>

        <div className="mt-3 flex items-center justify-between">
          <div className="flex flex-wrap gap-1">
            {/* {event.categories && event.categories.length > 0 ? (
              <>
                {event.categories.slice(0, 4).map((category) => (
                  <Badge
                    key={category.id}
                    variant="secondary"
                    className="text-xs"
                  >
                    {category.name}
                  </Badge>
                ))}
                {event.categories.length > 4 && (
                  <Badge variant="outline" className="text-xs">
                    +{event.categories.length - 4}
                  </Badge>
                )}
              </>
            ) : null} */}
          </div>

          <div className="flex gap-2">
            <Button asChild size="sm" variant="outline">
              <Link href={`/cmsdesk/historical-events/${event.id}/edit`}>
                <Edit className="mr-2 h-4 w-4" />
                Edit
              </Link>
            </Button>
            <Button
              size="sm"
              variant="outline"
              className="text-destructive hover:bg-destructive/10 hover:text-destructive"
              onClick={() => onDelete(event.id)}
            >
              <Trash2 className="mr-2 h-4 w-4" />
              Delete
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
