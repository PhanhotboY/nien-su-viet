import { Calendar, Users, Crown, BookOpen } from 'lucide-react';
import { Card, CardContent } from '@/components/ui/card';

export const EventStatistics = () => {
  interface StatProps {
    icon: React.ReactNode;
    value: string;
    label: string;
    description: string;
  }

  const stats: StatProps[] = [
    {
      icon: <Calendar className="h-8 w-8" />,
      value: '4000+',
      label: 'Năm lịch sử',
      description: 'Từ thời Hồng Bàng đến nay',
    },
    {
      icon: <BookOpen className="h-8 w-8" />,
      value: '1000+',
      label: 'Sự kiện',
      description: 'Được ghi nhận và xác thực',
    },
    {
      icon: <Users className="h-8 w-8" />,
      value: '500+',
      label: 'Nhân vật',
      description: 'Anh hùng dân tộc',
    },
    {
      icon: <Crown className="h-8 w-8" />,
      value: '20+',
      label: 'Triều đại',
      description: 'Qua các thời kỳ',
    },
  ];

  return (
    <section className="container mx-auto px-4 py-6 md:py-8">
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3 md:gap-4">
        {stats.map((stat, index) => (
          <Card
            key={index}
            className="group hover:shadow-md transition-all duration-300 hover:scale-[1.02]"
          >
            <CardContent className="p-4 md:p-5 text-center space-y-2">
              <div className="inline-flex p-2 rounded-full bg-primary/10 text-primary group-hover:bg-primary group-hover:text-primary-foreground transition-colors">
                {stat.icon}
              </div>
              <div className="space-y-1">
                <h3 className="text-2xl md:text-3xl font-bold bg-gradient-to-br from-primary to-secondary text-transparent bg-clip-text">
                  {stat.value}
                </h3>
                <p className="font-semibold text-sm md:text-base">
                  {stat.label}
                </p>
                <p className="text-xs text-muted-foreground hidden md:block">
                  {stat.description}
                </p>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
};
