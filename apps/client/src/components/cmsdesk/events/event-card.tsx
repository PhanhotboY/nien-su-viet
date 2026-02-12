'use client';

import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Calendar, Edit, Trash2, User } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { formatHistoricalEventDate } from '@/helper/date';
import { useTranslations } from 'next-intl';

interface EventCardProps {
  event: components['schemas']['HistoricalEventBriefResponseDto'];
  onDelete: (id: string) => void;
}

export function EventCard({ event, onDelete }: EventCardProps) {
  const t = useTranslations('EventPage');

  return (
    <Card className="group overflow-hidden transition-all hover:shadow-lg p-0 gap-2">
      <CardHeader className="p-0">
        <div className="relative aspect-video w-full overflow-hidden bg-muted">
          {event.thumbnail ? (
            <Image
              src={event.thumbnail}
              alt={event.name}
              fill
              className="object-cover transition-transform group-hover:scale-105"
            />
          ) : (
            <div className="flex h-full items-center justify-center bg-gradient-to-br from-primary/10 to-primary/5">
              <Calendar className="h-16 w-16 text-muted-foreground/40" />
            </div>
          )}
        </div>
      </CardHeader>
      <CardContent className="px-4 grow">
        <div className="mb-2 flex items-start justify-between gap-2">
          <h3 className="line-clamp-2 text-lg font-semibold leading-tight">
            {event.name}
          </h3>
        </div>

        <div className="mb-3 flex items-center gap-2 text-sm text-muted-foreground">
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

        {/* <p className="mb-3 line-clamp-3 text-sm text-muted-foreground">
          {event.content.replace(/<[^>]*>/g, '')}
        </p> */}

        {/* {event.categories && event.categories.length > 0 && (
          <div className="flex flex-wrap gap-1">
            {event.categories.slice(0, 3).map((category) => (
              <Badge key={category.id} variant="secondary" className="text-xs">
                {category.name}
              </Badge>
            ))}
            {event.categories.length > 3 && (
              <Badge variant="outline" className="text-xs">
                +{event.categories.length - 3}
              </Badge>
            )}
          </div>
        )} */}
      </CardContent>

      <CardFooter className="flex items-center justify-between border-t px-4 !py-2">
        <div className="flex items-center gap-2 text-sm text-muted-foreground">
          <User className="h-4 w-4" />
          <span className="line-clamp-1">
            {event.author?.name || 'Unknown'}
          </span>
        </div>
        <div className="flex gap-2">
          <Button asChild size="sm" variant="ghost" className="h-8 w-8 p-0">
            <Link href={`/cmsdesk/historical-events/${event.id}/edit`}>
              <Edit className="h-4 w-4" />
            </Link>
          </Button>
          <Button
            size="sm"
            variant="ghost"
            className="h-8 w-8 p-0 text-destructive hover:bg-destructive/10 hover:text-destructive"
            onClick={() => onDelete(event.id)}
          >
            <Trash2 className="h-4 w-4" />
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
