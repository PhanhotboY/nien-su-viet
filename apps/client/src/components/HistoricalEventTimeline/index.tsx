'use client';

import { useEffect, useRef, useState } from 'react';
import { DataSet } from 'vis-data';
import { Timeline, TimelineOptions } from 'vis-timeline';
import { ChevronLeft, ChevronRight, ZoomIn, ZoomOut } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { VisItem } from '@/interfaces/vis.interface';
import { createDate } from '@/helper/date';
import { components } from '@nsv-interfaces/nsv-api-documentation';
import { IPaginatedResponse } from '../../interfaces/response.interface';

import 'vis-timeline/styles/vis-timeline-graph2d.css';
import './index.css';
import { useRouter } from 'next/navigation';
import { getEvents } from '@/services/historical-event.service';
import { EventDetailDialog } from './EventDetailDialog';

export function HistoricalEventTimeline() {
  const [events, setEvents] = useState<IPaginatedResponse<
    components['schemas']['HistoricalEventBriefResponseDto']
  > | null>(null);
  const [previewItemId, setPreviewItemId] = useState<string | null>(null);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    getEvents({ limit: '1000' })
      .then((res) => {
        if (res.data && res.statusCode <= 400) {
          return setEvents(res.data);
        }
      })
      .catch(setError);
  }, []);
  const timelineRef = useRef<HTMLDivElement>(null);
  const router = useRouter();
  const [timeline, setTimeline] = useState<Timeline | null>(null);

  useEffect(() => {
    if (!timelineRef.current || !events) return;

    const items = new DataSet<VisItem>(
      events.data.map((event) => {
        const start = createDate(
          event.fromYear,
          event.fromMonth!,
          event.fromDay!,
        );
        const end = event.toYear
          ? createDate(event.toYear, event.toMonth!, event.toDay!)
          : undefined;

        return {
          id: event.id,
          content: '',
          start,
          end,
          title: event.name, // Tooltip preview
          type: end ? ('range' as const) : ('point' as const), // Range for events with duration
          // className: event.
          //   .map((c: any) => c.category.slug)
          //   .join(' '),
        };
      }),
    );

    const initStartDate = new Date();
    const initEndDate = new Date();
    initStartDate.setFullYear(initStartDate.getFullYear() - 450);
    initEndDate.setFullYear(initEndDate.getFullYear() + 50);

    const options: TimelineOptions = {
      stack: true,
      editable: false,
      selectable: true,
      showCurrentTime: false,
      start: initStartDate,
      end: initEndDate,
      minHeight: '70vh',
      format: {
        minorLabels: (date: any, scale, step) => {
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
              return date.format('HH:mm');
            case 'minute':
              return date.format('HH:mm');
            case 'second':
              return date.format('s');
            case 'millisecond':
              return date.format('SSS');
            default:
              return '';
          }
        },
        majorLabels: (date: any, scale, step) => {
          const year = date.year().toString();
          switch (scale) {
            case 'year':
              return '';
            case 'month':
              return year;
            case 'week':
              return `${date.format('MMMM')} ${year}`;
            case 'day':
              return `${date.format('MMMM')} ${year}`;
            case 'weekday':
              return `${date.format('MMMM')} ${year}`;
            case 'hour':
              return date.format('ddd D MMMM');
            case 'minute':
              return date.format('ddd D MMMM');
            case 'second':
              return date.format('D MMMM HH:mm');
            case 'millisecond':
              return date.format('HH:mm:ss');
            default:
              return '';
          }
        },
      },
      template: (item) => {
        return `
          <div class="flex items-center justify-between gap-2 px-2 py-1">
          <h2 class="text-sm font-semibold truncate">${item.title}</h2>
            <span class="text-sm font-medium truncate">${item.content}</span>
          </div>
        `;
      },
    };

    const newTimeline = new Timeline(timelineRef.current, items, options);

    // Handle click to open React dialog
    newTimeline.on('select', (properties) => {
      if (properties.items.length > 0) {
        const itemId = properties.items[0];
        setPreviewItemId(itemId);
      }
    });

    setTimeline(newTimeline);

    return () => {
      newTimeline.destroy();
    };
  }, [events, router]);

  const handleZoomIn = () => {
    if (timeline) {
      const currentRange = timeline.getWindow();
      const start = new Date(currentRange.start.valueOf());
      const end = new Date(currentRange.end.valueOf());
      const interval = end.getTime() - start.getTime();
      const newInterval = interval * 0.7;
      const center = new Date((start.getTime() + end.getTime()) / 2);
      const newStart = new Date(center.getTime() - newInterval / 2);
      const newEnd = new Date(center.getTime() + newInterval / 2);
      timeline.setWindow(newStart, newEnd);
    }
  };

  const handleZoomOut = () => {
    if (timeline) {
      const currentRange = timeline.getWindow();
      const start = new Date(currentRange.start.valueOf());
      const end = new Date(currentRange.end.valueOf());
      const interval = end.getTime() - start.getTime();
      const newInterval = interval * 1.3;
      const center = new Date((start.getTime() + end.getTime()) / 2);
      const newStart = new Date(center.getTime() - newInterval / 2);
      const newEnd = new Date(center.getTime() + newInterval / 2);
      timeline.setWindow(newStart, newEnd);
    }
  };

  const handleZoomDate = (
    date: Date = new Date(),
    zoomUnit: 'day' | 'month' | 'year' = 'day',
    zoomSpan: number = 15,
  ) => {
    if (timeline) {
      const start = new Date(date);
      const end = new Date(date);

      switch (zoomUnit) {
        case 'day':
          start.setDate(start.getDate() - zoomSpan);
          end.setDate(end.getDate() + zoomSpan);
          break;
        case 'month':
          start.setMonth(start.getMonth() - zoomSpan);
          end.setMonth(end.getMonth() + zoomSpan);
          break;
        case 'year':
        default:
          start.setFullYear(start.getFullYear() - zoomSpan);
          end.setFullYear(end.getFullYear() + zoomSpan);
          break;
      }

      timeline.setWindow(start, end);
    }
  };

  const handleToday = () => {
    handleZoomDate(new Date(), 'day', 15);
  };

  function FunctionalButtons() {
    return (
      <div className="p-4 border-b flex gap-2">
        <Button variant="outline" size="icon" onClick={handleZoomIn}>
          <ZoomIn className="h-4 w-4" />
        </Button>
        <Button variant="outline" size="icon" onClick={handleZoomOut}>
          <ZoomOut className="h-4 w-4" />
        </Button>
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            if (timeline) {
              const window = timeline.getWindow();
              const interval = window.end.getTime() - window.start.getTime();
              const distance = interval * 0.2;
              const newStart = new Date(window.start.getTime() - distance);
              const newEnd = new Date(window.end.getTime() - distance);
              timeline.setWindow(newStart, newEnd);
            }
          }}
          title="Move Left"
        >
          <ChevronLeft className="h-4 w-4" />
        </Button>
        <Button variant="outline" onClick={handleToday} className="text-xs">
          Today
        </Button>
        <Button
          variant="outline"
          size="icon"
          onClick={() => {
            if (timeline) {
              const window = timeline.getWindow();
              const interval = window.end.getTime() - window.start.getTime();
              const distance = interval * 0.2;
              const newStart = new Date(window.start.getTime() + distance);
              const newEnd = new Date(window.end.getTime() + distance);
              timeline.setWindow(newStart, newEnd);
            }
          }}
          title="Move Right"
        >
          <ChevronRight className="h-4 w-4" />
        </Button>
      </div>
    );
  }

  return (
    <Card className="py-0">
      <CardContent className="p-0">
        <FunctionalButtons />

        {!error ? (
          <div className="p-6">
            <div ref={timelineRef} className="timeline-container" />
          </div>
        ) : (
          <div className="p-6 text-red-600 font-semibold">
            Lỗi khi tải dữ liệu lịch sử: {error?.message}
          </div>
        )}

        <EventDetailDialog
          eventId={previewItemId}
          open={!!previewItemId}
          onOpenChange={() => setPreviewItemId(null)}
        />
      </CardContent>
    </Card>
  );
}
