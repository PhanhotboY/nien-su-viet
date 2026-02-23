import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { Calendar, TrendingUp } from 'lucide-react';
import { getImportantEvents } from '@/content/landing/important-events';
import { toEventPeriodString } from '@/helper/date';
import { getTranslations } from 'next-intl/server';

export const FeaturedPeriods = async () => {
  const tshared = await getTranslations('Shared');
  const importantEvents = await getImportantEvents();

  return (
    <section className="container mx-auto px-4 py-6 md:py-8">
      <div className="space-y-4">
        {/* Section Header */}
        <div className="text-center space-y-2">
          <h2 className="text-2xl md:text-3xl font-bold">
            Các giai đoạn lịch sử{' '}
            <span className="bg-gradient-to-r from-primary to-secondary text-transparent bg-clip-text">
              quan trọng
            </span>
          </h2>
          <p className="text-sm text-muted-foreground max-w-2xl mx-auto">
            Tìm hiểu về các thời kỳ đã định hình nên đất nước Việt Nam
          </p>
        </div>

        {/* Featured Periods Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {importantEvents.map((event, index) => (
            <Card
              key={index}
              className="group hover:shadow-md transition-all duration-300"
            >
              <CardHeader>
                <div className="flex items-start justify-between gap-4">
                  <div className="space-y-1 flex-1">
                    <CardTitle className="text-lg md:text-xl">
                      {event.name}
                    </CardTitle>
                    <div className="flex items-center gap-2 text-xs text-muted-foreground">
                      <Calendar className="h-4 w-4" />

                      <span>{toEventPeriodString(event, tshared)}</span>
                    </div>
                  </div>
                  <Badge variant="default" className="shrink-0">
                    Nổi bật
                  </Badge>
                </div>
                <CardDescription className="text-sm pt-1">
                  {event.description}
                </CardDescription>
              </CardHeader>

              <CardContent>
                <div className="space-y-1">
                  <p className="text-xs font-semibold text-muted-foreground">
                    Điểm nổi bật:
                  </p>
                  <ul className="space-y-1">
                    {event.highlights.map((highlight, idx) => (
                      <li key={idx} className="flex items-start gap-2 text-xs">
                        <span className="text-primary mt-0.5">•</span>
                        <span className="leading-relaxed">{highlight}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </div>
    </section>
  );
};
