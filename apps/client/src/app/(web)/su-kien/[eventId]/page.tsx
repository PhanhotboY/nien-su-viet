import { getEvent } from '@/services/historical-event.service';
import { notFound } from 'next/navigation';
import { Calendar, User, Tag, Clock } from 'lucide-react';
import { Card, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import TextRenderer from '@/components/TextRenderer';
import Link from 'next/link';
import { formatHistoricalEventDate } from '@/helper/date';

export default async function EventDetailPage({
  params,
}: {
  params: Promise<{ eventId: string }>;
}) {
  const { eventId } = await params;
  const response = await getEvent(eventId);

  if (!response?.data) {
    notFound();
  }

  const event = response.data;

  const startDate = formatHistoricalEventDate(
    event.fromYear,
    event.fromMonth,
    event.fromDay,
  );
  const endDate = formatHistoricalEventDate(
    event.toYear,
    event.toMonth,
    event.toDay,
  );

  return (
    <div className="bg-gradient-to-b from-amber-50 via-red-50 to-yellow-50 dark:from-gray-950 dark:via-gray-900 dark:to-gray-950">
      {/* Header with Vietnamese pattern */}
      <div className="relative bg-gradient-to-r from-red-700 via-yellow-600 to-red-700 dark:from-red-900 dark:via-yellow-900 dark:to-red-900">
        <div className="absolute inset-0 opacity-10">
          <div
            className="h-full w-full"
            style={{
              backgroundImage: `url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='1'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E")`,
            }}
          />
        </div>
        <div className="container text-center px-4 py-16 relative z-10">
          <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold text-white mb-6 drop-shadow-lg">
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
          <Card className="border-2 border-red-200 dark:border-red-900 bg-white/80 dark:bg-gray-900/80 backdrop-blur">
            <CardContent className="p-6">
              <div className="flex items-start gap-4">
                <div className="p-3 bg-gradient-to-br from-red-500 to-yellow-500 rounded-lg">
                  <Calendar className="h-6 w-6 text-white" />
                </div>
                <div className="flex-1">
                  <h3 className="font-semibold text-gray-900 dark:text-gray-100 mb-2">
                    Thời gian
                  </h3>
                  <div className="space-y-1">
                    <p className="text-sm text-gray-700 dark:text-gray-300">
                      <span className="font-medium">Bắt đầu:</span> {startDate}
                    </p>
                    {endDate && (
                      <p className="text-sm text-gray-700 dark:text-gray-300">
                        <span className="font-medium">Kết thúc:</span> {endDate}
                      </p>
                    )}
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Author Card */}
          {event.author && (
            <Card className="border-2 border-amber-200 dark:border-amber-900 bg-white/80 dark:bg-gray-900/80 backdrop-blur">
              <CardContent className="p-6">
                <div className="flex items-start gap-4">
                  <div className="p-3 bg-gradient-to-br from-amber-500 to-yellow-500 rounded-lg">
                    <User className="h-6 w-6 text-white" />
                  </div>
                  <div className="flex-1">
                    <h3 className="font-semibold text-gray-900 dark:text-gray-100 mb-2">
                      Tác giả
                    </h3>
                    <p className="text-sm text-gray-700 dark:text-gray-300">
                      {event.author.name}
                    </p>
                  </div>
                </div>
              </CardContent>
            </Card>
          )}

          {/* Categories Card */}
          {event.categories && event.categories.length > 0 && (
            <Card className="border-2 border-orange-200 dark:border-orange-900 bg-white/80 dark:bg-gray-900/80 backdrop-blur">
              <CardContent className="p-6">
                <div className="flex items-start gap-4">
                  <div className="p-3 bg-gradient-to-br from-orange-500 to-red-500 rounded-lg">
                    <Tag className="h-6 w-6 text-white" />
                  </div>
                  <div className="flex-1">
                    <h3 className="font-semibold text-gray-900 dark:text-gray-100 mb-3">
                      Danh mục
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
        <Card className="border-2 border-red-200 dark:border-red-900 bg-white/90 dark:bg-gray-900/90 backdrop-blur">
          <CardContent className="p-8">
            <div className="flex items-center gap-3 mb-6">
              <div className="h-1 w-12 bg-gradient-to-r from-red-500 to-yellow-500 rounded-full" />
              <h2 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
                Nội dung chi tiết
              </h2>
              <div className="h-1 flex-1 bg-gradient-to-r from-yellow-500 to-red-500 rounded-full" />
            </div>

            <Separator className="mb-6 bg-gradient-to-r from-red-200 via-yellow-200 to-red-200" />

            {event.content ? (
              <div className="prose prose-lg dark:prose-invert max-w-none prose-headings:text-red-900 dark:prose-headings:text-red-400 prose-a:text-red-600 dark:prose-a:text-red-400">
                <TextRenderer content={event.content} />
              </div>
            ) : (
              <p className="text-gray-500 dark:text-gray-400 italic text-center py-8">
                Nội dung đang được cập nhật...
              </p>
            )}
          </CardContent>
        </Card>

        {/* Back to Timeline */}
        <div className="mt-8 text-center">
          <Link
            href="/timeline"
            className="inline-flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-red-600 to-yellow-600 hover:from-red-700 hover:to-yellow-700 text-white font-medium rounded-lg shadow-lg hover:shadow-xl transition-all duration-200"
          >
            <Clock className="h-4 w-4" />
            Quay lại dòng thời gian
          </Link>
        </div>
      </div>
    </div>
  );
}
