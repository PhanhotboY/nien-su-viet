import { getEvent } from '@/services/historical-event.service';
import { notFound } from 'next/navigation';
import { Calendar, User, Tag, Clock } from 'lucide-react';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import TextRenderer from '@/components/TextRenderer';
import Link from 'next/link';
import { formatHistoricalEventDate } from '@/helper/date';
import { getTranslations } from 'next-intl/server';
import { Button } from '@/components/ui/button';

export default async function EventDetailPage({
  params,
}: {
  params: Promise<{ eventId: string }>;
}) {
  const { eventId } = await params;
  const response = await getEvent(eventId);
  const t = await getTranslations('EventPage');
  const tshared = await getTranslations('Shared');

  if (!response?.data) {
    notFound();
  }

  const event = response.data;

  const startDate = formatHistoricalEventDate({
    dateType: event.fromDateType,
    sharedTranslator: tshared,
    year: event.fromYear,
    month: event.fromMonth,
    day: event.fromDay,
  });
  const endDate = formatHistoricalEventDate({
    dateType: event.toDateType,
    sharedTranslator: tshared,
    year: event.toYear,
    month: event.toMonth,
    day: event.toDay,
  });

  return (
    <Card className="border-none py-0">
      {/* Header with Vietnamese pattern */}
      <div className="relative bg-gradient-to-r from-secondary via-primary to-secondary">
        {/*<div className="absolute inset-0 opacity-10">
          <div
            className="h-full w-full"
            style={{
              backgroundImage: `url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='1'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E")`,
            }}
          />
        </div>*/}
        <div className="container text-center px-4 py-16 relative z-10">
          <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold mb-6 drop-shadow-lg text-primary-foreground">
            {event.name}
          </h1>
        </div>
        {/* Bottom wave decoration */}
        <div className="absolute bottom-0 left-0 right-0">
          <svg
            viewBox="0 0 1200 120"
            preserveAspectRatio="none"
            className="w-full h-12 md:h-16"
          >
            <path
              d="M321.39,56.44c58-10.79,114.16-30.13,172-41.86,82.39-16.72,168.19-17.73,250.45-.39C823.78,31,906.67,72,985.66,92.83c70.05,18.48,146.53,26.09,214.34,3V0H0V27.35A600.21,600.21,0,0,0,321.39,56.44Z"
              className="fill-amber-50 dark:fill-gray-950"
            />
          </svg>
        </div>
      </div>
      {/* Main Content */}
      <div className="container mx-auto px-4 py-12 -mt-8">
        {/* Info Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          {/* Date Card */}
          <Card className="border-2 border-secondary backdrop-blur">
            <CardContent className="p-6">
              <div className="flex items-start gap-4">
                <div className="p-3 bg-gradient-to-br from-secondary to-primary rounded-lg">
                  <Calendar className="h-6 w-6 text-white" />
                </div>
                <div className="flex-1">
                  <h3 className="font-semibold mb-2">{tshared('period')}</h3>
                  <div className="space-y-1">
                    <p className="text-sm ">
                      <span className="font-medium">
                        {tshared('period-from')}:
                      </span>{' '}
                      {startDate}
                    </p>
                    {endDate && (
                      <p className="text-sm">
                        <span className="font-medium">
                          {tshared('period-to')}:
                        </span>{' '}
                        {endDate}
                      </p>
                    )}
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Author Card */}
          {event.author && (
            <Card className="border-2 border-primary backdrop-blur">
              <CardContent className="p-6">
                <div className="flex items-start gap-4">
                  <div className="p-3 bg-gradient-to-br from-primary/60 to-primary rounded-lg">
                    <User className="h-6 w-6 text-white" />
                  </div>
                  <div className="flex-1">
                    <h3 className="font-semibold mb-2">{tshared('author')}</h3>
                    <p className="text-sm">{event.author.name}</p>
                  </div>
                </div>
              </CardContent>
            </Card>
          )}

          {/* Categories Card */}
          {event.categories && event.categories.length > 0 && (
            <Card className="border-2 backdrop-blur">
              <CardContent className="p-6">
                <div className="flex items-start gap-4">
                  <div className="p-3 bg-gradient-to-br from-orange-500 to-red-500 rounded-lg">
                    <Tag className="h-6 w-6 text-white" />
                  </div>
                  <div className="flex-1">
                    <h3 className="font-semibold text-gray-900 dark:text-gray-100 mb-3">
                      {tshared('category')}
                    </h3>
                    <div className="flex flex-wrap gap-2">
                      {event.categories.map((cat: any) => (
                        <Badge
                          key={cat.category.id}
                          variant="secondary"
                          className="bg-gradient-to-r from-red-100 to-yellow-100 dark:from-red-950 dark:to-yellow-950 border border-red-200 dark:border-red-800"
                        >
                          {cat.category.name}
                        </Badge>
                      ))}
                    </div>
                  </div>
                </div>
              </CardContent>
            </Card>
          )}
        </div>

        {/* Content Card */}
        <Card className="border-2 border-secondary backdrop-blur">
          <CardContent className="p-8">
            <div className="flex items-center gap-3 mb-6">
              <div className="h-1 w-12 bg-gradient-to-r from-secondary to-primary rounded-full" />
              <h2 className="text-2xl font-bold">{t('detailed-content')}</h2>
              <div className="h-1 flex-1 bg-gradient-to-r from-primary to-secondary rounded-full" />
            </div>

            <Separator className="mb-6 bg-gradient-to-r from-secondary via-primary to-secondary" />

            {event.content ? (
              <div className="prose prose-lg dark:prose-invert max-w-none prose-headings:text-red-900 dark:prose-headings:text-red-400 prose-a:text-red-600 dark:prose-a:text-red-400">
                <TextRenderer content={event.content} />
              </div>
            ) : (
              <p className="italic text-center py-8">{t('empty-content')}</p>
            )}
          </CardContent>
        </Card>

        {/* Back to Timeline */}
        <div className="mt-8 text-center">
          <Button asChild>
            <Link href="/">{t('back-to-timeline')}</Link>
          </Button>
        </div>
      </div>
    </Card>
  );
}
