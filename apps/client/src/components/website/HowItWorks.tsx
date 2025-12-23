import { Card, CardContent, CardHeader, CardTitle } from '../ui/card';
import { MedalIcon, MapIcon, PlaneIcon, GiftIcon } from '../Icons';
import { JSX } from 'react';

interface FeatureProps {
  icon: JSX.Element;
  title: string;
  description: string;
}

const features: FeatureProps[] = [
  {
    icon: <MedalIcon />,
    title: 'Dòng thời gian tương tác',
    description:
      'Khám phá lịch sử Việt Nam qua dòng thời gian trực quan, dễ dàng tìm kiếm và xem chi tiết các sự kiện quan trọng',
  },
  {
    icon: <MapIcon />,
    title: 'Bản đồ lịch sử',
    description:
      'Xem các sự kiện lịch sử trên bản đồ Việt Nam, hiểu rõ bối cảnh địa lý và sự thay đổi lãnh thổ qua các thời kỳ',
  },
  {
    icon: <PlaneIcon />,
    title: 'Tài liệu phong phú',
    description:
      'Truy cập kho tàng tài liệu, hình ảnh, văn bản lịch sử được số hóa và phân loại khoa học',
  },
  {
    icon: <GiftIcon />,
    title: 'Học tập tương tác',
    description:
      'Trải nghiệm học lịch sử thú vị với các câu đố, trắc nghiệm và hoạt động tương tác',
  },
];

export const HowItWorks = () => {
  return (
    <section id="howItWorks" className="container text-center py-24 sm:py-32">
      <h2 className="text-3xl md:text-4xl font-bold ">
        Tính năng{' '}
        <span className="bg-gradient-to-b from-secondary/60 to-secondary text-transparent bg-clip-text">
          Nổi bật{' '}
        </span>
        Của Niên Sử Việt
      </h2>
      <p className="md:w-3/4 mx-auto mt-4 mb-8 text-xl text-muted-foreground">
        Khám phá lịch sử Việt Nam một cách sinh động và dễ dàng với các tính
        năng hiện đại và thân thiện với người dùng
      </p>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
        {features.map(({ icon, title, description }: FeatureProps) => (
          <Card key={title} className="bg-muted/50">
            <CardHeader>
              <CardTitle className="grid gap-4 place-items-center">
                {icon}
                {title}
              </CardTitle>
            </CardHeader>
            <CardContent>{description}</CardContent>
          </Card>
        ))}
      </div>
    </section>
  );
};
