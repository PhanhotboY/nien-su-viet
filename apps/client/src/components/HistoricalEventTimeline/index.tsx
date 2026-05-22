'use client';

import { useEffect, useRef, useState, useCallback, useMemo, memo } from 'react';
import { DataSet } from 'vis-data';
import { Timeline, TimelineOptions } from 'vis-timeline';
import { ChevronLeft, ChevronRight, ZoomIn, ZoomOut } from 'lucide-react';
import { useTranslations } from 'next-intl';
import dynamic from 'next/dynamic';
import 'vis-timeline/styles/vis-timeline-graph2d.css';

import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { VisItem } from '@/interfaces/vis.interface';
import { createDate } from '@/helper/date';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { IPaginatedResponse } from '../../interfaces/response.interface';
import { getEvents } from '@/services/historical-event.service';
import { useRouter } from '@/i18n/navigation';
import { Input } from '../ui/input';
import { createSearchIndex, fuzzySearchWithIndex } from '@/utils/search.util';

import './index.css';

const EventDetailDialog = dynamic(
  () =>
    import('./EventDetailDialog').then((mod) => ({
      default: mod.EventDetailDialog,
    })),
  { ssr: false },
);

const FunctionalButtons = memo(function FunctionalButtons({
  handleZoomIn,
  handleZoomOut,
  handleToday,
  handleZoomDate,
  searchTerm,
  setSearchTerm,
  timeline,
  startDate,
  endDate,
  t,
}: {
  handleZoomIn: () => void;
  handleZoomOut: () => void;
  handleToday: () => void;
  handleZoomDate: (start: Date, end: Date, zoomFactor: number) => void;
  searchTerm: string;
  setSearchTerm: (term: string) => void;
  timeline: Timeline | null;
  startDate: Date;
  endDate: Date;
  t: ReturnType<typeof useTranslations>;
}) {
  // Memoize button callbacks to prevent re-renders
  const handleMovePrev = useCallback(() => {
    if (timeline) {
      const interval = endDate.getTime() - startDate.getTime();
      const distance = interval * 0.2;
      const newStart = new Date(startDate.getTime() - distance);
      const newEnd = new Date(endDate.getTime() - distance);
      handleZoomDate(newStart, newEnd, 1);
    }
  }, [timeline, startDate, endDate, handleZoomDate]);

  const handleMoveNext = useCallback(() => {
    if (timeline) {
      const interval = endDate.getTime() - startDate.getTime();
      const distance = interval * 0.2;
      const newStart = new Date(startDate.getTime() + distance);
      const newEnd = new Date(endDate.getTime() + distance);
      handleZoomDate(newStart, newEnd, 1);
    }
  }, [timeline, startDate, endDate, handleZoomDate]);

  return (
    <div className="p-4 border-b flex gap-2">
      <Button
        variant="outline"
        size="icon"
        onClick={handleZoomIn}
        title={t('zoom-in')}
      >
        <ZoomIn className="h-4 w-4" />
      </Button>
      <Button
        variant="outline"
        size="icon"
        onClick={handleZoomOut}
        title={t('zoom-out')}
      >
        <ZoomOut className="h-4 w-4" />
      </Button>
      <Button
        variant="outline"
        size="icon"
        onClick={handleMovePrev}
        title={t('move-left')}
      >
        <ChevronLeft className="h-4 w-4" />
      </Button>
      <Button
        variant="outline"
        onClick={handleToday}
        className="text-xs"
        title={t('today')}
      >
        {t('today')}
      </Button>
      <Button
        variant="outline"
        size="icon"
        onClick={handleMoveNext}
        title={t('move-right')}
      >
        <ChevronRight className="h-4 w-4" />
      </Button>

      <Input
        type="text"
        placeholder={t('search-events')}
        className="ml-auto max-w-xs"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
    </div>
  );
});

export function HistoricalEventTimeline() {
  const [events, setEvents] = useState<IPaginatedResponse<
    components['schemas']['HistoricalEventBriefResponseDto']
  > | null>(null);
  const [eventsIndex, setEventsIndex] = useState<Map<
    string,
    components['schemas']['HistoricalEventBriefResponseDto'][]
  > | null>(null);
  const [previewItemId, setPreviewItemId] = useState<string | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const t = useTranslations('EventPage');
  const now = new Date();
  // Start with a default range of 15 years before and after today
  const [startDate, setStartDate] = useState(
    new Date(now.getFullYear() - 15, now.getMonth(), now.getDate()),
  );
  const [endDate, setEndDate] = useState(
    new Date(now.getFullYear(), now.getMonth(), now.getDate()),
  );
  const [searchTerm, setSearchTerm] = useState('');
  const [searchTermDebounced, setSearchTermDebounced] = useState('');

  // Debounce search input to avoid excessive filtering
  useEffect(() => {
    const timer = setTimeout(() => {
      setSearchTermDebounced(searchTerm);
    }, 300);
    return () => clearTimeout(timer);
  }, [searchTerm]);

  useEffect(() => {
    getEvents({ limit: '1000' })
      .then((res) => {
        if (Array.isArray(res.data) && res.statusCode <= 400) {
          setEventsIndex(createSearchIndex(res.data));
          return setEvents(res);
        }
      })
      .catch(setError);
  }, []);

  const timelineRef = useRef<HTMLDivElement>(null);
  const router = useRouter();
  const [timeline, setTimeline] = useState<Timeline | null>(null);

  // Cache event timestamps to avoid repeated createDate calls
  const eventTimestamps = useMemo(() => {
    if (!events) return new Map<string, { start: number; end: number }>();
    const cache = new Map<string, { start: number; end: number }>();
    events.data.forEach((event) => {
      const eventStart = createDate(
        event.fromYear,
        event.fromMonth!,
        event.fromDay!,
      ).getTime();
      const eventEnd = event.toYear
        ? createDate(event.toYear, event.toMonth!, event.toDay!).getTime()
        : eventStart;
      cache.set(event.id, { start: eventStart, end: eventEnd });
    });
    return cache;
  }, [events]);

  useEffect(() => {
    if (!timelineRef.current || !events) return;

    let timelineInstance: Timeline | null = null;
    let cancelled = false;
    let idleId: number | undefined;
    let timeoutId: number | undefined;

    const initTimeline = () => {
      if (cancelled || !timelineRef.current) return;

      const items = new DataSet<VisItem>(
        events.data
          .map((event) => {
            const timestamps = eventTimestamps.get(event.id);
            if (!timestamps) return null;

            const { start: eventStartMs, end: eventEndMs } = timestamps;
            const start = new Date(eventStartMs);
            const end =
              eventEndMs !== eventStartMs ? new Date(eventEndMs) : undefined;

            return {
              id: event.id,
              content: '',
              start,
              end,
              title: event.name,
              type: end ? ('range' as const) : ('point' as const),
            };
          })
          .filter(Boolean) as VisItem[],
      );

      // Format label functions - moved outside to reduce allocations
      const formatDateLabel = (date: any, scale: string): string => {
        switch (scale) {
          case 'year':
            return date.year().toString();
          case 'month':
            return date.format('MMM');
          case 'week':
            return date.format('w');
          case 'day':
            return date.format('D');
          case 'weekday':
            return date.format('ddd D');
          case 'hour':
          case 'minute':
            return date.format('HH:mm');
          case 'second':
            return date.format('s');
          case 'millisecond':
            return date.format('SSS');
          default:
            return '';
        }
      };

      const formatMajorLabel = (date: any, scale: string): string => {
        const year = date.year().toString();
        switch (scale) {
          case 'year':
            return '';
          case 'month':
            return year;
          case 'week':
          case 'day':
          case 'weekday':
            return `${date.format('MMMM')} ${year}`;
          case 'hour':
          case 'minute':
            return date.format('ddd D MMMM');
          case 'second':
            return date.format('D MMMM HH:mm');
          case 'millisecond':
            return date.format('HH:mm:ss');
          default:
            return '';
        }
      };

      const options: TimelineOptions = {
        stack: true,
        editable: false,
        selectable: true,
        showCurrentTime: false,
        start: startDate,
        end: endDate,
        minHeight: '70vh',
        maxHeight: '90vh',
        verticalScroll: false,
        horizontalScroll: false,
        format: {
          minorLabels: (date: any, scale) => formatDateLabel(date, scale),
          majorLabels: (date: any, scale) => formatMajorLabel(date, scale),
        },
        template: (item) =>
          `<div class="flex items-center justify-between gap-2 px-2 py-1"><h2 class="text-sm font-semibold truncate">${item.title}</h2><span class="text-sm font-medium truncate">${item.content}</span></div>`,
      };

      timelineInstance = new Timeline(timelineRef.current, items, options);

      // Handle click to open React dialog
      timelineInstance.on('select', (properties) => {
        if (properties.items.length > 0) {
          const itemId = properties.items[0];
          setPreviewItemId(itemId);
        }
      });

      timelineInstance.on('rangechange', (properties) => {
        setStartDate(new Date(properties.start));
        setEndDate(new Date(properties.end));
      });

      setTimeline(timelineInstance);
    };

    if (typeof window !== 'undefined' && window.requestIdleCallback) {
      idleId = window.requestIdleCallback(initTimeline, { timeout: 2500 });
    } else {
      timeoutId = window.setTimeout(initTimeline, 1);
    }

    return () => {
      cancelled = true;
      if (idleId !== undefined && window.cancelIdleCallback) {
        window.cancelIdleCallback(idleId);
      }
      if (timeoutId !== undefined) {
        window.clearTimeout(timeoutId);
      }
      if (timelineInstance) {
        timelineInstance.destroy();
      }
      setTimeline(null);
    };
  }, [events, router]);

  const handleZoomDate = useCallback(
    (
      start: Date = new Date(),
      end: Date = new Date(),
      zoomFactor: number = 1.5,
    ) => {
      if (timeline) {
        const interval = end.getTime() - start.getTime();
        const newInterval = interval * zoomFactor;
        const center = new Date((start.getTime() + end.getTime()) / 2);
        const newStart = new Date(center.getTime() - newInterval / 2);
        const newEnd = new Date(center.getTime() + newInterval / 2);

        setStartDate(newStart);
        setEndDate(newEnd);
        timeline.setWindow(newStart, newEnd);
      }
    },
    [timeline],
  );

  const FIFTEEN_DAYS_IN_MS = useMemo(() => 15 * 24 * 60 * 60 * 1000, []);

  // Zoom to today with a default range of 30 days (15 days before and after)
  const handleToday = useCallback(() => {
    const now = Date.now();
    handleZoomDate(
      new Date(now - FIFTEEN_DAYS_IN_MS),
      new Date(now + FIFTEEN_DAYS_IN_MS),
      1,
    );
  }, [handleZoomDate, FIFTEEN_DAYS_IN_MS]);

  const handleZoomIn = useCallback(() => {
    if (timeline) {
      handleZoomDate(startDate, endDate, 0.7);
    }
  }, [timeline, startDate, endDate, handleZoomDate]);

  const handleZoomOut = useCallback(() => {
    if (timeline) {
      handleZoomDate(startDate, endDate, 1.3);
    }
  }, [timeline, startDate, endDate, handleZoomDate]);

  // Helper to convert event to VisItem
  const eventToVisItem = useCallback(
    (
      event: components['schemas']['HistoricalEventBriefResponseDto'],
    ): VisItem => {
      const timestamps = eventTimestamps.get(event.id);
      if (!timestamps) {
        throw new Error(`No timestamps found for event ${event.id}`);
      }

      const { start: eventStartMs, end: eventEndMs } = timestamps;
      const start = new Date(eventStartMs);
      const end =
        eventEndMs !== eventStartMs ? new Date(eventEndMs) : undefined;

      return {
        id: event.id,
        content: '',
        start,
        end,
        title: event.name,
        type: end ? ('range' as const) : ('point' as const),
      };
    },
    [eventTimestamps],
  );

  // Memoize search results for efficiency
  const searchResults = useMemo(() => {
    if (!searchTermDebounced || !eventsIndex) return null;
    return new Set(
      fuzzySearchWithIndex(searchTermDebounced, eventsIndex).map((e) => e.id),
    );
  }, [searchTermDebounced, eventsIndex]);

  useEffect(
    function filterEventsOnDateChange() {
      if (!timeline || !events) return;

      const startDateMs = startDate.getTime();
      const endDateMs = endDate.getTime();
      const timelineRange = endDateMs - startDateMs;

      const filteredItems = events.data
        .filter((event) => {
          // Apply search filter first (fastest rejection)
          if (searchResults && !searchResults.has(event.id)) {
            return false;
          }

          const timestamps = eventTimestamps.get(event.id);
          if (!timestamps) return false;

          const { start: eventStart, end: eventEnd } = timestamps;
          const eventRange = eventEnd - eventStart;

          // Remove events that are too short compared to the current timeline range to avoid clutter
          if (eventRange > 0 && timelineRange / eventRange > 6) {
            return false;
          }

          // Check if event overlaps with visible timeline
          return (
            (eventStart >= startDateMs && eventStart <= endDateMs) ||
            (eventEnd >= startDateMs && eventEnd <= endDateMs) ||
            (eventStart <= startDateMs && eventEnd >= endDateMs)
          );
        })
        .map(eventToVisItem);

      // Limit to 20 items for performance
      const items = new DataSet<VisItem>(
        filteredItems
          .sort((a, b) => {
            // Prioritize events with duration, then sort by duration (longer first)
            const hasEndA = a.end != null;
            const hasEndB = b.end != null;
            if (hasEndA !== hasEndB) return hasEndB ? 1 : -1;

            const durationA =
              (a.end?.getTime() ?? a.start.getTime()) - a.start.getTime();
            const durationB =
              (b.end?.getTime() ?? b.start.getTime()) - b.start.getTime();
            return durationB - durationA;
          })
          .slice(0, 20),
      );
      try {
        timeline.setItems(items);
      } catch (e) {
        console.error(e);
      }
    },
    [startDate, endDate, searchResults],
  );

  return (
    <Card className="py-0">
      <CardContent className="p-0">
        <FunctionalButtons
          handleZoomIn={handleZoomIn}
          handleZoomOut={handleZoomOut}
          handleToday={handleToday}
          handleZoomDate={handleZoomDate}
          searchTerm={searchTerm}
          setSearchTerm={setSearchTerm}
          timeline={timeline}
          startDate={startDate}
          endDate={endDate}
          t={t}
        />

        {!error ? (
          <div className="p-6">
            <div ref={timelineRef} className="timeline-container" />
          </div>
        ) : (
          <div className="p-6 text-red-600 font-semibold">
            {t('error-loading-event')} {error?.message}
          </div>
        )}

        {!!previewItemId && (
          <EventDetailDialog
            eventId={previewItemId}
            open={!!previewItemId}
            onOpenChange={() => setPreviewItemId(null)}
          />
        )}
      </CardContent>
    </Card>
  );
}
