import { MousePointer, ZoomIn, Move } from 'lucide-react';

export const QuickTips = () => {
  const tips = [
    {
      icon: <MousePointer className="h-4 w-4" />,
      title: 'Nhấp để xem chi tiết',
    },
    {
      icon: <ZoomIn className="h-4 w-4" />,
      title: 'Phóng to/thu nhỏ',
    },
    {
      icon: <Move className="h-4 w-4" />,
      title: 'Kéo để di chuyển',
    },
  ];

  return (
    <section className="container mx-auto px-4 pb-4">
      <div className="flex flex-wrap items-center justify-center gap-4 md:gap-6 text-sm text-muted-foreground">
        <span className="font-medium">💡 Hướng dẫn:</span>
        {tips.map((tip, index) => (
          <div key={index} className="flex items-center gap-2">
            <div className="text-primary">{tip.icon}</div>
            <span>{tip.title}</span>
          </div>
        ))}
      </div>
    </section>
  );
};
